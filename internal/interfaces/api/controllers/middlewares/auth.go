// File: internal/interfaces/api/middlewares/middlewares.go
package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Auth là middleware xác thực token JWT
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			c.Abort()
			return
		}

		// Kiểm tra Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		token := parts[1]
		// Đây là nơi để thêm logic xác thực token
		// Ví dụ: validateToken(token)

		// Cho phép tiếp tục nếu token hợp lệ
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Có thể lưu thông tin người dùng vào context
		// c.Set("user_id", userID)

		c.Next()
	}
}

// Logger là middleware ghi log request
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Add query to path if it exists
		if raw != "" {
			path = path + "?" + raw
		}

		// Process request
		c.Next()

		// Log details after request is completed
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method

		log.Printf("[%d] %s %s %s %v", statusCode, method, path, clientIP, latency)
	}
}

// CORS là middleware xử lý Cross-Origin Resource Sharing
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// ErrorHandler là middleware bắt và xử lý lỗi toàn cục
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Chỉ xử lý lỗi nếu có lỗi và response chưa được gửi
		if len(c.Errors) > 0 && !c.Writer.Written() {
			err := c.Errors.Last().Err

			// Gửi response lỗi mặc định
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	}
}

// Recovery là middleware xử lý panic và ghi log
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log lỗi panic
				log.Printf("Panic recovered: %v\n%s", err, debug.Stack())

				// Trả về lỗi 500 cho client
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Internal server error",
				})
			}
		}()

		c.Next()
	}
}

// RateLimiter là struct để kiểm soát số lượng requests
type RateLimiter struct {
	ips    map[string]int64
	mutex  sync.Mutex
	limit  int
	window time.Duration
}

// NewRateLimiter tạo một rate limiter mới
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		ips:    make(map[string]int64),
		limit:  limit,
		window: window,
	}
}

// RateLimit là middleware giới hạn số lượng requests
func (rl *RateLimiter) RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now().UnixNano()

		rl.mutex.Lock()
		defer rl.mutex.Unlock()

		// Xóa các entries cũ
		for key, timestamp := range rl.ips {
			if now-timestamp > rl.window.Nanoseconds() {
				delete(rl.ips, key)
			}
		}

		// Kiểm tra và cập nhật lượt truy cập
		count := 0
		for existingIP := range rl.ips {
			if existingIP == ip {
				count++
			}
		}

		if count >= rl.limit {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			c.Abort()
			return
		}

		// Thêm truy cập mới
		rl.ips[ip+":"+time.Now().String()] = now

		c.Next()
	}
}
