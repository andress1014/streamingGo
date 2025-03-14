package repositories

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
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
	// Cargar configuración de AWS con la región correcta
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &streamRepository{
		s3Client: client,
		bucket:   "test-go-streaming", // 🔹 Nombre del bucket desde la imagen
		region:   "us-east-1",
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
