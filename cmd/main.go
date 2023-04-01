package main

import (
	"legocy-minio-storage/internal/app"
	"os"
)

const configFilepath = "/internal/config/json/config.json"

func main() {
	cwd, _ := os.Getwd()
	_app := app.New(cwd + configFilepath)

	_app.Run()
}
