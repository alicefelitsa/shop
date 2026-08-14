package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"shop/config"
	"sync"
	"time"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var Clients = &ClientManager{
	clients:    make(map[string]*Client),
	broadcast:  make(chan *Message, 1000),
	register:   make(chan *Client, 100),
	unregister: make(chan *Client, 100),
	mu:         sync.RWMutex{},
	closed:     false,
}

type ClientManager struct {
	clients    map[string]*Client
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
	closed     bool
	closeOnce  sync.Once
}

// Client 客户端数据结构
type Client struct {
	conn      *websocket.Conn
	ID        string
	Role      int
	UCode     string
	limiter   *rate.Limiter
	send      chan []byte
	closeChan chan struct{}
	mu        sync.RWMutex
	lastPong  time.Time
	closed    bool
	closeOnce sync.Once
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

// StartClientManager 启动客户端管理器
func StartClientManager() {
	go Clients.run()
}

func (manager *ClientManager) run() {
	heartbeatTicker := time.NewTicker(30 * time.Second)
	defer heartbeatTicker.Stop()

	for {
		select {
		case client := <-manager.register:
			manager.mu.Lock()
			manager.clients[client.ID] = client
			manager.mu.Unlock()
			if client.Role == 0 {
				log.Printf("✅ 客服上线: %s (角色: %d)", client.ID, client.Role)
			}
			if client.Role == 1 {
				log.Printf("✅ 访客上线: %s (角色: %d)", client.ID, client.Role)
			}

		case client := <-manager.unregister:
			manager.mu.Lock()
			if _, ok := manager.clients[client.ID]; ok {
				delete(manager.clients, client.ID)
				client.close() // 关闭客户端连接
				if client.Role == 0 {
					log.Printf("❌ 客服离线: %s (角色: %d)", client.ID, client.Role)
				}
				if client.Role == 1 {
					log.Printf("❌ 访客离线: %s (角色: %d)", client.ID, client.Role)
				}
			}
			manager.mu.Unlock()

		case message := <-manager.broadcast:
			manager.mu.RLock()
			client, ok := manager.clients[message.To]
			manager.mu.RUnlock()

			if ok && client != nil {
				// 使用非阻塞方式发送，避免阻塞broadcast通道
				select {
				case client.send <- serializeMessage(message):
					// 发送成功
				default:
					log.Printf("⚠️ 客户端 %s 发送通道已满，丢弃消息", message.To)
					// 不立即注销，只丢弃消息
				}
			}

		case <-heartbeatTicker.C:
			manager.checkHeartbeat()
		}
	}
}

// 检查心跳
func (manager *ClientManager) checkHeartbeat() {
	manager.mu.RLock()
	clients := make([]*Client, 0, len(manager.clients))
	for _, client := range manager.clients {
		clients = append(clients, client)
	}
	manager.mu.RUnlock()

	now := time.Now()
	for _, client := range clients {
		client.mu.RLock()
		lastPong := client.lastPong
		client.mu.RUnlock()

		if now.Sub(lastPong) > 60*time.Second {
			log.Printf("❌ 客户端 %s (角色: %d) 心跳超时，强制断开", client.ID, client.Role)
			manager.unregister <- client
		}
	}
}

// UserWs 用户WebSocket连接
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
		return
	}

	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("❌ 升级连接失败: %v (ucode: %v)", err, ucode)
		return
	}

	client := &Client{
		conn:      conn,
		ID:        ucode,
		Role:      0,
		UCode:     ucode,
		limiter:   rate.NewLimiter(2, 2),
		send:      make(chan []byte, 100),
		closeChan: make(chan struct{}),
		lastPong:  time.Now(),
		closed:    false,
	}

	Clients.register <- client

	// 启动读写协程
	go client.writePump()
	go client.readPump(ucode)
}

// ClientWs 客户端WebSocket连接
func ClientWs(c *gin.Context) {
	uuid := c.Query("uuid")
	ucode := c.Query("ucode")

	if uuid == "" || ucode == "" {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "缺少参数"})
		return
	}

	if !websocket.IsWebSocketUpgrade(c.Request) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": 405,
			"msg":  "请求方法不允许",
		})
		return
	}

	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("❌ 升级连接失败: %v (uuid: %v)", err, uuid)
		return
	}

	client := &Client{
		conn:      conn,
		ID:        uuid,
		Role:      1,
		UCode:     ucode,
		limiter:   rate.NewLimiter(2, 2),
		send:      make(chan []byte, 100),
		closeChan: make(chan struct{}),
		lastPong:  time.Now(),
		closed:    false,
	}

	Clients.register <- client

	// 发送上线通知
	go func() {
		time.Sleep(100 * time.Millisecond)
		onlineReport(client.ID, client.UCode)
	}()

	// 启动读写协程
	go client.writePump()
	go client.readPump(uuid)
}

