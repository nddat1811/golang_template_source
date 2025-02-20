package utils


import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/http"
	"time"
)


var (
	BUCKET_NAME        = ""
	MINIO_ENDPOINT     = ""
	MINIO_USERNAME     = ""
	MINIO_PASSWORD     = ""
	MINIO_PUBLIC_DOMAIN = "https://hcm.mobifone.vn/minio"
)


func getMinioClient() (*minio.Client, error) {
	return minio.New(MINIO_ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(MINIO_USERNAME, MINIO_PASSWORD, ""),
		Secure: false,
	})
}


func UploadImageToMinio(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, NewResponse("Invalid file upload", nil))
		return
	}
	// folderName := c.PostForm("folder")
	fileName := file.Filename

	client, err := getMinioClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewResponse("Failed to connect to Minio", nil))
		return
	}

	fileData, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewResponse("Failed to open file", nil))
		return
	}
	defer fileData.Close()

	objectName := fmt.Sprintf("%s_%s", fileName, time.Now().Format("20060102"))

	_, err = client.PutObject(context.Background(), BUCKET_NAME, objectName, fileData, file.Size, minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewResponse("Failed to upload file to Minio", nil))
		return
	}

	shareLink := fmt.Sprintf("%s/%s/%s", MINIO_PUBLIC_DOMAIN, BUCKET_NAME, objectName)
	c.JSON(http.StatusOK, NewResponse("ok",  shareLink))
}