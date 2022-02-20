package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	healthCheck "service_health/pkg/health-check"
	"time"
)

const (
	ContentTypeHeader = "Content-Type"
	JsonContentType   = "application/json"
)

type healthCheckHandler struct {
	probes                 []healthCheck.HealthCheckProbe
	maxHealthCheckDuration time.Duration
}

// NewHealthCheckHandler factory.
func NewHealthCheckHandler(probes []healthCheck.HealthCheckProbe, maxHealthCheckDuration time.Duration) *healthCheckHandler {
	return &healthCheckHandler{
		probes:                 probes,
		maxHealthCheckDuration: maxHealthCheckDuration,
	}
}

// Check all the health-check probes and returns a HealthResponse. If all the services checked are running returns
// HealthyStatus, if any of the services fails the health check or the request times out returns UnhealthyStatus.
func (h *healthCheckHandler) Check(w http.ResponseWriter, r *http.Request) {
	log.Println("handling health-check")
	ctx, cancel := context.WithTimeout(r.Context(), h.maxHealthCheckDuration)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	for i := 0; i < len(h.probes); i++ {
		probe := h.probes[i]
		g.Go(func() error {
			err := probe.Check(ctx)
			if err != nil {
				return fmt.Errorf("%s: %s", probe.Name(), err.Error())
			}

			return nil
		})
	}

	response := &HealthResponse{
		Status: HealthyStatus,
	}
	statusCode := http.StatusOK

	err := g.Wait()
	if err != nil {
		response = &HealthResponse{
			Status: UnhealthyStatus,
			Reason: err.Error(),
		}
		statusCode = http.StatusServiceUnavailable
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("failed to write response with error: %s", err.Error())
	}

	w.Header().Set(ContentTypeHeader, JsonContentType)
	w.WriteHeader(statusCode)
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatalf("failed to write response with error: %s", err.Error())
	}
}
