package repository

import (
	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
	"github.com/google/uuid"
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

func (r *StocksPostgresRepo) GetStocks(filter domain.GetStocksFilterModel)([]domain.StockModel, error){
	var stocks []domain.StockModel
	query:= r.getFilteredQuery(filter)
	query.Offset(filter.Page * filter.Size).Limit(filter.Size)
	query.Limit(filter.Size)
	err := query.Find(&stocks).Error
	return stocks, err
}

func (r *StocksPostgresRepo) GetCountWithFilter(filter domain.GetStocksFilterModel)(int64, error){
	query := r.getFilteredQuery(filter)
	var count int64
	err := query.Count(&count).Error
	return count, err
}


func ( r *StocksPostgresRepo) getFilteredQuery(filter domain.GetStocksFilterModel) *gorm.DB{
	query := r.db.Model(&domain.StockModel{})
	if filter.TextSearch != ""{
		query = query.Where("ticker ILIKE ? OR company ILIKE ?", "%"+filter.TextSearch+"%", "%"+filter.TextSearch+"%")
	}
	if len(filter.RatingFromList) > 0{
		query = query.Where("rating_from IN ?", filter.RatingFromList)
	}
	if len(filter.RatingToList) > 0{
		query = query.Where("rating_to IN ?", filter.RatingToList)
	}
	if filter.TimeStart != nil{
		query = query.Where("time >= ?", *filter.TimeStart)
	}
	if filter.TimeEnd != nil{
		query = query.Where("time <= ?", *filter.TimeEnd)
	}
	if len(filter.ActionList) > 0{
		query = query.Where("action IN ?", filter.ActionList)
	}
	if filter.TargetStart != nil && filter.TargetEnd != nil{
		if *filter.TargetStart > *filter.TargetEnd{
			query = query.Where("target_from >= ?", *filter.TargetStart)
			query = query.Where("target_to <= ?", *filter.TargetEnd)
		}else{
			query = query.Where("target_from <= ?", *filter.TargetStart)
			query = query.Where("target_to >= ?", *filter.TargetEnd)
		}
	}
	if filter.Sort != ""{
		query = query.Order(filter.Sort)
	}
	return query
}


func (r *StocksPostgresRepo) GetStockById(id uuid.UUID) (*domain.StockModel, error){
	stockModel := new(domain.StockModel)
	err := r.db.Where("id = ?", id).First(stockModel).Error
	return stockModel, err
}


func (r *StocksPostgresRepo) GetBasicTopStocks() ([]domain.StockModel, error){
	var stocks []domain.StockModel
	err := r.db.Model(&domain.StockModel{}).
		Order("((target_to - target_from)/target_from*100) DESC"). 
		Limit(30).
		Find(&stocks).Error
	return stocks, err
}