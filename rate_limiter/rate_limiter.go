package rate_limiter

import (
	"context"
	"fmt"
	"go-rate-limiter/redis"
	"time"
)

func RateLimit(key string, limit int, duration time.Duration) bool {
	rdb := redis.GetClient()
	ctx := context.Background()

	key = fmt.Sprintf("rate_limit:%s", key)
	count, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		return false
	}
	if count == 1 {
		rdb.Expire(ctx, key, duration)
	}
	return count <= int64(limit)
}
