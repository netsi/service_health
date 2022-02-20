package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service_health/cmd/service/handler"
	healthCheck "service_health/pkg/health-check"
	"syscall"
	"time"
)

const (
	httpEndpointEnvVariableKey    = "HTTP_ENDPOINT_CHECK"
	httpServerPort                = 8080
	defaultMaxHealthCheckDuration = 5 * time.Second
	defaultReadTimeout            = 60 * time.Second
	defaultWriteTimeout           = 30 * time.Second
)

func main() {
	httpEndpointCheck := os.Getenv(httpEndpointEnvVariableKey)

	healthCheckHandler := handler.NewHealthCheckHandler(
		healthCheck.DefaultHealthCheckProbes(httpEndpointCheck),
		defaultMaxHealthCheckDuration,
	)

	http.HandleFunc("/health", healthCheckHandler.Check)
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", httpServerPort),
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatalf("failed to start the http server on port %d with error %s.\n", httpServerPort, err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("finished execution.")
}
