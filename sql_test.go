package belajar_golang_db

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(m *testing.M) {
	db := GetConn()
	defer db.Close()

	db_ctx := context.Background()

	_, err := db.ExecContext(db_ctx, "INSERT INTO customer(id, name) VALUES('eko', 'eko')")

	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
