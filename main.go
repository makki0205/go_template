package main

import (
	"github.com/itsjamie/gin-cors"
	"github.com/gin-gonic/gin"
	"github.com/makki0205/web/websocket"
	"github.com/makki0205/web/midelware"
)

func main(){
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")

	r.LoadHTMLGlob("view/*")
	//ウェルカムページ
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	//wsハンドラ
	r.GET("/ws", func(c *gin.Context) {
		websocket.GetHandle()(c.Writer, c.Request)
	})
	api := r.Group("/api")
	// crosの許可
	api.Use(cors.Middleware(middleware.CorsConfig))
	api.GET("/makki", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "katsuramaki taiki",
			"sex": "man",
			"email": "llxo2_5oxll@icloud.com",
		})
	})
}
