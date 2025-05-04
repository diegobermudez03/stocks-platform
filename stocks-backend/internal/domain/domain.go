package domain

import "errors"

var (
	ErrMakingHttpCall = errors.New("error making http call")
	ErrInvalidPayloadInHttpCall = errors.New("invalid payload in http call")
	ErrInternalErrorWritingToDb = errors.New("internal error writing to db")
)

type StocksService interface {
	PopulateDatabase() error
	GetStocks(filter GetStocksFilter) ([]StockDTO, error)
	GetActions() ([]ActionDTO, error)
	GetRatings() ([]RatingDTO, error)
}