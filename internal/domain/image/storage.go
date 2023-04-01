package image

import (
	"context"
	"legocy-minio-storage/internal/domain/image/models"
)

type ImageStorage interface {
	Connect() error
	IsReady() bool
	UploadFile(ctx context.Context, image models.ImageUnit, bucketName string) (string, error)
}
