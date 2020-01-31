package victims

import (
	"Multi-Honeypot/internal/app/website/errors"
	"Multi-Honeypot/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
}

func (*Handler) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error_code": errors.CodeInvalidQuery,
			})
			return
		}
		type victim struct {
			ID        uint      `gorm:"primary_key" json:"id"`
			FromIP    string    `gorm:"size:50" json:"from_ip"`
			Service   string    `gorm:"size:50" json:"service"`
			CreatedAt time.Time `json:"created_at"`
		}
		var victims []victim
		if v := db.Model(&models.Victim{}).Limit(limit).Scan(&victims); v.Error != nil {
			panic(v.Error)
		}
		c.JSON(http.StatusOK, gin.H{
			"error_code": 0,
			"victim":     victims,
		})
	}
}

func (*Handler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)
		id, err := strconv.Atoi(c.Param("victimId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error_code": errors.CodeInvalidResourceID,
			})
			return
		}
		var victim models.Victim
		if v := db.Where(&models.Victim{ID: uint(id)}).First(&victim); v.Error == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_code": errors.CodeNotFound,
			})
			return
		} else if v.Error != nil {
			panic(v.Error)
		}
		c.JSON(http.StatusOK, gin.H{
			"error_code": 0,
			"victim":     victim,
		})
	}
}
