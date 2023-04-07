package delievery

import (
	"context"
	"legocy-minio-storage/proto"
	"legocy-minio-storage/proto/mapper"
	"log"
)

func (h ImageServer) UploadImage(
	ctx context.Context, req *proto.UploadImageRequest) (*proto.UploadImageResponse, error) {

	log.Printf("request %v", req.Meta.BucketName)

	image := *mapper.FromUploadImageRequest(req)
	log.Println(image.ID)

	url, err := h.storage.UploadFile(ctx, image, req.Meta.BucketName)
	if err != nil {
		return nil, err
	}

	response := &proto.UploadImageResponse{
		ImageURL: url,
	}

	return response, nil
}
