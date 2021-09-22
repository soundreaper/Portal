package config

import "os"

// S3Config represents the variables that are needed to upload files to an S3 bucket.
type S3Config struct {
	Region    string
	Bucket    string
	SecretKey string
	AccessKey string
}

// GetS3Config will get a new s3 config from the environment.
func GetS3Config() *S3Config {
	return &S3Config{
		Region:    os.Getenv("S3_REGION"),
		Bucket:    os.Getenv("S3_BUCKET"),
		SecretKey: os.Getenv("S3_SECRET_KEY"),
		AccessKey: os.Getenv("S3_ACCESS_KEY"),
	}
}
