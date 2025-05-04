package repository

import (
	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
	"gorm.io/gorm"
)

type StocksPostgresRepo struct {
	db *gorm.DB
}

func NewStocksPostgresRepo(db *gorm.DB) StocksRepo{
	return &StocksPostgresRepo{
		db: db,
	}
}

func (r *StocksPostgresRepo) CreateStockRecord(stock domain.StockModel) error {
	return r.db.Create(&stock).Error
}


func (r *StocksPostgresRepo) GetRecordsCount() (int64, error){
	var count int64
	err := r.db.Model(&domain.StockModel{}).Count(&count).Error
	return count, err
}

func (r *StocksPostgresRepo) GetRatings()([]domain.ParamValueModel, error){
	var ratings []domain.ParamValueModel
	err := r.db.Model(&domain.StockModel{}).
		Select("rating_from as Value, COUNT(*) as Count").
		Group("rating_from").
		Scan(&ratings).Error
	return ratings, err
}


func (r *StocksPostgresRepo) GetActions()([]domain.ParamValueModel, error){
	var ratings []domain.ParamValueModel
	err := r.db.Model(&domain.StockModel{}).
		Select("action as Value, COUNT(*) as Count").
		Group("action").
		Scan(&ratings).Error
	return ratings, err
}