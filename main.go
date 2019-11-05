package main

import (
	"example-db/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server.StartServer()
}
