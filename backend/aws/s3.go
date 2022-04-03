package aws

import (
	"education-aws/config"
	"fmt"
	"mime/multipart"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Handler struct {
	cfg        *config.AWSConfig
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func NewS3Handler(cfg *config.AWSConfig) *S3Handler {
	session, err := connectToAWS(cfg.AccessKeyID, cfg.SecretKey, cfg.Region)
	if err != nil {
		panic(err)
	}

	return &S3Handler{
		cfg:        cfg,
		downloader: s3manager.NewDownloader(session),
		uploader:   s3manager.NewUploader(session),
	}
}

func (handler *S3Handler) Upload(file multipart.File, filename string) { // nolint:interfacer // ...
	_, err := handler.uploader.Upload(&s3manager.UploadInput{
		Bucket: &handler.cfg.BucketName,
		ACL:    aws.String("public-read"),
		Key:    aws.String(path.Join("file-loader-v2", "uploaded-files", filename)),
		Body:   file,
	})

	if err != nil {
		fmt.Println(handler.cfg.BucketName)
		fmt.Println(handler.cfg.Region)
		panic(err)
	}
}

func (handler *S3Handler) Download(filename string) []byte {
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := handler.downloader.Download(buf, &s3.GetObjectInput{
		Bucket: &handler.cfg.BucketName,
		Key:    aws.String(path.Join("file-loader-v2", "uploaded-files", filename)),
	})

	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func connectToAWS(accesskeyID, secretKey, region string) (*session.Session, error) {
	awsConfig := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accesskeyID, secretKey, ""),
		Endpoint:    aws.String(fmt.Sprintf("https://s3.%s.amazonaws.com", region)),
	}

	return session.NewSession(awsConfig)
}
