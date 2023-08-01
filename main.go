package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Asad2730/S3_UploadSave_Go_Gin/routes"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("fail to load .env")
	}

	r := gin.Default()

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.MaxMultipartMemory = 8 << 20 //8 MiB

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Printf("error: %v", err.Error())
	}

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	r.GET("/", routes.GetIndex)

	r.POST("/", func(c *gin.Context) {
		routes.PostForm(c, uploader)
	})

	r.Run(":8080")

}
