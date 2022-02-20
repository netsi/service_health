package health_check_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"service_health/internal/mysql/mocks"
	healthCheck "service_health/pkg/health-check"
	"testing"
)

func Test_mysqlHealthCheck_Check(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		pingErr error
		wantErr bool
	}{
		{name: "ping returns error", pingErr: errors.New("connection failure"), wantErr: true},
		{name: "success", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			connectionMock := &mocks.MySQLConnection{}
			connectionMock.On("PingContext", mock.Anything).Return(tt.pingErr)

			h := healthCheck.NewMySQLHealthCheckProbeWithInterfaces(connectionMock)
			if err := h.Check(ctx); (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlHealthCheck_Name(t *testing.T) {
	want := "MySQL"
	h := healthCheck.NewMySQLHealthCheckProbeWithInterfaces(nil)
	if got := h.Name(); got != want {
		t.Errorf("Name() = %v, want %v", got, want)
	}
}
