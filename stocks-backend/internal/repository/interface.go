package repository

import (
	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
	"github.com/google/uuid"
)

type StocksRepo interface {
	CreateStockRecord(stock domain.StockModel) error 
	GetRecordsCount() (int64, error)
	GetRatings()([]domain.ParamValueModel, error)
	GetActions()([]domain.ParamValueModel, error)
	GetStocks(filter domain.GetStocksFilterModel)([]domain.StockModel, error)
	GetStockById(id uuid.UUID) (*domain.StockModel, error)
}
