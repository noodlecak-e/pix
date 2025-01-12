package main

import (
	"github.com/gin-gonic/gin"
	"github.com/noodlecak-e/pix/internal/resources/event"
)

func main() {
	eventHandler := event.NewHandler()

	r := gin.Default()

	r.POST("/event", eventHandler.Create)

	r.Run()
}
