package main

import (
	"flag"

	"github.com/CleysonPH/client-server-api/internal/transport/server"
)

func main() {
	addr := flag.String("addr", ":8080", "server address")
	dsn := flag.String("dsn", "file:db.sqlite3?cache=shared", "database connection string")
	flag.Parse()
	server.Run(*dsn, *addr)
}
