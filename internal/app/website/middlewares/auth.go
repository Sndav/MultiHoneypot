package middlewares

import (
	"Multi-Honeypot/internal/app/website/errors"
	config2 "Multi-Honeypot/internal/pkg/config"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := c.MustGet("config").(*config2.Config)
		token := c.Request.Header.Get("Authorization")
		if token != "" {
			s := strings.Split(token, " ")
			if len(s) == 2 {
				if s[0] == "Basic" {
					payload, _ := base64.StdEncoding.DecodeString(s[1])
					pair := strings.SplitN(string(payload), ":", 2)
					if len(pair) == 2 {
						if pair[0] == config.Get("global", "username") &&
							pair[1] == config.Get("global", "password") {
							c.Next()
							return
						}
					}
				}
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error_code": errors.CodeForbidden,
		})
		return
	}
}
