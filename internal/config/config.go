package config

import (
	"encoding/json"
	"io/ioutil"
	"legocy-minio-storage/pkg/helpers"
)

type AppConfig struct {
	MinioConf MinioConfig `json:"minio"`
}

type MinioConfig struct {
	Url         string `json:"url"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Token       string `json:"token"`
	SecretToken string `json:"secret_token"`
	Ssl         bool   `json:"ssl"`
}

func SetupFromJSON(fp string) (*AppConfig, error) {
	var cfg AppConfig

	if fileExists := helpers.FileExists(fp); !fileExists {
		return nil, ErrConfigFileDoesNotExist
	}

	raw, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
