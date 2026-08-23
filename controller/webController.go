package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"shop/config"
	"shop/function"
	"time"
)

type WebController struct {
	db *gorm.DB
}

// NewWebController 创建控制器
func NewWebController() *WebController {
	return &WebController{
		db: config.Mysql,
	}
}

// GetProduct 获取产品数据
func (wc *WebController) GetProduct(c *gin.Context) {
	var code, count int
	var message string
	productData := make([]map[string]interface{}, 0)
	_ = wc.db.Raw("select * from product order by id desc" + config.PageLimit(c)).Scan(&productData).Error
	if len(productData) > 0 {
		var domain string
		_ = config.Mysql.Raw("select domain from config").Scan(&domain).Error
		for k, val := range productData {
			productData[k]["album"] = domain + val["album"].(string)
			if t, ok := val["ctime"].(time.Time); ok {
				productData[k]["ctime"] = t.Format("2006-01-02 15:04:05")
			}
		}
	}
	_ = wc.db.Raw("select count(id) from product").Scan(&count).Error
	categoryData := make([]map[string]interface{}, 0)
	_ = wc.db.Raw("select * from category order by id asc").Scan(&categoryData).Error
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
func (wc *WebController) GetProductDetail(c *gin.Context) {
	var code int
	var message string
	productData := make([]map[string]interface{}, 0)
	relatedProducts := make([]map[string]interface{}, 0)
	_ = wc.db.Raw("select * from product where id = ?", c.Query("id")).Scan(&productData).Error
	if len(productData) > 0 {
		var domain string
		_ = wc.db.Raw("select domain from config").Scan(&domain).Error
		for k, val := range productData {
			productData[k]["album"] = domain + val["album"].(string)
			if t, ok := val["ctime"].(time.Time); ok {
				productData[k]["ctime"] = t.Format("2006-01-02 15:04:05")
			}
		}
		_ = wc.db.Raw("select * from product where category=? and id != ? order by id desc limit 8", productData[0]["category"], c.Query("id")).Scan(&relatedProducts).Error
		if len(relatedProducts) > 0 {
			for k, val := range relatedProducts {
				relatedProducts[k]["album"] = domain + val["album"].(string)
				if t, ok := val["ctime"].(time.Time); ok {
					relatedProducts[k]["ctime"] = t.Format("2006-01-02 15:04:05")
				}
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
func (wc *WebController) AddMessage(c *gin.Context) {
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
	data["ip"] = ip
	data["ip_address"] = function.GetIpAddress(ip)
	data["ctime"] = time.Now()
	result := wc.db.Table("message").Create(data)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "message": "提交失败"})
		return
	}
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "提交成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "提交失败"})
	}
}

// GetContactInfo 获取联系方式配置
func (wc *WebController) GetContactInfo(c *gin.Context) {
	data := make([]map[string]interface{}, 0)
	err := wc.db.Raw("select * from contact order by id asc limit 1").Scan(&data).Error
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
