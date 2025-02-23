package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/noodlecak-e/pix/internal/repository"
	"github.com/noodlecak-e/pix/internal/resources/event"
	"github.com/noodlecak-e/pix/internal/resources/picture"
	"github.com/noodlecak-e/pix/pkg/files/localstorage"
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

	repository := repository.New(db)

	r := gin.Default()

	eventHandler := event.NewHandler(*repository)
	r.POST("/events", eventHandler.Create)
	r.GET("/events/:event_id", eventHandler.Get)
	r.GET("/events", eventHandler.BulkGet)

	// userHandler := user.NewHandler(*repository)
	// r.POST("/events/:event_id/users", userHandler.Create)
	// r.GET("/events/:event_id/users/:user_id", userHandler.Get)
	// r.GET("/events/:event_id/users", userHandler.BulkGet)

	storage := localstorage.NewHandler()

	pictureHandler := picture.NewHandler(*repository, storage)
	r.POST("/events/:event_id/users/:user_id/pictures", pictureHandler.Create)
	r.GET("/events/:event_id/users/:user_id/pictures/:picture_id", pictureHandler.Get)
	r.GET("/events/:event_id/users/:user_id/pictures", pictureHandler.BulkGet)

	r.Run()
}
