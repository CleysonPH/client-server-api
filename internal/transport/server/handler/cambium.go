package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/CleysonPH/client-server-api/internal/repository"
	"github.com/CleysonPH/client-server-api/internal/service"
)

func GetCambium(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	if from == "" {
		from = "USD"
	}
	to := r.URL.Query().Get("to")
	if to == "" {
		to = "BRL"
	}

	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	cambiumResp, err := service.GetCambium(ctx, from, to)
	if err != nil {
		log.Println(err)
		if ctx.Err() == context.DeadlineExceeded {
			http.Error(w, "timeout to get data", http.StatusGatewayTimeout)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	cambiumData := cambiumResp[from+to]
	cambiumM := &repository.CambiumModel{
		Code:       cambiumData.Code,
		Codein:     cambiumData.Codein,
		Name:       cambiumData.Name,
		High:       cambiumData.High,
		Low:        cambiumData.Low,
		VarBid:     cambiumData.VarBid,
		PctChange:  cambiumData.PctChange,
		Bid:        cambiumData.Bid,
		Ask:        cambiumData.Ask,
		Timestamp:  cambiumData.Timestamp,
		CreateDate: cambiumData.CreateDate,
	}

	ctx, cancel = context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()
	if err := repository.InsertCambium(ctx, cambiumM); err != nil {
		log.Println(err)
		if ctx.Err() == context.DeadlineExceeded {
			http.Error(w, "timeout to insert data", http.StatusGatewayTimeout)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(cambiumData.Bid))
}
