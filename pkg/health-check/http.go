package health_check

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type httpHealthCheckProbe struct {
	endpoint string
}

// NewHttpHealthCheckProbe returns a new instance of *httpHealthCheckProbe with the endpoint given.
func NewHttpHealthCheckProbe(endpoint string) *httpHealthCheckProbe {
	return &httpHealthCheckProbe{
		endpoint: endpoint,
	}
}

// Name returns the name of the HealthCheckProbe implementation.
func (h httpHealthCheckProbe) Name() string {
	return "Http"
}

// Check the health of the endpoint.
func (h *httpHealthCheckProbe) Check(ctx context.Context) error {
	log.Printf("checking: %s\n", h.endpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, h.endpoint, nil)
	if err != nil {
		log.Printf("failed to create request with error: %s\n", err.Error())
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("failed to create request with error: %s\n", err.Error())
		return err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("health check returned status code %d\n", resp.StatusCode)
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	return nil
}
