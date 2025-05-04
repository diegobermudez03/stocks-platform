package domain

import "github.com/google/uuid"

type StockModel struct {
	ID     uuid.UUID
	Ticker string
}