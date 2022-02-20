package handler

// HealthStatus type.
type HealthStatus string

const (
	HealthyStatus   HealthStatus = "healthy"
	UnhealthyStatus HealthStatus = "unhealthy"
)

// HealthResponse represents the object returned by the healthCheckHandler.
type HealthResponse struct {
	Status HealthStatus `json:"status"`
	Reason string       `json:"reason,omitempty"`
}
