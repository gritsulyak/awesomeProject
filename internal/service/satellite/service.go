package satellite

import (
	"context"
	"errors"
	"time"

	"github.com/gritsulyak/awesomeProject/internal/service/cache"
	"github.com/redis/go-redis/v9"

	"github.com/gritsulyak/awesomeProject/internal/model"
)

//go:generate moq -out repository_moq_test.go . Repository
type Repository interface {
	Create(ctx context.Context, s model.Satellite) error
	GetByName(ctx context.Context, name string) (*model.Satellite, error)
}

//go:generate moq -out cache_moq_test.go . Cache
type Cache[T cache.Cacheable] interface {
	Set(ctx context.Context, key string, value T, expiration time.Duration) error
	Get(ctx context.Context, key string, destination T) error
}

type Service struct {
	repo Repository
	c    Cache[*model.Satellite]
}

func NewService(repo Repository, c Cache[*model.Satellite]) *Service {
	return &Service{
		repo: repo,
		c:    c,
	}
}

func (s *Service) GetSatelliteByName(ctx context.Context, name string) (*model.Satellite, error) {
	var res model.Satellite
	err := s.c.Get(ctx, name, &res)
	if err == nil {
		return &res, nil
	}

	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	resFromDB, err := s.repo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return resFromDB, nil
}
