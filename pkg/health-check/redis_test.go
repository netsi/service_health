package health_check_test

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/mock"
	"service_health/internal/redis/mocks"
	healthCheck "service_health/pkg/health-check"
	"testing"
)

func Test_redisHealthCheck_Check(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name     string
		redisErr error
		wantErr  bool
	}{
		{name: "returns error", redisErr: errors.New("connection failure"), wantErr: true},
		{name: "success", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			connectionMock := &mocks.RedisConnection{}
			connectionMock.On("Do", mock.Anything, "role").Return(redis.NewCmdResult("", tt.redisErr))

			h := healthCheck.NewRedisHealthCheckProbeWithInterfaces(connectionMock)
			if err := h.Check(ctx); (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisHealthCheck_Name(t *testing.T) {
	want := "Redis"
	h := healthCheck.NewRedisHealthCheckProbeWithInterfaces(nil)
	if got := h.Name(); got != want {
		t.Errorf("Name() = %v, want %v", got, want)
	}
}
