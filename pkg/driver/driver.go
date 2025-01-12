package driver

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func OpenDB(user, password, addr, database string) (*sql.DB, error) {
	// Connect to the ClickHouse server
	db := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{addr},
		Auth: clickhouse.Auth{
			Database: database,
			Username: user,
			Password: password,
		},
		DialTimeout: time.Second,
		Compression: &clickhouse.Compression{Method: clickhouse.CompressionLZ4},
		Settings:    clickhouse.Settings{"max_execution_time": 60},
	})
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}

// Auth: clickhouse.Auth{
// 	Database: "default", // Default database
// 	Username: "default", // Default username
// 	Password: "",        // Default password (if any)
// },
