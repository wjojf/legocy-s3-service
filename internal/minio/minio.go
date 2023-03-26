package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"legocy-minio-storage/internal/config"
	"legocy-minio-storage/internal/domain/image"
	"log"
)

type MinioProvider struct {
	minioAuthData
	client *minio.Client
}

func (m *MinioProvider) IsReady() bool {
	return m.client != nil
}

type minioAuthData struct {
	url         string
	user        string
	password    string
	token       string
	secretToken string
	ssl         bool
}

func (m *MinioProvider) Connect() error {
	var err error

	// if already connected - return
	if m.client != nil {
		return nil
	}

	m.client, err = minio.New(m.url, &minio.Options{
		Creds:  credentials.NewStaticV4(m.token, m.secretToken, ""),
		Secure: m.ssl,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func NewMinioProvider(config config.MinioConfig) (image.ImageStorage, error) {
	//Client will be initialized by Connect() method
	return &MinioProvider{
		minioAuthData: minioAuthData{
			password:    config.Password,
			url:         config.Url,
			user:        config.User,
			ssl:         config.Ssl,
			token:       config.Token,
			secretToken: config.SecretToken,
		},
		client: nil,
	}, nil
}
