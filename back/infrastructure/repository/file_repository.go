package repository

import (
	"context"
	"fmt"
	"io"
	"os"

	"back/domain/repository"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type fileRepository struct {
	storage *s3.Client
}

func NewFileRepository(storage *s3.Client) repository.FileRepository {
	return &fileRepository{storage: storage}
}

func (r *fileRepository) Upload(file io.ReadSeeker, contentType string) (url string, err error) {
	bucket := os.Getenv("AWS_BUCKET")
	if bucket == "" {
		return "", fmt.Errorf("AWS_BUCKET is not set")
	}

	key := fmt.Sprintf("user_icons/%s", uuid.New().String())

	_, err = r.storage.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload to S3: %w", err)
	}

	region := os.Getenv("AWS_DEFAULT_REGION")
	if region == "" {
		region = "ap-northeast-1"
	}

	url = fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, key)
	return url, nil
}
