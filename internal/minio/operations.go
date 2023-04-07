package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"legocy-minio-storage/internal/domain/image/models"
	"legocy-minio-storage/pkg/helpers"
	"log"
)

func (m *MinioProvider) UploadFile(ctx context.Context, object models.ImageUnit, bucketName string) (string, error) {

	if !IsValidBucketName(bucketName) {
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
	return object.GetObjectFilepath(bucketName, imageName), err
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

func (m *MinioProvider) DownloadFile(ctx context.Context, bucketName string, imageName string) ([]byte, error) {
	file, err := m.client.GetObject(
		ctx,
		bucketName,
		imageName,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return helpers.StreamToByte(file), nil
}
