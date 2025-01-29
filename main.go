package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/noodlecak-e/pix/internal/resources/event"
	"github.com/noodlecak-e/pix/internal/resources/picture"
	"github.com/noodlecak-e/pix/internal/resources/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=pix_user password=pix_pass dbname=pix port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	r := gin.Default()

	eventHandler := event.NewHandler(db)
	r.POST("/events", eventHandler.Create)
	r.GET("/events/:id", eventHandler.Get)

	userHandler := user.NewHandler(db)
	r.POST("/users", userHandler.Create)
	r.GET("/users/:id", userHandler.Get)

	pictureHandler := picture.NewHandler(db)
	r.POST("/pictures", pictureHandler.Create)

	r.Run()
}
