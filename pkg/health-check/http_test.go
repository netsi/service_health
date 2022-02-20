package health_check_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	healthCheck "service_health/pkg/health-check"
	"testing"
)

func Test_httpHealthCheck_Check(t *testing.T) {

	tests := []struct {
		name            string
		invalidEndpoint bool
		responseCode    int
		contextCanceled bool
		wantErr         bool
	}{
		{name: "invalid endpoint", invalidEndpoint: true, wantErr: true},
		{name: "fail", responseCode: http.StatusBadGateway, wantErr: true},
		{name: "canceled context", contextCanceled: true, wantErr: true},
		{name: "success", responseCode: http.StatusOK, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancelFunc := context.WithCancel(context.Background())
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.contextCanceled {
					cancelFunc()
				}
				w.WriteHeader(tt.responseCode)
			}))
			defer ts.Close()

			url := ts.URL
			if tt.invalidEndpoint {
				url = "http://:localhost"
			}

			h := healthCheck.NewHttpHealthCheckProbe(url)
			if err := h.Check(ctx); (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_httpHealthCheck_Name(t *testing.T) {
	want := "Http"
	h := healthCheck.NewHttpHealthCheckProbe("")
	if got := h.Name(); got != want {
		t.Errorf("Name() = %v, want %v", got, want)
	}
}
