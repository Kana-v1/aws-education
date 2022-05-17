package tests

import (
	"education-aws/aws"
	"education-aws/config"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var s3Handler *aws.S3Handler
var existingTestFileName = "test_file.txt"
var nonExistingFileName = "non_exists.file"

func init() {
	godotenv.Load("tests.env")
	s3Handler = aws.NewS3Handler(config.NewConfig())
}

func TestS3Upload(t *testing.T) {
	reader, err := os.Open(existingTestFileName)
	if err != nil {
		t.Fail()
	}

	err = s3Handler.Upload(reader, existingTestFileName)
	if err != nil {
		t.Fail()
	}
}

func TestS3DownLoad(t *testing.T) {
	file, err := s3Handler.Download(existingTestFileName)
	if err != nil {
		t.Fail()
	}

	if len(file) <= 0 {
		t.Fail()
	}
}

func TestS3DownloadNonExistsFile(t *testing.T) {
	_, err := s3Handler.Download(nonExistingFileName)
	if err == nil {
		t.Fail()
	}
}
