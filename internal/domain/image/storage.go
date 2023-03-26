package image

import (
	"context"
	"legocy-minio-storage/internal/domain/image/models"
)

type ImageStorage interface {
	Connect() error
	IsReady() bool
	UploadFile(context.Context, models.ImageUnit, string) (string, error)
}
