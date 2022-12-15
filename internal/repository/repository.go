package repository

import (
	"context"

	"github.com/google/uuid"
)

type CambiumModel struct {
	ID         string
	Code       string
	Codein     string
	Name       string
	High       string
	Low        string
	VarBid     string
	PctChange  string
	Bid        string
	Ask        string
	Timestamp  string
	CreateDate string
}

func InsertCambium(ctx context.Context, cm *CambiumModel) error {
	query := `
		INSERT INTO cambium (
			id,
			code,
			codein,
			name,
			high,
			low,
			var_bid,
			pct_change,
			bid,
			ask,
			timestamp,
			create_date
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	id := uuid.New().String()
	_, err := db.ExecContext(
		ctx,
		query,
		id,
		cm.Code,
		cm.Codein,
		cm.Name,
		cm.High,
		cm.Low,
		cm.VarBid,
		cm.PctChange,
		cm.Bid,
		cm.Ask,
		cm.Timestamp,
		cm.CreateDate,
	)
	if err != nil {
		return err
	}
	cm.ID = id
	return nil
}
