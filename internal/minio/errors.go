package minio

import "errors"

var ErrInvalidBucketName = errors.New("bucket with given name does not exist")
