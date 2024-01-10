package belajar_golang_db

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConn(t *testing.T) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/belajar_golang_db")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}