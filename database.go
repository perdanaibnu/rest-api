package belajar_golang_db

import (
	"database/sql"
	"time"
)

func GetConn() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/belajar_golang_db")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
