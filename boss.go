package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"shop/config"
	"shop/route"
)

func main() {
	gin.SetMode(gin.ReleaseMode)  // 设置为生产模式
	router := route.SetupRouter() //设置路由地址
	//启动服务
	if err := router.Run(":" + config.Cj.GetString("bossPort")); err != nil {
		log.Fatal("服务器启动失败：", err)
	}
}
