package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"shop/config"
	"shop/function"
	"strconv"
	"strings"
	"time"
)

// AdminLogin 管理员登录
func AdminLogin(c *gin.Context) {
	var code int
	var avatar string
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	_ = config.Mysql.QueryRow("select avatar from config").Scan(&avatar)
	resData, err := config.MysqlQuery("select id from admin where account = ? && password = ?", data["account"], data["password"])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": err.Error()})
		return
	}
	if len(resData) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "账户或密码不正确"})
		return
	}
	token := function.CreateARandomString(30)
	err = config.Redis.Set(config.Ctx, token, resData[0]["id"], time.Minute*43200).Err()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "创建授权令牌失败：" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": "登录成功",
		"token":   token,
		"avatar":  avatar,
		"account": data["account"],
	})
}

// Captcha 获取登录验证码
func Captcha(c *gin.Context) {
	var code int
	var message string
	var data map[string]interface{}
	response, err := http.Get("https://v2.eleadmin.com/api/captcha")
	if err != nil {
		code = 400
		message = "获取验证码失败"
	} else {
		defer response.Body.Close()
		all, err := io.ReadAll(response.Body)
		if err != nil {
			code = 400
			message = "获取验证码失败"
		} else {
			resData := make(map[string]interface{})
			err = json.Unmarshal(all, &resData)
			if err != nil {
				code = 400
				message = "获取验证码失败"
			} else {
				code = 0
				message = "获取数据成功"
				data = resData["data"].(map[string]interface{})
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

// AdminLogout 管理员退出
func AdminLogout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "退出登录",
	})
}

func AuthUser(c *gin.Context) {
	var code int
	var message string
	data := make(map[string]interface{})
	uid, _ := config.Redis.Get(config.Ctx, c.GetHeader("Authorization")).Result()
	resData, err := config.MysqlQuery("select * from admin where id = ?", uid)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "message": err.Error()})
		return
	}
	if len(resData) > 0 {
		code = 0
		message = "操作成功"
		data["avatar"] = "https://cdn.eleadmin.com/20200610/avatar.jpg"
		data["nickname"] = resData[0]["account"]
		data["userId"] = resData[0]["id"]
		data["account"] = resData[0]["account"]
	} else {
		code = 400
		message = "管理员不存在"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

// GetMovie 获取视频
func GetMovie(c *gin.Context) {
	var code, count int
	var where string
	title := c.Query("title")
	category := c.Query("category")
	region := c.Query("region")
	year := c.Query("year")
	format := c.Query("format")
	if title != "" {
		where += fmt.Sprintf("title like '%%%v%%' && ", title)
	}
	if category != "" {
		where += fmt.Sprintf("category like '%%%v%%' && ", category)
	}
	if region != "" {
		where += fmt.Sprintf("region like '%%%v%%' && ", region)
	}
	if year != "" {
		where += fmt.Sprintf("year like '%%%v%%' && ", year)
	}
	if format != "" {
		where += fmt.Sprintf("format like '%%%v%%' && ", format)
	}
	if where != "" {
		where = fmt.Sprintf(" where %v", strings.TrimRight(where, " && "))
	}
	fmt.Println(where)
	data, err := config.MysqlQuery("select * from video" + where + " order by id desc" + config.PageLimit(c))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": err.Error()})
		return
	}
	err = config.Mysql.QueryRow("select count(id) from video" + where).Scan(&count)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 501, "message": err.Error()})
		return
	}
	var playUri string
	_ = config.Mysql.QueryRow("select play_uri from config").Scan(&playUri)
	if len(data) > 0 {
		for k, val := range data {
			data[k]["cover"] = playUri + val["play_url"].(string) + "cover.png"
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": "操作成功",
		"count":   count,
		"data":    data,
	})
}

// AddMovie 添加影片
func AddMovie(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	timer := time.Now().Format("2006-01-02 15:04:05")
	exec, _ := config.Mysql.Exec(`insert into video (title,score,category,region,year,plot,format,play_url,director,episode,ctime) values (?,?,?,?,?,?,?,?,?,?,?)`,
		data["title"], data["score"], data["category"], data["region"], data["year"], data["plot"], data["format"], data["play_url"], data["director"], data["episode"], timer)
	id, _ := exec.LastInsertId()
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// SaveMovie 修改影片
func SaveMovie(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	exec, _ := config.Mysql.Exec(`update video set title=?,score=?,category=?,region=?,year=?,plot=?,format=?,play_url=?,director=?,episode=? where id=?`,
		data["title"], data["score"], data["category"], data["region"], data["year"], data["plot"], data["format"], data["play_url"], data["director"], data["episode"], data["id"])
	affected, _ := exec.RowsAffected()
	if affected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// DelMovie 删除影片
func DelMovie(c *gin.Context) {
	ids := c.Query("ids")
	exec, _ := config.Mysql.Exec("delete from video where id " + "in(" + ids + ")")
	affected, _ := exec.RowsAffected()
	if affected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// GetPlayVideo 获取视频数据
func GetPlayVideo(c *gin.Context) {
	videoData := make(map[string]interface{})
	_ = c.BindJSON(&videoData)
	data, _ := config.MysqlQuery("select * from video where id = ?", videoData["id"])
	if len(data) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "视频不存在"})
		return
	}
	var playUri string
	_ = config.Mysql.QueryRow("select play_uri from config").Scan(&playUri)
	episodeData := make([]map[string]string, 0)
	if data[0]["format"].(string) == "mp4" {
		files, _ := os.ReadDir("." + data[0]["play_url"].(string) + "video")
		if len(files) > 0 {
			for key, file := range files {
				if strings.Contains(file.Name(), ".mp4") == true {
					episodeDataMap := make(map[string]string)
					episodeDataMap["page"] = strconv.Itoa(key + 1)
					episodeDataMap["playUri"] = playUri + data[0]["play_url"].(string) + "video/" + file.Name()
					episodeData = append(episodeData, episodeDataMap)
				}
			}
		}
	}
	if data[0]["format"].(string) == "m3u8" {
		episodeDataMap := make(map[string]string)
		episodeDataMap["page"] = strconv.Itoa(1)
		episodeDataMap["playUri"] = playUri + data[0]["play_url"].(string) + "m3u8/index.m3u8"
		episodeData = append(episodeData, episodeDataMap)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "操作成功",
		"episode": episodeData,
	})
}
