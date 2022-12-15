package server

import (
	"net/http"

	"github.com/CleysonPH/client-server-api/internal/repository"
	"github.com/CleysonPH/client-server-api/internal/transport/server/handler"
)

func Run() {
	repository.InitDb("file:db.sqlite3?cache=shared")
	defer repository.CloseDb()

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", handler.GetCambium)
	http.ListenAndServe(":8081", mux)
}