// 读取消息
func (c *Client) readPump(key string) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("❌ readPump 发生panic: %v (客户端: %v)", r, key)
		}
		Clients.unregister <- c
	}()

	c.conn.SetReadLimit(5120)
	_ = c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.conn.SetPongHandler(func(string) error {
		c.mu.Lock()
		c.lastPong = time.Now()
		c.mu.Unlock()
		_ = c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		log.Printf("💚 收到Pong: %s (角色: %d)", c.ID, c.Role)
		return nil
	})

	for {
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("❌ 读取消息错误: %v (客户端: %v)", err, key)
			}
			break
		}

		if messageType == websocket.TextMessage {
			c.handleMessage(message)
		}

		_ = c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	}
}

// 写入消息
func (c *Client) writePump() {
	pingTicker := time.NewTicker(20 * time.Second)
	defer func() {
		pingTicker.Stop()
		c.closeOnce.Do(func() {
			c.closed = true
			close(c.send)
		})
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// 通道已关闭
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			_ = c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, _ = w.Write(message)

			// 批量写入
			n := len(c.send)
			for i := 0; i < n; i++ {
				_, _ = w.Write([]byte{'\n'})
				msg, ok := <-c.send
				if !ok {
					break
				}
				_, _ = w.Write(msg)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-pingTicker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
			log.Printf("🧡 发送心跳给: %s (角色: %d)", c.ID, c.Role)

		case <-c.closeChan:
			return
		}
	}
}

// 处理消息
func (c *Client) handleMessage(message []byte) {
	var msg Message
	if err := json.Unmarshal(message, &msg); err != nil {
		log.Printf("❌ 解析消息失败: %v", err)
		return
	}

	// 速率限制
	if !c.limiter.Allow() {
		log.Printf("⚠️ 发送速率过快: %s", c.ID)
		return
	}

	// 设置消息时间
	if msg.CTime == "" {
		msg.CTime = time.Now().Format("2006-01-02 15:04:05")
	}
	msg.From = c.ID

	// 转发消息
	if msg.To != "" {
		Clients.broadcast <- &msg
	}
}

// 上线通知
func onlineReport(uuid, ucode string) {
	timer := time.Now().Format("2006-01-02 15:04:05")
	message := &Message{
		From:  uuid,
		To:    ucode,
		Types: "online",
		CTime: timer,
		Role:  1,
	}

	Clients.broadcast <- message

	go clientStatus(uuid, 0, timer)
}

// 离线通知
func offlineReport(uuid, ucode string) {
	timer := time.Now().Format("2006-01-02 15:04:05")
	message := &Message{
		From:  uuid,
		To:    ucode,
		Types: "offline",
		CTime: timer,
		Role:  1,
	}

	Clients.broadcast <- message

	go clientStatus(uuid, 1, timer)
}

// 安全发送消息
func (c *Client) safeSend(message []byte) bool {
	if c.isClosed() {
		return false
	}

	select {
	case c.send <- message:
		return true
	default:
		return false
	}
}

// 检查客户端是否已关闭
func (c *Client) isClosed() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.closed
}

// 关闭连接
func (c *Client) close() {
	c.closeOnce.Do(func() {
		c.mu.Lock()
		c.closed = true
		close(c.closeChan)
		c.mu.Unlock()

		_ = c.conn.Close()

		// 如果是访客，发送离线通知
		if c.Role == 1 {
			offlineReport(c.ID, c.UCode)
		}
	})
}

// 序列化消息
func serializeMessage(msg *Message) []byte {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("❌ 序列化消息失败: %v", err)
		return []byte("{}")
	}
	return data
}

// GetOnlineCount 获取在线客户端数量
func GetOnlineCount() int {
	Clients.mu.RLock()
	defer Clients.mu.RUnlock()
	return len(Clients.clients)
}

// GetOnlineClients 获取在线客户端列表
func GetOnlineClients() []string {
	Clients.mu.RLock()
	defer Clients.mu.RUnlock()

	clients := make([]string, 0, len(Clients.clients))
	for id := range Clients.clients {
		clients = append(clients, id)
	}
	return clients
}

// 访客在线状态调整
func clientStatus(uuid string, status int, timer string) {
	if status == 0 {
		//上线
		_, _ = config.Mysql.Exec("update client set status=?,online_time=? where uuid=?", status, timer, uuid)
	}
	if status == 1 {
		//离线
		_, _ = config.Mysql.Exec("update client set status=?,offline_time=? where uuid=?", status, timer, uuid)
	}
}
