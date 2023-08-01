package routes

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func PostForm(c *gin.Context, uploader *manager.Uploader) {
	file, err := c.FormFile("image")

	if err != nil {
		c.HTML(200, "index.html", gin.H{
			"error": "failed to upload image",
		})
		return
	}

	err = c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)
	if err != nil {
		c.HTML(200, "index.html", gin.H{
			"error": "failed to upload image",
		})
		return
	}

	f, openErr := file.Open()

	if openErr != nil {
		c.HTML(200, "index.html", gin.H{
			"error": "failed to upload image",
		})
		return
	}

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(""),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})

	if uploadErr != nil {
		c.HTML(200, "index.html", gin.H{
			"error": "failed to upload image",
		})
		return
	}

	c.HTML(200, "index.html", gin.H{
		"image": result.Location,
	})
}
