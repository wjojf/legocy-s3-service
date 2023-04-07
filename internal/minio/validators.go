package minio

const (
	UserObjectsBucketName      = "users"
	MarketItemImagesBucketName = "marketItems"
)

func IsValidBucketName(bucketName string) bool {
	return bucketName == UserObjectsBucketName ||
		bucketName == MarketItemImagesBucketName
}
