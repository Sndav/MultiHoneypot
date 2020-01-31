package victims

import "github.com/gin-gonic/gin"

func RegisterRouter(g *gin.Engine) {
	h := &Handler{}
	r := g.Group("victim")
	r.GET("", h.List())
	r.GET(":victimId", h.Get())
}
