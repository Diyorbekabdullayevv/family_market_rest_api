package dbconnect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {

	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	conStr := fmt.Sprintf("root:%s@tcp(127.0.0.1%s)/%s", dbPassword, dbPort, dbName)
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
