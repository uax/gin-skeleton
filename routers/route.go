package routers

import (
	c "gin-skeleton/controller"
	"gin-skeleton/middleware"
	"gin-skeleton/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/test", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		uid, _ := c.Get("uid")
		c.JSON(200, gin.H{
			"message": "test",
			"uid":     uid,
		})
	})
	r.POST("/login", c.Login)

	return r
}
