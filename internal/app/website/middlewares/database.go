package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetDB(_gorm *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx := _gorm.Begin()
		c.Set("db", tx)
		c.Next()
		if v := tx.Commit(); v.Error != nil {
			panic(v.Error)
		}
	}
}
