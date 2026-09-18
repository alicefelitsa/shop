package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/controller"
	"shop/middleware"
)

// SetupRouter 设置路由地址
func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	//router.Use(middleware.NewIPRateLimiter(120, 20, 1*time.Minute))
	//router.Use(middleware.QueueRateLimiterBlocking(config.Conf.GetInt("server.queueCapacity"), config.Conf.GetInt("server.suddenCapacity"), 5*time.Second))
	//router.Use(middleware.QueueRateLimiter(config.Conf.GetInt("server.queueCapacity"), config.Conf.GetInt("server.suddenCapacity")))
	//router.Use(middleware.ULimiter("10-S"))
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "页面不存在",
		})
	})
	boss := router.Group("/api/boss", middleware.BossAuth)
	{
		bossController := controller.NewBossController()
		boss.POST("/login", bossController.AdminLogin)
		boss.GET("/logout", bossController.AdminLogout)
		boss.GET("/captcha", bossController.Captcha)
		boss.GET("/auth/user", bossController.AuthUser)
		boss.GET("/GetMessage", bossController.GetMessage)
		boss.GET("/DelMessage", bossController.DelMessage)
		boss.GET("/GetCartIntent", bossController.GetCartIntent)
		boss.GET("/DelCartIntent", bossController.DelCartIntent)
		boss.GET("/GetProductList", bossController.GetProductList)
		boss.POST("/UploadImage", bossController.UploadImage)
		boss.POST("/AddProduct", bossController.AddProduct)
		boss.POST("/SaveProduct", bossController.SaveProduct)
		boss.GET("/DelProduct", bossController.DelProduct)
		boss.GET("/GetCategory", bossController.GetCategory)
		boss.POST("/AddCategory", bossController.AddCategory)
		boss.POST("/SaveCategory", bossController.SaveCategory)
		boss.GET("/DelCategory", bossController.DelCategory)
		boss.GET("/GetContactSetting", bossController.GetContactSetting)
		boss.POST("/SaveContactSetting", bossController.SaveContactSetting)
		boss.GET("/GetConfigSetting", bossController.GetConfigSetting)
		boss.POST("/SaveConfigSetting", bossController.SaveConfigSetting)

	}
	web := router.Group("/api/web")
	{
		webController := controller.NewWebController()
		web.GET("/GetProduct", webController.GetProduct)
		web.GET("/GetProductDetail", webController.GetProductDetail)
		web.POST("/AddMessage", webController.AddMessage)
		web.POST("/AddCartIntent", webController.AddCartIntent)
		web.GET("/GetContactInfo", webController.GetContactInfo)
		web.GET("/GetSiteConfig", webController.GetSiteConfig)

	}
	//上传的图片静态服务（URL 前缀需与 UploadImage 返回的 url 保持一致）
	router.Static("/uploads", "./uploads")
	//ws := router.Group("/ws")
	//{
	//	controller.StartClientManager()
	//	ws.GET("/user", controller.UserWs)
	//	ws.GET("/client", controller.ClientWs)
	//}
	return router
}
