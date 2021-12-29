package db

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dataSource := "root:root@tcp(localhost:3306)/productos"
	txdb.Register("txdb", "mysql", dataSource)
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", "identificador")

	if err == nil {
		return db, db.Ping()
	}

	return db, err
}
