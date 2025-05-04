package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrMakingHttpCall = errors.New("error making http call")
	ErrInvalidPayloadInHttpCall = errors.New("invalid payload in http call")
	ErrInternalErrorWritingToDb = errors.New("internal error writing to db")
	ErrInternalError = errors.New("An error ocurred, try again later")
)

type StocksService interface {
	PopulateDatabase() error
	GetStocks(filter GetStocksFilter) ([]StockDTO, error)
	GetActions() ([]ActionDTO, error)
	GetRatings() ([]RatingDTO, error)
	GetStockFullData(stockId uuid.UUID) (*StockDataDTO, error)
	GetRecommendations()([]RecommendationDTO, error)
}