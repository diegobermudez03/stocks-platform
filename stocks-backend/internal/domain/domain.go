package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrMakingHttpCall = errors.New("error making http call")
	ErrInvalidPayloadInHttpCall = errors.New("invalid payload in http call")
	ErrInternalErrorWritingToDb = errors.New("internal error writing to db")
	ErrInternalError = errors.New("an error ocurred, try again later")
)

type StocksService interface {
	PopulateDatabase() error
	GetStocks(filter GetStocksFilter) (*StocksReturnDTO, error)
	GetActions() ([]ActionDTO, error)
	GetRatings() ([]RatingDTO, error)
	GetStockFullData(stockId uuid.UUID) (*StockDataDTO, error)
	GetRecommendations()([]RecommendationDTO, error)
}

type ExternalApiService interface{
	GetCompanyProfile(symbol string)(*CompanyProfileDTO, error)
	GetLatestNews(symbol string)([]NewsDTO, error)
	GetStockSentiment(symbol string)(*InternalSentimentDTO, error)
}