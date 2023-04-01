package delievery

import (
	img "legocy-minio-storage/internal/domain/image"
	"legocy-minio-storage/proto"
)

type ImageServer struct {
	proto.S3ServiceServer
	storage img.ImageStorage
}

func NewImageServer(storage img.ImageStorage) *ImageServer {
	return &ImageServer{
		storage: storage,
	}
}
