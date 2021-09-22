package s3

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/soundreaper/portal/config"
)

// S3 is a wrapper around the s3 session
type S3 struct {
	session *session.Session
}

// NewS3Session creates a new S3 session wrapper
func NewS3Session() *S3 {
	// get the s3 config
	c := config.GetS3Config()

	// create a new S3 session
	s, err := session.NewSession(&aws.Config{Region: aws.String(c.Region), Credentials: credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, "")})
	if err != nil {
		log.Fatal("S3 Connection Failed: ", err)
	}

	return &S3{
		session: s,
	}
}

// Upload will upload file to S3. It takes a multipart file header.
func (s *S3) Upload(fileHeader *multipart.FileHeader) (string, string, error) {
	c := config.GetS3Config()

	// open the file header so that we can read into the file buffer
	file, err := fileHeader.Open()
	if err != nil {
		return "", "", err
	}

	// get the size
	size := fileHeader.Size
	// read the file buffer into a new byte array
	buffer := make([]byte, size)
	_, err = file.Read(buffer)
	if err != nil {
		return "", "", err
	}

	// get the current time for the file name
	// TODO: discuss if we want to change this
	now := time.Now().Nanosecond()
	strconv.Itoa(now)

	tmpFileName := fmt.Sprintf("receipts/%d-%s", now, fileHeader.Filename)

	// upload the file to S3
	_, err = s3.New(s.session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(c.Bucket),
		Key:                  aws.String(tmpFileName),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})

	if err != nil {
		return tmpFileName, "", err
	}

	// return the URL to the s3 image in the bucket
	// TODO: maybe not hard code this if we decide to make the images private rather than public
	url := fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", c.Bucket, c.Region, tmpFileName)
	return url, tmpFileName, nil
}
