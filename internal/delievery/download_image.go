package delievery

import (
	"context"
	"legocy-minio-storage/internal/minio"
	"legocy-minio-storage/proto"
)

func (h ImageServer) DownloadImage(
	ctx context.Context, req *proto.DownloadImageRequest) (*proto.DownloadImageResponse, error) {

	if !minio.IsValidBucketName(req.BucketName) {
		return nil, minio.ErrInvalidBucketName
	}

	// TODO:
	image, err := h.storage.DownloadFile(ctx, req.BucketName, req.ImageName)
	if err != nil {
		return nil, err
	}

	return &proto.DownloadImageResponse{Data: image}, nil
}
