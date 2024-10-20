package infra

import (
	"calamity-ai/internal"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Greeting struct {
	gorm.Model
	Content string
}

type BodyDto struct {
	Content string
}

func SetupRouter() (router *gin.Engine) {
	dsn := internal.DatabaseConfig.DatabaseDsn
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	router = gin.Default()

	router.GET("/t", func(ctx *gin.Context) {
		var greetings []Greeting
		db.Find(&greetings)

		ctx.JSON(200, gin.H{"greetings": greetings})
	})

	router.POST("/t", func(ctx *gin.Context) {
		var body BodyDto
		ctx.ShouldBindJSON(&body)

		greeting := Greeting{Content: body.Content}

		db.Create(&greeting)

		ctx.JSON(201, gin.H{"message": "success"})
	})

	return
}
