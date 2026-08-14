package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // 设置为生产模式
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "页面不存在",
		})
	})
	api := router.Group("/api")
	{
		go handleData()
		api.GET("/send", sendData)
	}
	//启动服务
	if err := router.Run(":8888"); err != nil {
		log.Fatal("服务器启动失败：", err)
	}
}

var data = make(chan Task, 100)

// Task 定义任务结构
type Task struct {
	Data   string
	Result chan string
}

// 发送数据
func sendData(c *gin.Context) {
	dataValue := c.Query("data")
	if dataValue == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数data不能为空",
		})
		return
	}

	task := Task{
		Data:   dataValue,
		Result: make(chan string),
	}

	select {
	case data <- task:
		select {
		case res := <-task.Result:
			c.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": fmt.Sprintf("处理完成：%s", res),
			})
		case <-c.Request.Context().Done():
			log.Printf("⚠️ 客户端请求取消，放弃等待结果: %v", dataValue)
		case <-time.After(3 * time.Second):
			c.JSON(http.StatusGatewayTimeout, gin.H{
				"code":    504,
				"message": "任务超时，处理失败",
			})
		}
	default:
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code":    503,
			"message": "发送失败，系统繁忙",
		})
	}
}

func handleData() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("❌ 处理数据发生panic: %v", err)
		}
	}()
	for task := range data {
		// 模拟业务处理耗时
		time.Sleep(2 * time.Second)

		select {
		case task.Result <- task.Data:
			log.Printf("✅ 处理完成并返回结果: %v", task.Data)
		default:
			log.Printf("⚠️ 发送结果超时，丢弃结果: %v", task.Data)
		}
	}
}
