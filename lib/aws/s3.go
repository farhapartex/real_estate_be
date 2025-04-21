package aws

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client *s3.Client
	bucket string
}

func NewS3Client(cfg aws.Config, bucket string) *S3Client {
	return &S3Client{
		client: s3.NewFromConfig(cfg),
		bucket: bucket,
	}
}

func (s *S3Client) GeneratePresignUploadURL(ctx context.Context, key string, contentType string, maxFileSize int64, expiredAt time.Duration) (string, error) {
	client := s3.NewPresignClient(s.client)

	input := &s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(key),
		ContentType:   aws.String(contentType),
		ContentLength: aws.Int64(maxFileSize),
	}

	presignResult, err := client.PresignPutObject(ctx, input, func(opts *s3.PresignOptions) {
		opts.Expires = expiredAt
	})

	if err != nil {
		return "", errors.New("failed to generate upload url")
	}

	return presignResult.URL, nil
}

func (s *S3Client) DeleteFile(ctx context.Context, key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}

	_, err := s.client.DeleteObject(ctx, input)
	if err != nil {
		return errors.New("failed to generate upload url")
	}

	return nil
}
