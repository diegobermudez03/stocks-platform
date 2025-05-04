package repository

import "github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"

type StocksRepo interface {
	CreateStockRecord(stock domain.StockModel) error 
	GetRecordsCount() (int64, error)
	GetRatings()([]domain.ParamValueModel, error)
	GetActions()([]domain.ParamValueModel, error)
}
