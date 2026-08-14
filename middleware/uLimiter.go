package middleware

import (
	"github.com/gin-gonic/gin"
	libredis "github.com/redis/go-redis/v9"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"log"
)

// ULimiter 并发限流器中间件，定义了限流的规则。格式 ULimiter("10-S") "10-S"表示：每秒最多允许处理 10 个请求
// 你可以根据需要调整，例如 "100-M"（每分钟100次）或 "5000-H"（每小时5000次）
func ULimiter(rules string) gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted(rules)
	if err != nil {
		log.Fatal("创建限流器失败：", err)
	}
	option, _ := libredis.ParseURL("redis://localhost:6379/5")
	client := libredis.NewClient(option)
	stores, err := sredis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix: "limiter_gin",
	})
	if err != nil {
		log.Fatal("ULimiter限流器创建redis链接失败：", err)
	}
	return mgin.NewMiddleware(limiter.New(stores, rate))
}
