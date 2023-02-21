package repo

import (
	"context"
	"time"
)

type Cache interface {
	Put(ctx context.Context, key string, value string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}
