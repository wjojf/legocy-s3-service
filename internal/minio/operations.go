package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"legocy-minio-storage/internal/domain/image/models"
	"log"
)

func (m *MinioProvider) UploadFile(ctx context.Context, object models.ImageUnit, bucketName string) (string, error) {

	if !isValidBucketName(bucketName) {
		return "", ErrInvalidBucketName
	}

	err := m.creatBucketIfPossible(ctx, bucketName)
	if err != nil {
		return "", err
	}

	imageName := object.GenerateObjectName()
	log.Println(imageName)

	uploadInfo, err := m.client.PutObject(
		ctx,
		bucketName,
		imageName,
		object.Payload,
		object.PayloadSize,
		minio.PutObjectOptions{},
	)

	log.Println(fmt.Sprintf("Sending Image to Minio: %v", uploadInfo))
	return object.GetObjectURL(m.url, bucketName, imageName), err
}

func (m *MinioProvider) creatBucketIfPossible(ctx context.Context, bucketName string) error {

	// Make a new bucket.
	err := m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	// err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := m.client.BucketExists(ctx, bucketName)
		log.Println("bucket exists: ", exists)
		if errBucketExists == nil && exists {
			log.Println("Bucket already exists")
		} else {
			return err
		}
	}
	log.Println("Successfully created ", bucketName)
	return nil
}
