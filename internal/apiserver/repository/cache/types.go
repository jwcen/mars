package cache

import (
	"context"
	"time"

	"github.com/ecodeclub/ekit"
)

type Cache interface {
	Set(ctx context.Context, key string, val any, exp time.Duration) error
	Get(ctx context.Context, key string) ekit.AnyValue
}
