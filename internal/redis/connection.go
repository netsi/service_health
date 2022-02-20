package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

// RedisConnection defines just the API we use for the current implementation to help us with the unit tests.
//go:generate mockery --name RedisConnection
type RedisConnection interface {
	Do(ctx context.Context, args ...interface{}) *redis.Cmd
}

func NewConnection(endpoint, password string) RedisConnection {
	return redis.NewClient(&redis.Options{
		Addr:     endpoint,
		Password: password,
	})
}
