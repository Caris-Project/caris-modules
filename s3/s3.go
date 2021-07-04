package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
	"time"
)

var (
	s      *session.Session
	u      *s3manager.Uploader
	err    error
	bucket string
)

// Init ...
func Init(region, accessKey, secretKey, bucketKey string) {
	s, err = session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})

	if err != nil {
		log.Fatal(err)
	}
	u = s3manager.NewUploader(s)
	bucket = bucketKey
}

// UploadFileAudio ...
func UploadFileAudio(path string, filename string) (link string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file %q, %v", path, err)
	}
	data, err := u.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(filename),
		Body:        f,
		ContentType: aws.String("audio/mpeg"),
		ACL:         aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}
	return data.Location, err
}

// SignedURL ...
func SignedURL(filename string, d time.Duration) (string, error) {
	svc := s3.New(s)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	return req.Presign(d)
}
