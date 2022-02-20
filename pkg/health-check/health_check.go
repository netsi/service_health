package health_check

import "context"

// HealthCheckProbe defines functions required by the HealthCheckProbe probe implementations
//go:generate mockery --name HealthCheckProbe
type HealthCheckProbe interface {
	Name() string
	Check(ctx context.Context) error
}

// DefaultHealthCheckProbes returns a list of HealthCheckProbe items.
func DefaultHealthCheckProbes(httpEndpoint string) []HealthCheckProbe {
	probes := []HealthCheckProbe{
		NewMySQLHealthCheckProbe(MySQLSettings{
			Host:     "mysql",
			User:     "root",
			Port:     3306,
			Password: "password",
		}),
		NewRedisHealthCheckProbe("redis-leader:6379", ""),
	}
	if httpEndpoint != "" {
		probes = append(probes, NewHttpHealthCheckProbe(httpEndpoint))
	}

	return probes
}
