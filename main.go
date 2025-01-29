package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/noodlecak-e/pix/internal/resources/event"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=pix_user password=pix_pass dbname=pix port=5432 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	eventHandler := event.NewHandler()

	r := gin.Default()

	r.POST("/events", eventHandler.Create)

	r.Run()
}
