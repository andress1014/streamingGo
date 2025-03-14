package services

import (
	"context"
	"goStreaming/pkg/api/stream/interfaces/repositories"
	"io"
)

type StreamService interface {
	GetStream(ctx context.Context, videoKey string) (io.ReadCloser, error)
}

type streamService struct {
	repo repositories.StreamRepository
}

func NewStreamService(repo repositories.StreamRepository) StreamService {
	return &streamService{repo: repo}
}

func (s *streamService) GetStream(ctx context.Context, videoKey string) (io.ReadCloser, error) {
	return s.repo.GetVideoStream(ctx, videoKey)
}
