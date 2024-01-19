package server

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/devhindo/owls/chat"
	"github.com/devhindo/owls/video"
)

func RUN() {
	router := gin.Default()

	router.GET("/chat", chat.HandleChat)
	router.GET("/video", video.HandleVideo)

	log.Println("Listening on port 8080...")

	router.Run()
}