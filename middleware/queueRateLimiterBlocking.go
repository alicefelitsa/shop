package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"strconv"
	"time"
)

// QueueRateLimiterBlocking 并发请求控制器中间件（阻塞版本）
// rps：每秒并发请求数
// burst：突发并发请求数
// maxWait：超时时间
func QueueRateLimiterBlocking(rps int, burst int, maxWait time.Duration) gin.HandlerFunc {
	if rps > 0 && burst > 0 {
		fmt.Println(fmt.Sprintf("并发请求控制器启动：正常%v并发，突发%v并发", rps, burst))
	} else {
		log.Fatal("并发请求控制器启动失败")
	}
	limiter := rate.NewLimiter(rate.Limit(rps), burst)
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), maxWait)
		defer cancel()
		// 记录开始时间
		start := time.Now()
		// 尝试获取令牌，如果超时则返回429
		err := limiter.Wait(ctx)
		// 记录等待时间
		waitTime := time.Since(start)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code":      429,
				"error":     "请求太多，请稍后重试",
				"wait_time": waitTime.String(),
			})
			return
		}
		// 设置限流头部信息
		c.Header("X-RateLimit-Limit", strconv.Itoa(rps))
		c.Header("X-RateLimit-Burst", strconv.Itoa(burst))
		c.Header("X-RateLimit-Remaining", strconv.FormatFloat(limiter.Tokens(), 'f', 0, 64))
		c.Header("X-RateLimit-Reset", time.Now().Add(time.Second).Format(time.RFC3339))
		c.Header("X-RateLimit-Wait-Time", waitTime.String())
		c.Next()
	}
}
