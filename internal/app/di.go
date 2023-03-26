package app

import (
	"google.golang.org/grpc"
	"legocy-minio-storage/internal/config"
	"legocy-minio-storage/internal/delievery"
	"legocy-minio-storage/internal/domain/image"
	"legocy-minio-storage/internal/minio"
	"legocy-minio-storage/proto"
)

func (a *App) IsReady() bool {
	return a.GetStorage() != nil && a.setS3Sever() != nil
}

func (a *App) getConfig() *config.AppConfig {
	return a.config
}

func (a *App) setConfig(fp string) {
	cfg, err := config.SetupFromJSON(fp)
	if err != nil {
		panic(err)
	}
	a.config = cfg
}

func (a *App) GetStorage() *image.ImageStorage {
	return a.storage
}

func (a *App) setStorage(minioCfg config.MinioConfig) (err error) {
	if a.GetStorage() != nil {
		err = errStorageAlreadySet
		return
	}

	storage, err := minio.NewMinioProvider(
		minioCfg.Url,
		minioCfg.User, minioCfg.Password,
		minioCfg.Token, minioCfg.SecretToken, minioCfg.Ssl,
	)
	if err != nil {
		return
	}

	a.storage = &storage
	return
}

func (a *App) GetS3Server() *grpc.Server {
	return a.server
}

func (a *App) setS3Sever() error {
	if a.GetS3Server() != nil {
		return errServerAlreadySet
	}

	a.server = grpc.NewServer()
	proto.RegisterS3ServiceServer(a.server, delievery.LegocyS3Server{})
	return nil
}
