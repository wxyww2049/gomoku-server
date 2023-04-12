package router

import (
	"github.com/gin-gonic/gin"
	"gomoku-server/middleware"
	"gomoku-server/pkg/app"
	"gomoku-server/service"
	"gomoku-server/websocket"
)

func Setup(engine *gin.Engine) {
	//处理cors
	engine.Use(middleware.Cors())
	//静态文件
	//engine.Static("")
	user := engine.Group("/user")
	{
		hub := service.ExUserService
		user.POST("/test", app.HandlerFunc(hub.Test))
	}
	m := websocket.InitMelody()
	engine.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
}
