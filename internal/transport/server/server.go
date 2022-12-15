package server

import (
	"net/http"

	"github.com/CleysonPH/client-server-api/internal/repository"
	"github.com/CleysonPH/client-server-api/internal/transport/server/handler"
)

func Run(dsn, addr string) {
	repository.InitDb(dsn)
	defer repository.CloseDb()

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", handler.GetCambium)
	http.ListenAndServe(addr, mux)
}
