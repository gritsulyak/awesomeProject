package satellite

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/gritsulyak/awesomeProject/internal/model"
)

// BenchmarkGetSatelliteByName_CacheHit measures throughput when the satellite
// is served directly from cache (hot path).
func BenchmarkGetSatelliteByName_CacheHit(b *testing.B) {
	moon := &model.Satellite{Name: "moon"}

	cacheMock := &CacheMock[*model.Satellite]{
		GetFunc: func(_ context.Context, _ string, dst *model.Satellite) error {
			*dst = *moon
			return nil
		},
		SetFunc: func(_ context.Context, _ string, _ *model.Satellite, _ time.Duration) error {
			return nil
		},
	}
	repoMock := &RepositoryMock{}

	svc := NewService(repoMock, cacheMock)
	ctx := context.Background()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = svc.GetSatelliteByName(ctx, "moon")
	}
}

// BenchmarkGetSatelliteByName_CacheMiss measures throughput when cache misses
// and the service falls through to the repository (cold path).
func BenchmarkGetSatelliteByName_CacheMiss(b *testing.B) {
	moon := &model.Satellite{Name: "moon"}

	cacheMock := &CacheMock[*model.Satellite]{
		GetFunc: func(_ context.Context, _ string, _ *model.Satellite) error {
			return redis.Nil
		},
		SetFunc: func(_ context.Context, _ string, _ *model.Satellite, _ time.Duration) error {
			return nil
		},
	}
	repoMock := &RepositoryMock{
		GetByNameFunc: func(_ context.Context, _ string) (*model.Satellite, error) {
			return moon, nil
		},
	}

	svc := NewService(repoMock, cacheMock)
	ctx := context.Background()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = svc.GetSatelliteByName(ctx, "moon")
	}
}
