package app

import (
	"google.golang.org/grpc"
	"legocy-minio-storage/internal/config"
	"legocy-minio-storage/internal/domain/image"
	"log"
	"net"
)

type App struct {
	config  config.AppConfig
	storage image.ImageStorage
	server  *grpc.Server
}

func New(configFilepath string) *App {
	app := App{}

	// Load Config
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

	if !app.IsReady() {
		panic("DI failed")
	}

	log.Println("App ready")
	return &app
}

func (a *App) Run() {
	listener, err := net.Listen("tcp", a.config.Port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	log.Printf("server started at %v", listener.Addr())
	if err := a.server.Serve(listener); err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}
