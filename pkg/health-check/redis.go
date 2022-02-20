package health_check

import (
	"context"
	"log"
	internalRedis "service_health/internal/redis"
)

type redisHealthCheckProbe struct {
	client internalRedis.RedisConnection
}

func NewRedisHealthCheckProbe(endpoint, password string) *redisHealthCheckProbe {
	client := internalRedis.NewConnection(endpoint, password)

	return NewRedisHealthCheckProbeWithInterfaces(client)
}

func NewRedisHealthCheckProbeWithInterfaces(client internalRedis.RedisConnection) *redisHealthCheckProbe {
	return &redisHealthCheckProbe{
		client: client,
	}
}

func (r redisHealthCheckProbe) Name() string {
	return "Redis"
}

func (r *redisHealthCheckProbe) Check(ctx context.Context) error {
	log.Println("checking redis")
	_, err := r.client.Do(ctx, "role").Result()
	return err
}
