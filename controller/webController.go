package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/config"
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
