package config

import "os"

type AWSConfig struct {
	SecretKey   string
	AccessKeyID string
	Region      string
	BucketName  string
}

func NewConfig() *AWSConfig{
	return &AWSConfig{
		SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AccessKeyID: os.Getenv("AWS_ACCESS_KEY_ID"),
		Region: os.Getenv("AWS_REGION"),
		BucketName: os.Getenv("BUCKET_NAME"),
	}
}
