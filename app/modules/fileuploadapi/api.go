package fileuploadapi

import (
	"bytes"
	"fantlab/shared"
	"fantlab/utils"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v6"
	"github.com/segmentio/ksuid"
)

type Controller struct {
	services *shared.Services
}

func NewController(services *shared.Services) *Controller {
	return &Controller{services: services}
}

func (c *Controller) UploadImage(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file")

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, "file not found")
		return
	}

	file, err := fileHeader.Open()
	defer file.Close()

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, "bad file")
		return
	}

	img, _, err := image.Decode(file)

	if err != nil {
		utils.ShowError(ctx, http.StatusBadRequest, "provided file is not an image")
		return
	}

	var imageBuf bytes.Buffer
	err = jpeg.Encode(&imageBuf, img, nil)

	reader := bytes.NewReader(imageBuf.Bytes())

	// TODO: add image resize - https://github.com/disintegration/imaging
	// TODO: add webp images

	fileId := ksuid.New().String()
	fileName := fileId + ".jpg"

	_, err = c.services.S3Client.PutObject(
		c.services.Config.MinioImagesBucket,
		fileName,
		reader,
		reader.Size(),
		minio.PutObjectOptions{ContentType: "image/jpeg"},
	)

	if err != nil {
		utils.ShowError(ctx, http.StatusInternalServerError, "failed to save image")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"url": c.services.UrlFormatter.GetImageUrl(fileName),
	})
}
