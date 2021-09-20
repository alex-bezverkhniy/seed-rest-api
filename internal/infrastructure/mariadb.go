package infrastructure

import (
	"database/sql"
	"time"
)

// Connect to MariaDB
func ConnectToMariaDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(mariadb:3306)/fiber_dmca")
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
