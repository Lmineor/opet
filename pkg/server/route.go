package server

import (
	"github.com/gin-gonic/gin"
	"opet/pkg/server/api/v1"
	"opet/pkg/server/common"
	mdw "opet/pkg/server/middlewares"
)

func initGinEngine(flag *Flags) *gin.Engine {

	// 设置gin模式
	gin.SetMode(flag.GinMode)
	engine := gin.Default()

	engine.Use(mdw.Cors())

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(common.AllGood, gin.H{"message": "pong"})
	})

	v1 := engine.Group("v1")
	{
		v1.GET("key", api.PutKey)
		v1.POST("key", api.PutKey)
		v1.GET("prefix_key", api.QueryKeyWithPrefix)
	}
	return engine
}
