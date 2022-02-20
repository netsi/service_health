package handler_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"service_health/cmd/service/handler"
	healthCheck "service_health/pkg/health-check"
	healthCheckMocks "service_health/pkg/health-check/mocks"
	"testing"
	"time"
)

func Test_healthCheckHandler_Check(t *testing.T) {
	tests := []struct {
		name         string
		probes       []healthCheck.HealthCheckProbe
		wantStatus   int
		wantResponse string
	}{
		{
			name:         "no probes configured",
			probes:       nil,
			wantStatus:   http.StatusOK,
			wantResponse: `{"status":"healthy"}`,
		},
		{
			name: "single probe, success",
			probes: []healthCheck.HealthCheckProbe{
				NewMockedHealthCheckProbe(nil),
			},
			wantStatus:   http.StatusOK,
			wantResponse: `{"status":"healthy"}`,
		},
		{
			name: "multiple probes, success",
			probes: []healthCheck.HealthCheckProbe{
				NewMockedHealthCheckProbe(nil),
				NewMockedHealthCheckProbe(nil),
			},
			wantStatus:   http.StatusOK,
			wantResponse: `{"status":"healthy"}`,
		},
		{
			name: "multiple probes, one fail",
			probes: []healthCheck.HealthCheckProbe{
				NewMockedHealthCheckProbe(nil),
				NewMockedHealthCheckProbe(errors.New("some meaningful error")),
			},
			wantStatus:   http.StatusServiceUnavailable,
			wantResponse: `{"status":"unhealthy","reason":"mock-probe: some meaningful error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/health", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			h := handler.NewHealthCheckHandler(tt.probes, time.Second)
			httpHandler := http.HandlerFunc(h.Check)
			httpHandler.ServeHTTP(rr, req)

			assert.Equal(t, tt.wantStatus, rr.Code)
			body, err := ioutil.ReadAll(rr.Body)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantResponse, string(body))
		})
	}
}

func NewMockedHealthCheckProbe(err error) healthCheck.HealthCheckProbe {
	mockProbe := &healthCheckMocks.HealthCheckProbe{}
	mockProbe.On("Check", mock.Anything).Return(err)
	mockProbe.On("Name").Return("mock-probe")

	return mockProbe
}
