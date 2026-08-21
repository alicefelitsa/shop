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
	//router.Use(middleware.QueueRateLimiterBlocking(config.Cj.GetInt("QueueCapacity"), config.Cj.GetInt("SuddenCapacity"), 5*time.Second))
	//router.Use(middleware.QueueRateLimiter(config.Cj.GetInt("QueueCapacity"), config.Cj.GetInt("SuddenCapacity")))
	//router.Use(middleware.ULimiter("10-S"))
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "页面不存在",
		})
	})
	boss := router.Group("/api/boss", middleware.BossAuthorization)
	{
		boss.POST("/login", controller.AdminLogin)
		boss.GET("/logout", controller.AdminLogout)
		boss.GET("/captcha", controller.Captcha)
		boss.GET("/auth/user", controller.AuthUser)
		boss.GET("/GetMessage", controller.GetMessage)
		boss.GET("/DelMessage", controller.DelMessage)
		boss.GET("/GetProductList", controller.GetProductList)
		boss.POST("/UploadImage", controller.UploadImage)
		boss.POST("/AddProduct", controller.AddProduct)
		boss.POST("/SaveProduct", controller.SaveProduct)
		boss.GET("/DelProduct", controller.DelProduct)
		boss.GET("/GetCategory", controller.GetCategory)
		boss.POST("/AddCategory", controller.AddCategory)
		boss.POST("/SaveCategory", controller.SaveCategory)
		boss.GET("/DelCategory", controller.DelCategory)
		boss.GET("/GetContactSetting", controller.GetContactSetting)
		boss.POST("/SaveContactSetting", controller.SaveContactSetting)
		boss.GET("/GetConfigSetting", controller.GetConfigSetting)
		boss.POST("/SaveConfigSetting", controller.SaveConfigSetting)

	}
	web := router.Group("/api/web")
	{
		web.GET("/GetProduct", controller.GetProduct)
		web.GET("/GetProductDetail", controller.GetProductDetail)
		web.POST("/AddMessage", controller.AddMessage)
		web.GET("/GetContactInfo", controller.GetContactInfo)

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
