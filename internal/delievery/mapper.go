package delievery

import (
	"bytes"
	"legocy-minio-storage/internal/domain/image/models"
	"legocy-minio-storage/proto"
)

func FromUploadImageRequest(req *proto.UploadImageRequest) *models.ImageUnit {
	m := &models.ImageUnit{
		ID:          int(req.Meta.Id),
		Payload:     bytes.NewReader(req.Data),
		PayloadName: "",
		PayloadSize: int64(len(req.Data)),
	}
	m.PayloadName = m.GenerateObjectName()
	return m
}

func ImageUnitToImageResponse(
	m *models.ImageUnit, baseUrl, bucketName, filename string) *proto.UploadImageResponse {
	return &proto.UploadImageResponse{
		ImageURL: m.GetObjectURL(baseUrl, bucketName, filename),
	}
}
