package main

import (
	"legocy-minio-storage/internal/app"
	"log"
	"os"
)

const configFilepath = "/internal/config/json/config.json"

func main() {
	cwd, _ := os.Getwd()
	_app := app.NewApp(cwd + configFilepath)

	//TODO:
	log.Println(_app)
}
