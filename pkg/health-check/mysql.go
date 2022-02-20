package health_check

import (
	"context"
	"fmt"
	"log"
	"service_health/internal/mysql"
)

type MySQLSettings struct {
	Host     string
	User     string
	Port     int
	Password string
}

type mysqlHealthCheckProbe struct {
	db mysql.MySQLConnection
}

// NewMySQLHealthCheckProbe returns an instance of *mysqlHealthCheckProbe with the given settings.
func NewMySQLHealthCheckProbe(settings MySQLSettings) *mysqlHealthCheckProbe {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/",
		settings.User,
		settings.Password,
		settings.Host,
		settings.Port,
	)
	db, err := mysql.NewConnection(dsn)
	if err != nil {
		log.Fatalf("failed to initialize the MySQL connection with error %s", err.Error())
	}

	return NewMySQLHealthCheckProbeWithInterfaces(db)
}

// NewMySQLHealthCheckProbeWithInterfaces returns an instance of *mysqlHealthCheckProbe with the given interfaces.
func NewMySQLHealthCheckProbeWithInterfaces(connection mysql.MySQLConnection) *mysqlHealthCheckProbe {
	return &mysqlHealthCheckProbe{
		db: connection,
	}
}

// Name returns the name of the probe.
func (m mysqlHealthCheckProbe) Name() string {
	return "MySQL"
}

// Check if the service has access to the MySQL instance.
func (m *mysqlHealthCheckProbe) Check(ctx context.Context) error {
	log.Println("checking mysql connection")

	err := m.db.PingContext(ctx)
	if err != nil {
		log.Printf("Couldn't connect to mysql server: %s", err.Error())
		return err
	}

	return nil
}
