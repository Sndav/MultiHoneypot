package middlewares

import (
	"Multi-Honeypot/internal/pkg/config"
	"github.com/gin-gonic/gin"
)

func SetConfig(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}
