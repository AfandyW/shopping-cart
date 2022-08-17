package app

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func NewDB(DBPort, DBHost, DBUser, DBPassword, DBName string) (*sql.DB, error) {
	// urlDatabase := "postgres://mapple@localhost:5432/go_rest?sslmode=disable"
	psqlconn := fmt.Sprintf(`dbname=%s user=%s password=%s host=%s port=%s  sslmode=disable`, DBName, DBUser, DBPassword, DBHost, DBPort)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, fmt.Errorf("error Open db : %s", err.Error())
	}

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("error setup db : %s", err.Error())
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db, nil
}
