package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"shop/config"
	"shop/function"
	"strings"
	"time"
)

// AdminLogin 管理员登录
func AdminLogin(c *gin.Context) {
	var code int
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
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

// GetMessage 获取客户留言列表
func GetMessage(c *gin.Context) {
	var code, count int
	var where string
	name := c.Query("name")
	email := c.Query("email")
	subject := c.Query("subject")
	if name != "" {
		where += fmt.Sprintf("name like '%%%v%%' && ", name)
	}
	if email != "" {
		where += fmt.Sprintf("email like '%%%v%%' && ", email)
	}
	if subject != "" {
		where += fmt.Sprintf("subject like '%%%v%%' && ", subject)
	}
	if where != "" {
		where = fmt.Sprintf(" where %v", strings.TrimRight(where, " && "))
	}
	data, err := config.MysqlQuery("select * from message" + where + " order by id desc" + config.PageLimit(c))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": err.Error()})
		return
	}
	err = config.Mysql.QueryRow("select count(id) from message" + where).Scan(&count)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 501, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": "操作成功",
		"count":   count,
		"data":    data,
	})
}

// DelMessage 删除客户留言
func DelMessage(c *gin.Context) {
	ids := c.Query("ids")
	exec, _ := config.Mysql.Exec("delete from message where id " + "in(" + ids + ")")
	affected, _ := exec.RowsAffected()
	if affected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// GetProductList 获取产品列表
func GetProductList(c *gin.Context) {
	var code, count int
	var where string
	name := c.Query("name")
	category := c.Query("category")
	if name != "" {
		where += fmt.Sprintf("name like '%%%v%%' && ", name)
	}
	if category != "" {
		where += fmt.Sprintf("category like '%%%v%%' && ", category)
	}
	if where != "" {
		where = fmt.Sprintf(" where %v", strings.TrimRight(where, " && "))
	}
	data, err := config.MysqlQuery("select * from product" + where + " order by id desc" + config.PageLimit(c))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": err.Error()})
		return
	}
	err = config.Mysql.QueryRow("select count(id) from product" + where).Scan(&count)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 501, "message": err.Error()})
		return
	}
	if len(data) > 0 {
		var domain string
		_ = config.Mysql.QueryRow("select domain from config").Scan(&domain)
		for k, val := range data {
			data[k]["album"] = domain + val["album"].(string)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": "操作成功",
		"count":   count,
		"data":    data,
	})
}

// trimAlbumDomain 去掉图册地址中的平台域名前缀，避免保存时重复拼接
func trimAlbumDomain(album string) string {
	var domain string
	_ = config.Mysql.QueryRow("select domain from config").Scan(&domain)
	if domain != "" {
		album = strings.TrimPrefix(album, domain)
	}
	return album
}

// UploadImage 上传图片
func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "获取上传文件失败"})
		return
	}
	if !function.IsFileExist("./uploads") {
		_ = os.Mkdir("./uploads", 0777)
	}
	ok, _, _, fileAddress := function.SaveImageFile(c, "./uploads", file)
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "上传失败，仅支持png/jpg/gif/jpeg图片"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "操作成功",
		"url":     "/uploads" + fileAddress,
	})
}

// AddProduct 添加产品
func AddProduct(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	album, _ := data["album"].(string)
	timer := time.Now().Format("2006-01-02 15:04:05")
	exec, _ := config.Mysql.Exec(`insert into product (name,price,level,category,Introduction,purity,album,details,ctime) values (?,?,?,?,?,?,?,?,?)`,
		data["name"], data["price"], data["level"], data["category"], data["Introduction"], data["purity"], trimAlbumDomain(album), data["details"], timer)
	id, _ := exec.LastInsertId()
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// SaveProduct 修改产品
func SaveProduct(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	album, _ := data["album"].(string)
	exec, _ := config.Mysql.Exec(`update product set name=?,price=?,level=?,category=?,Introduction=?,purity=?,album=?,details=? where id=?`,
		data["name"], data["price"], data["level"], data["category"], data["Introduction"], data["purity"], trimAlbumDomain(album), data["details"], data["id"])
	affected, _ := exec.RowsAffected()
	if affected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// DelProduct 删除产品
func DelProduct(c *gin.Context) {
	ids := c.Query("ids")
	exec, _ := config.Mysql.Exec("delete from product where id " + "in(" + ids + ")")
	affected, _ := exec.RowsAffected()
	if affected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// GetCategory 获取分类列表
func GetCategory(c *gin.Context) {
	var code, count int
	var where string
	name := c.Query("name")
	if name != "" {
		where = fmt.Sprintf(" where name like '%%%v%%'", name)
	}
	data, err := config.MysqlQuery("select * from category" + where + " order by id asc" + config.PageLimit(c))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": err.Error()})
		return
	}
	err = config.Mysql.QueryRow("select count(id) from category" + where).Scan(&count)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 501, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": "操作成功",
		"count":   count,
		"data":    data,
	})
}

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	exec, _ := config.Mysql.Exec(`insert into category (name) values (?)`, data["name"])
	id, _ := exec.LastInsertId()
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// SaveCategory 修改分类
func SaveCategory(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	exec, _ := config.Mysql.Exec(`update category set name=? where id=?`, data["name"], data["id"])
	affected, _ := exec.RowsAffected()
	if affected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// DelCategory 删除分类
func DelCategory(c *gin.Context) {
	ids := c.Query("ids")
	exec, _ := config.Mysql.Exec("delete from category where id " + "in(" + ids + ")")
	affected, _ := exec.RowsAffected()
	if affected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// GetContactSetting 获取联系方式配置
func GetContactSetting(c *gin.Context) {
	data, err := config.MysqlQuery("select * from contact order by id asc limit 1")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "操作成功",
		"data":    data,
	})
}

// SaveContactSetting 保存联系方式配置
func SaveContactSetting(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	email, _ := data["email"].(string)
	phone, _ := data["phone"].(string)
	address, _ := data["address"].(string)
	businessHours, _ := data["business_hours"].(string)
	var count int
	_ = config.Mysql.QueryRow("select count(id) from contact").Scan(&count)
	var exec sql.Result
	var err error
	if count == 0 {
		exec, err = config.Mysql.Exec(`insert into contact (email,phone,address,business_hours) values (?,?,?,?)`,
			email, phone, address, businessHours)
	} else {
		exec, err = config.Mysql.Exec(`update contact set email=?,phone=?,address=?,business_hours=? order by id asc limit 1`,
			email, phone, address, businessHours)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": err.Error()})
		return
	}
	affected, _ := exec.RowsAffected()
	if affected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// GetConfigSetting 获取系统配置
func GetConfigSetting(c *gin.Context) {
	data, err := config.MysqlQuery("select * from config order by id asc limit 1")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "操作成功",
		"data":    data,
	})
}

// SaveConfigSetting 保存系统配置
func SaveConfigSetting(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	domain, _ := data["domain"].(string)
	var count int
	_ = config.Mysql.QueryRow("select count(id) from config").Scan(&count)
	var exec sql.Result
	var err error
	if count == 0 {
		exec, err = config.Mysql.Exec(`insert into config (domain) values (?)`, domain)
	} else {
		exec, err = config.Mysql.Exec(`update config set domain=? order by id asc limit 1`, domain)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": err.Error()})
		return
	}
	affected, _ := exec.RowsAffected()
	if affected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}
