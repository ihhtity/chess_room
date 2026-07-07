package middleware

import (
	"net/http"
	"strings"

	"chess-room-backend/pkg/config"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		isAllowed := false
		if origin == "" {
			isAllowed = true
		} else {
			for _, allowedOrigin := range config.Cfg.CORS.AllowedOrigins {
				if strings.HasPrefix(origin, allowedOrigin) {
					isAllowed = true
					break
				}
			}
		}

		if isAllowed {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
