package config

import (
	"os"
	"strings"
)

type AWSConfig struct {
	SecretKey   string
	AccessKeyID string
	Region      string
	BucketName  string
}

func NewConfig() *AWSConfig {
	return &AWSConfig{
		SecretKey:   strings.Trim(os.Getenv("AWS_SECRET_ACCESS_KEY"), `"`),
		AccessKeyID: strings.Trim(os.Getenv("AWS_ACCESS_KEY_ID"), `"`),
		Region:      strings.Trim(os.Getenv("AWS_REGION"), `"`),
		BucketName:  strings.Trim(os.Getenv("BUCKET_NAME"), `"`),
	}
}
