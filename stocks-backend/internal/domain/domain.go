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
	ErrUnableToLiveConnection = errors.New("unable to create live price connection")
	ErrInvalidStockId = errors.New("invalid stock id")
)

type StocksService interface {
	PopulateDatabase(allOrNothing int) error
	GetStocks(filter GetStocksFilter) (*StocksReturnDTO, error)
	GetActions() ([]ActionDTO, error)
	GetRatings() ([]RatingDTO, error)
	GetStockFullData(stockId uuid.UUID) (*StockDataDTO, error)
	GetRecommendations()([]RecommendationDTO, error)
	SuscribeStockPrice(stockId uuid.UUID)(chan PriceUpdateDTO, error)
	UnsuscribeFromStock(stockId uuid.UUID, channel chan PriceUpdateDTO)
}

type ExternalApiService interface{
	GetCompanyProfile(symbol string)(*CompanyProfileDTO, error)
	GetLatestNews(symbol string)([]NewsDTO, error)
	GetStockSentiment(symbol string)(*InternalSentimentDTO, error)
	StartLiveConnection() (chan string, chan StockPriceUpdate, error)
	CloseLiveConnection()
}