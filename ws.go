package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // 设置为生产模式
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "页面不存在",
		})
	})
	ws := router.Group("/ws")
	{
		ws.GET("/user", UserWs)
		ws.GET("/client", ClientWs)
	}
	//启动服务
	if err := router.Run(":9000"); err != nil {
		log.Fatal("服务器启动失败：", err)
	}
}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var Clients sync.Map //储存用户连接

// Client 客户端数据结构
type Client struct {
	conn      *websocket.Conn //连接对象
	key       string          // ucode 或 uuid
	role      int             // 0:客服 1:访客
	ucode     string          // 所属客服
	limiter   *rate.Limiter   //限制发送速率
	send      chan []byte     //接收消息
	done      chan struct{}   //通知关闭心跳
	mu        sync.RWMutex    //读写锁
	closeOnce sync.Once       //用于只执行一次程序
	isClosed  bool            //关闭连接状态
}

// Message 发送消息结构
type Message struct {
	From    string      `json:"from"`
	To      string      `json:"to"`
	Types   string      `json:"types"`
	Content interface{} `json:"content"`
	CTime   string      `json:"ctime"`
	Role    int         `json:"role"`
}

// UserWs 客服连线
func UserWs(c *gin.Context) {
	ucode := c.Query("ucode")
	if ucode == "" {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "缺少ucode"})
		return
	}
	if !websocket.IsWebSocketUpgrade(c.Request) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": 405,
			"msg":  "请求方法不允许",
		})
	}
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("❌ 升级连接失败: %v：%v", err, ucode)
		return
	}
	//注册用户
	client := &Client{
		conn:    conn,
		key:     ucode,
		role:    0,
		ucode:   ucode,
		limiter: rate.NewLimiter(2, 2),
		done:    make(chan struct{}),
		send:    make(chan []byte, 100),
	}
	//保存用户
	Clients.Store(ucode, client)
	log.Println("✅ 客服上线：", ucode)

	//启动读写线程
	go client.writePump()
	go client.readPump()
}

// ClientWs 访客连线
func ClientWs(c *gin.Context) {
	uuid := c.Query("uuid")
	ucode := c.Query("ucode")
	if ucode == "" || uuid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "缺少参数"})
		return
	}
	if !websocket.IsWebSocketUpgrade(c.Request) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": 405,
			"msg":  "请求方法不允许",
		})
	}
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("❌ 升级连接失败: %v：%v", err, ucode)
		return
	}
	//注册用户
	client := &Client{
		conn:    conn,
		key:     uuid,
		role:    1,
		ucode:   ucode,
		limiter: rate.NewLimiter(2, 2),
		done:    make(chan struct{}),
		send:    make(chan []byte, 100),
	}
	//保存用户
	Clients.Store(uuid, client)
	log.Println("✅ 访客上线：", uuid)

	//发送上线通知
	go func() {
		time.Sleep(100 * time.Millisecond)
		client.onlineReport()
	}()

	//启动读写线程
	go client.writePump()
	go client.readPump()
}

// 读取消息
func (c *Client) readPump() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("❌ readPump 发生panic: %v (客户端: %v)", r, c.key)
		}
		c.safeClose()
	}()

	//限制读取消息大小
	c.conn.SetReadLimit(5120)

	// ✅ 设置读超时（60 秒）
	_ = c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	// ✅ 客户回复Pong自动续期
	c.conn.SetPongHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		log.Printf("💚 收到Pong: %s (角色: %d)", c.key, c.role)
		return nil
	})

	// ✅ 客户发来的Ping自动续期
	c.conn.SetPingHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		log.Printf("🧡 收到Ping: %s (角色: %d)", c.key, c.role)
		return nil
	})

	for {
		//接收信息
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("❌ WebSocket连接出错：", c.key, ":", err)
			return
		}
		if messageType == websocket.TextMessage {
			var msg Message
			if err := json.Unmarshal(message, &msg); err != nil {
				log.Printf("❌ 解析消息失败: %v", err)
				continue
			}
			fmt.Println("收到的信息：", msg)

			if c.limiter.Allow() {
				if heObj, status := Clients.Load(msg.To); status {
					he := heObj.(*Client)
					select {
					case he.send <- message: //发送消息
					default:
						log.Printf("⚠️ 客户端 %s 发送通道已满，丢弃消息", he.key)
					}
				}
			} else {
				log.Println("⚠️ 发送速率过快，请稍后再试：", c.key, c.role)
			}
		}
		// ✅ 每次读完刷新
		_ = c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	}
}

// 写入消息
func (c *Client) writePump() {
	pingTicker := time.NewTicker(20 * time.Second)
	defer func() {
		pingTicker.Stop()
		c.safeClose()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// 通道已关闭
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			var msg Message
			_ = json.Unmarshal(message, &msg)
			_ = c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second)) //设置消息写超时
			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("⚠️ 发送超时：", msg.From, "to", c.key, err.Error())
			}

		case <-pingTicker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
			log.Printf("🧡 发送Ping: %s (角色: %d)", c.key, c.role)

		case <-c.done:
			log.Printf("🤍 心跳结束：: %s (角色: %d)", c.key, c.role)
			return
		}
	}
}

// 访客上线通知客服
func (c *Client) onlineReport() {
	if heObj, status := Clients.Load(c.ucode); status {
		he := heObj.(*Client)
		message, _ := json.Marshal(&Message{
			From:  c.key,
			To:    c.ucode,
			Types: "online",
			CTime: time.Now().Format("2006-01-02 15:04:05"),
			Role:  1,
		})
		select {
		case he.send <- message: //发送消息
		default:
			log.Printf("⚠️ 客户端 %s 发送通道已满，丢弃消息", c.ucode)
		}
	}
}

// 访客离线通知客服
func (c *Client) offlineReport() {
	if heObj, status := Clients.Load(c.ucode); status {
		he := heObj.(*Client)
		message, _ := json.Marshal(&Message{
			From:  c.key,
			To:    c.ucode,
			Types: "offline",
			CTime: time.Now().Format("2006-01-02 15:04:05"),
			Role:  1,
		})
		select {
		case he.send <- message: //发送消息
		default:
			log.Printf("⚠️ 客户端 %s 发送通道已满，丢弃消息", c.ucode)
		}
	}
}

// 离线处理，安全关闭
func (c *Client) safeClose() {
	c.closeOnce.Do(func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		//判断是否已关闭
		if c.isClosed {
			return
		}

		//设置状态为关闭
		c.isClosed = true

		//访客通知客服离线
		if c.role == 1 {
			c.offlineReport()
		}

		// 关闭 done 通道
		if c.done != nil {
			close(c.done)
		}

		// 关闭 send 通道
		if c.send != nil {
			close(c.send)
		}

		// 关闭连接
		if c.conn != nil {
			_ = c.conn.Close()
		}

		// 删除客户端映射
		Clients.Delete(c.key)

		// 记录日志
		if c.role == 0 {
			log.Println("⭕ 客服离线：", c.key)
		}
		if c.role == 1 {
			log.Println("⭕ 访客离线：", c.key)
		}
	})
}
