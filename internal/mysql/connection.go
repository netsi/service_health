package mysql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// MySQLConnection defines just the API we use for the current implementation to help us with the unit tests
//go:generate mockery --name MySQLConnection
type MySQLConnection interface {
	Close() error
	PingContext(ctx context.Context) error
}

func NewConnection(dsn string) (MySQLConnection, error) {
	return sql.Open("mysql", dsn)
}
