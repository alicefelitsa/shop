package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/config"
	"shop/function"
	"time"
)

// GetProduct 获取产品数据
func GetProduct(c *gin.Context) {
	var code, count int
	var message string
	productData, _ := config.MysqlQuery("select * from product order by id desc" + config.PageLimit(c))
	if len(productData) > 0 {
		var domain string
		_ = config.Mysql.QueryRow("select domain from config").Scan(&domain)
		for k, val := range productData {
			productData[k]["album"] = domain + val["album"].(string)
		}
	}
	_ = config.Mysql.QueryRow("select count(id) from product").Scan(&count)
	categoryData, _ := config.MysqlQuery("select * from category order by id asc")
	message = "获取完成"
	c.JSON(http.StatusOK, gin.H{
		"code":         code,
		"message":      message,
		"productData":  productData,
		"categoryData": categoryData,
		"count":        count,
	})
}

// GetProductDetail 获取产品数据详情
func GetProductDetail(c *gin.Context) {
	var code int
	var message string
	var relatedProducts, productData []map[string]interface{}
	productData, _ = config.MysqlQuery("select * from product where id = ?", c.Query("id"))
	if len(productData) > 0 {
		var domain string
		_ = config.Mysql.QueryRow("select domain from config").Scan(&domain)
		for k, val := range productData {
			productData[k]["album"] = domain + val["album"].(string)
		}
		relatedProducts, _ = config.MysqlQuery("select * from product where category=? and id != ? order by id desc limit 8", productData[0]["category"], c.Query("id"))
		if len(relatedProducts) > 0 {
			for k, val := range relatedProducts {
				relatedProducts[k]["album"] = domain + val["album"].(string)
			}
		}
	}
	message = "获取完成"
	c.JSON(http.StatusOK, gin.H{
		"code":            code,
		"message":         message,
		"productData":     productData,
		"relatedProducts": relatedProducts,
	})
}

// AddMessage 客户提交留言
func AddMessage(c *gin.Context) {
	data := make(map[string]interface{})
	_ = c.BindJSON(&data)
	name, _ := data["name"].(string)
	email, _ := data["email"].(string)
	subject, _ := data["subject"].(string)
	content, _ := data["content"].(string)
	if name == "" || email == "" || subject == "" || content == "" {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "请填写完整的提交信息"})
		return
	}
	ip := c.ClientIP()
	exec, err := config.Mysql.Exec(`insert into message (name,email,ip,ip_address,subject,content,ctime) values (?,?,?,?,?,?,?)`,
		name, email, ip, function.GetIpAddress(ip), subject, content, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": "提交失败"})
		return
	}
	id, _ := exec.LastInsertId()
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "提交成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "提交失败"})
	}
}

// GetContactInfo 获取联系方式配置
func GetContactInfo(c *gin.Context) {
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
