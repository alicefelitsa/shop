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
		boss.GET("/GetMovie", controller.GetMovie)
		boss.POST("/AddMovie", controller.AddMovie)
		boss.POST("/SaveMovie", controller.SaveMovie)
		boss.GET("/DelMovie", controller.DelMovie)
		boss.POST("/GetPlayVideo", controller.GetPlayVideo)

	}
	user := router.Group("/api/user", middleware.UserAuthorization)
	{
		user.POST("/login", controller.Login)
		user.GET("/getMovieList", controller.GetMovieList)
		user.GET("/getVideoDetails", controller.GetVideoDetails)

	}
	file := router.Group("/")
	{
		file.Static("/movie", "./movie")
	}
	//ws := router.Group("/ws")
	//{
	//	controller.StartClientManager()
	//	ws.GET("/user", controller.UserWs)
	//	ws.GET("/client", controller.ClientWs)
	//}
	return router
}
