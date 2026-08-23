package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"net/http"
	"shop/config"
	"shop/function"
	"strings"
	"time"
)

type BossController struct {
	db *gorm.DB
}

// NewBossController 创建控制器
func NewBossController() *BossController {
	return &BossController{
		db: config.Mysql,
	}
}

// AdminLogin 管理员登录
func (bc *BossController) AdminLogin(c *gin.Context) {
	var code int
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	resData := make([]map[string]interface{}, 0)
	err := bc.db.Raw("select id from admin where account = ? && password = ?", data["account"], data["password"]).Scan(&resData).Error
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
func (bc *BossController) Captcha(c *gin.Context) {
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
func (bc *BossController) AdminLogout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "退出登录",
	})
}

func (bc *BossController) AuthUser(c *gin.Context) {
	var code int
	var message string
	data := make(map[string]interface{})
	uid, _ := config.Redis.Get(config.Ctx, c.GetHeader("Authorization")).Result()
	resData := make([]map[string]interface{}, 0)
	err := bc.db.Raw("select * from admin where id = ?", uid).Scan(&resData).Error
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
func (bc *BossController) GetMessage(c *gin.Context) {
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
	data := make([]map[string]interface{}, 0)
	err := bc.db.Raw("select * from message" + where + " order by id desc" + config.PageLimit(c)).Scan(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": err.Error()})
		return
	}
	for _, row := range data {
		for col, val := range row {
			if t, ok := val.(time.Time); ok {
				row[col] = t.Format("2006-01-02 15:04:05")
			}
		}
	}
	err = bc.db.Raw("select count(id) from message" + where).Scan(&count).Error
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
func (bc *BossController) DelMessage(c *gin.Context) {
	ids := c.Query("ids")
	result := bc.db.Exec("delete from message where id " + "in(" + ids + ")")
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// GetProductList 获取产品列表
func (bc *BossController) GetProductList(c *gin.Context) {
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
	data := make([]map[string]interface{}, 0)
	err := bc.db.Raw("select * from product" + where + " order by id desc" + config.PageLimit(c)).Scan(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": err.Error()})
		return
	}
	for _, row := range data {
		for col, val := range row {
			if t, ok := val.(time.Time); ok {
				row[col] = t.Format("2006-01-02 15:04:05")
			}
		}
	}
	err = bc.db.Raw("select count(id) from product" + where).Scan(&count).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 501, "message": err.Error()})
		return
	}
	if len(data) > 0 {
		var domain string
		_ = bc.db.Raw("select domain from config").Scan(&domain).Error
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
func (bc *BossController) trimAlbumDomain(album string) string {
	var domain string
	_ = bc.db.Raw("select domain from config").Scan(&domain).Error
	if domain != "" {
		album = strings.TrimPrefix(album, domain)
	}
	return album
}

// UploadImage 上传图片
func (bc *BossController) UploadImage(c *gin.Context) {
	filePath := "/uploads"
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "获取上传文件失败"})
		return
	}
	fileName, fileUrl, err := function.SaveImageFile(c, "."+filePath, file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "上传失败：" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"message":  "操作成功",
		"url":      filePath + fileUrl,
		"fileName": fileName,
	})
}

// AddProduct 添加产品
func (bc *BossController) AddProduct(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	album, _ := data["album"].(string)
	data["album"] = bc.trimAlbumDomain(album)
	data["ctime"] = time.Now()
	result := bc.db.Table("product").Create(data)
	if result.Error == nil && result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// SaveProduct 修改产品
func (bc *BossController) SaveProduct(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	album, _ := data["album"].(string)
	result := bc.db.Exec(`update product set name=?,price=?,level=?,category=?,Introduction=?,purity=?,album=?,details=? where id=?`,
		data["name"], data["price"], data["level"], data["category"], data["Introduction"], data["purity"], bc.trimAlbumDomain(album), data["details"], data["id"])
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// DelProduct 删除产品
func (bc *BossController) DelProduct(c *gin.Context) {
	ids := c.Query("ids")
	result := bc.db.Exec("delete from product where id " + "in(" + ids + ")")
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// GetCategory 获取分类列表
func (bc *BossController) GetCategory(c *gin.Context) {
	var code, count int
	var where string
	name := c.Query("name")
	if name != "" {
		where = fmt.Sprintf(" where name like '%%%v%%'", name)
	}
	data := make([]map[string]interface{}, 0)
	err := bc.db.Raw("select * from category" + where + " order by id asc" + config.PageLimit(c)).Scan(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": err.Error()})
		return
	}
	err = bc.db.Raw("select count(id) from category" + where).Scan(&count).Error
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
func (bc *BossController) AddCategory(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	result := bc.db.Table("category").Create(data)
	if result.Error == nil && result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// SaveCategory 修改分类
func (bc *BossController) SaveCategory(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	result := bc.db.Exec(`update category set name=? where id=?`, data["name"], data["id"])
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// DelCategory 删除分类
func (bc *BossController) DelCategory(c *gin.Context) {
	ids := c.Query("ids")
	result := bc.db.Exec("delete from category where id " + "in(" + ids + ")")
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// GetContactSetting 获取联系方式配置
func (bc *BossController) GetContactSetting(c *gin.Context) {
	data := make([]map[string]interface{}, 0)
	err := bc.db.Raw("select * from contact order by id asc limit 1").Scan(&data).Error
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
func (bc *BossController) SaveContactSetting(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	email, _ := data["email"].(string)
	phone, _ := data["phone"].(string)
	address, _ := data["address"].(string)
	businessHours, _ := data["business_hours"].(string)
	var count int64
	_ = bc.db.Raw("select count(id) from contact").Scan(&count).Error
	var result *gorm.DB
	if count == 0 {
		result = bc.db.Exec(`insert into contact (email,phone,address,business_hours) values (?,?,?,?)`,
			email, phone, address, businessHours)
	} else {
		result = bc.db.Exec(`update contact set email=?,phone=?,address=?,business_hours=? order by id asc limit 1`,
			email, phone, address, businessHours)
	}
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": result.Error.Error()})
		return
	}
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}

// GetConfigSetting 获取系统配置
func (bc *BossController) GetConfigSetting(c *gin.Context) {
	data := make([]map[string]interface{}, 0)
	err := bc.db.Raw("select * from config order by id asc limit 1").Scan(&data).Error
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
func (bc *BossController) SaveConfigSetting(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	domain, _ := data["domain"].(string)
	var count int64
	_ = bc.db.Raw("select count(id) from config").Scan(&count).Error
	var result *gorm.DB
	if count == 0 {
		result = bc.db.Exec(`insert into config (domain) values (?)`, domain)
	} else {
		result = bc.db.Exec(`update config set domain=? order by id asc limit 1`, domain)
	}
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": result.Error.Error()})
		return
	}
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "操作成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "操作失败"})
	}
}
