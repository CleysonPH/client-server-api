package main

import (
	"flag"

	"github.com/CleysonPH/client-server-api/internal/transport/client"
)

func main() {
	addr := flag.String("addr", "http://localhost:8080", "server address")
	filename := flag.String("filename", "cotacao.txt", "file name")
	flag.Parse()
	client.Run(*addr, *filename)
}
