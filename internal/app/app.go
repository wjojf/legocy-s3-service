package app

import (
	"google.golang.org/grpc"
	"legocy-minio-storage/internal/config"
	"legocy-minio-storage/internal/domain/image"
)

type App struct {
	config  *config.AppConfig
	storage *image.ImageStorage
	server  *grpc.Server
}

func NewApp(configFilepath string) *App {
	app := App{}

	// Setup Config
	app.setConfig(configFilepath)

	//Setup Storage
	err := app.setStorage(app.getConfig().MinioConf)
	if err != nil {
		panic(err)
	}

	// Set Server
	err = app.setS3Sever()
	if err != nil {
		panic(err)
	}

	return &app
}
