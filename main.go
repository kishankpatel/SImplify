package main

import (
	"github.com/kishankpatel/simplify/db"
	"github.com/kishankpatel/simplify/routes"

	_ "github.com/lib/pq"
)

func main() {
	db.InitDb()
	routes.Handler()
}
