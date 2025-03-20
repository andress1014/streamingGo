package repositories

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StreamRepository interface {
	GetVideoStream(ctx context.Context, key string) (io.ReadCloser, error)
}

type streamRepository struct {
	s3Client *s3.Client
	bucket   string
	region   string
}

func NewStreamRepository() (*streamRepository, error) {
	// 🔥 Quemando credenciales (Solo para pruebas locales)
	awsAccessKey := "ddd"
	awsSecretKey := "ddddddd/CFfdPAUYRsZNV3lrOOldRtb8Tg7wg"
	awsRegion := "us-east-1"

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			awsAccessKey, awsSecretKey, "",
		)),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &streamRepository{
		s3Client: client,
		bucket:   "andress1014-go-stream", // Nombre del bucket
		region:   awsRegion,
	}, nil
}

func (r *streamRepository) GetVideoStream(ctx context.Context, key string) (io.ReadCloser, error) {
	resp, err := r.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
