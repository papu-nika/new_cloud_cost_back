package server

import (
	"io"
	"sync"

	"github.com/gin-gonic/gin"
)

// クライアントを管理するための構造体
type Client struct {
	Channel chan string
}

var clients = make(map[*Client]struct{})
var clientsMutex sync.Mutex

// メッセージを保存
var messages []string = []string{"a", "b", "c"}
var messagesMutex sync.Mutex

// Gin の SSE ハンドラー
func SSEHandler(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

	client := &Client{Channel: make(chan string)}
	clientsMutex.Lock()
	clients[client] = struct{}{}
	clientsMutex.Unlock()

	defer func() {
		clientsMutex.Lock()
		delete(clients, client)
		close(client.Channel)
		clientsMutex.Unlock()
	}()

	// クライアントにリアルタイムデータを送信
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-client.Channel; ok {
			c.SSEvent("message", msg) // Gin の SSE メソッド
			return true
		}
		return false
	})
}

// クライアントにデータをブロードキャスト
func broadcastMessage(msg string) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	for client := range clients {
		client.Channel <- msg
	}
}

// クライアントからのメッセージを受信する POST API
func PostHandler(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

	var jsonData struct {
		Message string `json:"message"`
	}

	// JSON のデコード
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// メッセージを保存
	messagesMutex.Lock()
	messages = append(messages, jsonData.Message)
	messagesMutex.Unlock()

	// すべてのクライアントにメッセージをブロードキャスト
	broadcastMessage(jsonData.Message)

	c.JSON(200, gin.H{"message": "Received successfully"})
}

// 追加: `/messages` エンドポイントで JSON を返す**
func GetMessagesHandler(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	// c.Writer.Header().Set("Access-Control-Allow-Methods", "GET POST")
	// c.Writer.Header().Set("Content-Type", "text/event-stream")
	// c.Writer.Header().Set("Cache-Control", "no-cache")
	// c.Writer.Header().Set("Connection", "keep-alive")

	messagesMutex.Lock()
	defer messagesMutex.Unlock()

	c.JSON(200, gin.H{"messages": messages})
}
