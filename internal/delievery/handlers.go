package delievery

import (
	"context"
	"legocy-minio-storage/internal/domain/image/models"
	"legocy-minio-storage/proto"
	"log"
)

func (s LegocyS3Server) UploadImage(
	ctx context.Context, req *proto.UploadImageRequest) (*proto.UploadImageResponse, error) {

	var imageUnit *models.ImageUnit = FromUploadImageRequest(req)
	log.Println(imageUnit.ID)

	//TODO:
	return nil, nil
}
