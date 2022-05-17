package aws

import (
	"education-aws/config"
	"fmt"
	"io"
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
		fmt.Println(err)
		panic(err)
	}

	return &S3Handler{
		cfg:        cfg,
		downloader: s3manager.NewDownloader(session),
		uploader:   s3manager.NewUploader(session),
	}
}

func (handler *S3Handler) Upload(file io.Reader, filename string) error { // nolint:interfacer // ...
	ui := &s3manager.UploadInput{
		Bucket: aws.String("elasticbeanstalk-us-east-1-317712438203"),
		ACL:    aws.String("public-read-write"),
		Key:    aws.String(path.Join("gitlab_file_loader", "uploaded-files", filename)),
		Body:   file,
	}
	_, err := handler.uploader.Upload(ui)

	if err != nil {
		err = fmt.Errorf("%v; bucket name: %s; region: %s; %w", *ui.Bucket, handler.cfg.BucketName, handler.cfg.Region, err)
		return err
	}

	return nil
}

func (handler *S3Handler) Download(filename string) ([]byte, error) {
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := handler.downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String("elasticbeanstalk-us-east-1-317712438203"),
		Key:    aws.String(path.Join("gitlab_file_loader", "uploaded-files", filename)),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func connectToAWS(accesskeyID, secretKey, region string) (*session.Session, error) {
	awsConfig := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accesskeyID, secretKey, ""),
		Endpoint:    aws.String("https://s3.amazonaws.com"),
	}

	return session.NewSession(awsConfig)
}
