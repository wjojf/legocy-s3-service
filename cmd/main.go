package main

import (
	"legocy-minio-storage/internal/app"
	"os"
)

const configFilepath = "/internal/config/json/config.json"

func main() {
	cwd, _ := os.Getwd()
	_app := app.NewApp(cwd + configFilepath)

	_, err := _app.HealthCheck()
	if err != nil {
		panic("Healtcheck failed")
	}
}
