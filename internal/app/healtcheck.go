package app

import (
	"context"
	"legocy-minio-storage/internal/domain/image/models"
)

func (a *App) HealthCheck() (string, error) {
	storage := a.GetStorage()
	return storage.UploadFile(context.Background(), models.ImageUnit{
		ID:          0,
		Payload:     nil,
		PayloadName: "healtcheck",
		PayloadSize: 0,
	}, "users")
}
