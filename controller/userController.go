package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"shop/config"
	"shop/function"
	"strconv"
	"strings"
	"time"
)

// Login 用户登录
func Login(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	if len(data) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "参数缺少"})
		return
	}
	resData, _ := config.MysqlQuery("select id from user where account=? && password=?", data["account"], data["password"])
	if len(resData) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "账号或密码错误"})
		return
	}
	token := function.CreateARandomString(30)
	err := config.Redis.Set(config.Ctx, token, resData[0]["id"], time.Minute*43200).Err()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "登录令牌创建失败：" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"token":   token,
		"message": "登录成功",
	})
}

// GetMovieList 获取视频列表
func GetMovieList(c *gin.Context) {
	var code, count int
	var message, parameter, where string
	category := c.Query("category")
	keyword := c.Query("keyword")
	if category != "" {
		parameter += fmt.Sprintf("category like '%%%v%%' && ", category)
	}
	if keyword != "" {
		parameter += fmt.Sprintf("(title like '%%%v%%' || director like '%%%v%%') && ", keyword, keyword)
	}
	if parameter != "" {
		where = fmt.Sprintf(" where %v", strings.TrimRight(parameter, " && "))
	}

	videoData, _ := config.MysqlQuery("select * from video" + where + " order by year desc,id desc" + config.PageLimit(c))
	if len(videoData) > 0 {
		domainHost := "http://" + c.Request.Host
		for k, val := range videoData {
			videoData[k]["cover"] = domainHost + val["play_url"].(string) + "cover.png"
		}
	}
	_ = config.Mysql.QueryRow("select count(id) from video" + where).Scan(&count)
	categoryData, _ := config.MysqlQuery("select * from category order by sort,id")
	message = "获取完成"
	c.JSON(http.StatusOK, gin.H{
		"code":         code,
		"message":      message,
		"videoData":    videoData,
		"categoryData": categoryData,
		"count":        count,
	})
}

// GetVideoDetails 获取视频详情
func GetVideoDetails(c *gin.Context) {
	vid := c.Query("vid")
	if vid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "参数缺少"})
		return
	}
	videoData, _ := config.MysqlQuery("select * from video where id=?", vid)
	if len(videoData) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "视频不存在"})
		return
	}
	domainHost := "http://" + c.Request.Host
	//集数的播放链接
	episodeData := make([]map[string]string, 0)
	if videoData[0]["format"].(string) == "mp4" {
		files, _ := os.ReadDir("." + videoData[0]["play_url"].(string) + "video")
		if len(files) > 0 {
			for key, file := range files {
				//fmt.Println(file.Name())
				if strings.Contains(file.Name(), ".mp4") == true {
					episodeDataMap := make(map[string]string)
					episodeDataMap["page"] = strconv.Itoa(key + 1)
					episodeDataMap["playUri"] = domainHost + videoData[0]["play_url"].(string) + "video/" + file.Name()
					episodeData = append(episodeData, episodeDataMap)
				}
			}
		}
	}
	if videoData[0]["format"].(string) == "m3u8" {
		episodeDataMap := make(map[string]string)
		episodeDataMap["page"] = strconv.Itoa(1)
		episodeDataMap["playUri"] = domainHost + videoData[0]["play_url"].(string) + "m3u8/index.m3u8"
		episodeData = append(episodeData, episodeDataMap)
	}
	//封面图
	videoData[0]["cover"] = domainHost + videoData[0]["play_url"].(string) + "cover.png"
	c.JSON(http.StatusOK, gin.H{
		"code":        0,
		"message":     "获取完成",
		"videoData":   videoData,
		"episodeData": episodeData,
	})
}
