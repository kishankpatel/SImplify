package db

import (
	"database/sql"
	"fmt"
)

type Env struct {
	db *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "francium"
	password = "francium"
	dbname   = "simplify"
)

// InitDb - establish database connection
func InitDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("Successfully connected!")
	return db
}
