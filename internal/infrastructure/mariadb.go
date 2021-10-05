package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

func ConnectToMariaDB() (*sql.DB, error) {
	host := getEnv("API_DB_HOST", "mariadb")
	port := getEnv("API_DB_PORT", "3306")
	dbName := getEnv("API_DB_NAME", "fiber_dmca")
	userName := getEnv("API_DB_USER", "root")
	pwd := getEnv("API_DB_PWD", "")

	return ConnectToMariaDBParams(host, port, userName, pwd, dbName)
}

func getEnv(name string, defaultVal string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		return defaultVal
	}

	return val
}

// Connect to MariaDB
func ConnectToMariaDBParams(host string, port string, userName string, pwd string, dbName string) (*sql.DB, error) {
	// db, err := sql.Open("mysql", "root:12345@tcp(mariadb:3306)/fiber_dmca")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, pwd, host, port, dbName))
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
