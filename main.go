package main

import (
	"hotdog/handler"
	"hotdog/service"

	"github.com/gin-gonic/gin"
)

func main() {
	service.Init()

	r := gin.Default()
	r.POST("/translate/zh2en", handler.Zh2En)
	r.POST("/translate/en2zh", handler.En2Zh)
	r.POST("/summarize", handler.Summarize)

	r.Run(":8080")
}
