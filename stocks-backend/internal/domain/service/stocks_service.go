package service

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/repository"
	"github.com/google/uuid"
)

type StocksServiceImpl struct {
	repo repository.StocksRepo
	apiUrl string
	apiToken string
	externalAPIUrl string
	externalAPIKey string
}

func NewStocksService(repo repository.StocksRepo, apiUrl, apiToken, externalAPIUrl, externalAPIKey string) domain.StocksService{
	return &StocksServiceImpl{
		repo : repo,
		apiUrl: apiUrl,
		apiToken: apiToken,
		externalAPIUrl: externalAPIUrl,
		externalAPIKey: externalAPIKey,
	}
}

/*
	PopulateDatabase is the method that takes care of making the HTTP call to the API if the db isnt
	populated yet, and populates it with that data
*/
func(s *StocksServiceImpl) PopulateDatabase() error{
	count, err := s.repo.GetRecordsCount()
	if err != nil{
		return err 
	}
	//if count is equal to 0, means that db isnt populated yet, so we try to populate it
	if count == 0{
		if err := s.populateWithAPI(); err != nil{
			return err 
		}
	}
	return nil
}

/*
	Main method, to retrieve the stocks based on the filters
*/
func (s *StocksServiceImpl) GetStocks(filter domain.GetStocksFilter) ([]domain.StockDTO, error){
	if filter.Page > 0{
		filter.Page -= 1
	}
	if filter.Size == 0{
		filter.Size = 12
	}
	var startTime *time.Time
	if filter.TimeStart != ""{
		if parsedTime, err := time.Parse(time.RFC3339Nano, filter.TimeStart); err ==nil{
			startTime = &parsedTime
		}
	}
	var endTime *time.Time
	if filter.TimeStart != ""{
		if parsedTime, err := time.Parse(time.RFC3339Nano, filter.TimeEnd); err ==nil{
			endTime = &parsedTime
		}
	}
	var targetStart *float64
	if filter.TargetStart != 0{
		targetStart = &filter.TargetStart
	}
	var targetEnd *float64
	if filter.TargetEnd != 0{
		targetEnd = &filter.TargetEnd
	}
	stocks, err := s.repo.GetStocks(domain.GetStocksFilterModel{
		Page: filter.Page,
		Size: filter.Size,
		TextSearch: filter.TextSearch,
		RatingFromList: filter.RatingFromList,
		RatingToList: filter.RatingToList,
		TimeStart: startTime,
		TimeEnd: endTime,
		Sort: domain.SortMap[filter.Sort],
		ActionList: filter.ActionList,
		TargetStart: targetStart,
		TargetEnd: targetEnd,
	})
	if err != nil{
		return nil, err
	}
	stocksDTO := make([]domain.StockDTO, len(stocks))
	for i, stock := range stocks{
		stocksDTO[i] = *s.stockModelToDTO(&stock)
	}
	return stocksDTO, nil
}


/*
	Method to retrieve all possible actions
*/
func (s *StocksServiceImpl) GetActions() ([]domain.ActionDTO, error){
	actions, err := s.repo.GetActions()
	if err != nil{
		return nil, err
	}
	actionsDto := make([]domain.ActionDTO, len(actions))
	for i, action := range actions{
		actionsDto[i] = domain.ActionDTO{
			Action: action.Value,
			Count: action.Count,
		}
	}
	return actionsDto, nil
}


/*
	Method to retrieve all possible ratings
*/
func (s *StocksServiceImpl) GetRatings() ([]domain.RatingDTO, error){
	ratings, err := s.repo.GetRatings()
	if err != nil{
		return nil, err
	}
	ratingsDTO := make([]domain.RatingDTO, len(ratings))
	for i, rating := range ratings{
		ratingsDTO[i] = domain.RatingDTO{
			Rating: rating.Value,
			Count: rating.Count,
		}
	}
	return ratingsDTO, nil
}


/*
	Method to retrieve the full information of the stock, which means, the company profile and latest news
*/
func (s *StocksServiceImpl)  GetStockFullData(stockId uuid.UUID) (*domain.StockDataDTO, error){
	stock, err := s.repo.GetStockById(stockId)
	if err != nil{
		return nil, err 
	}
	//get the company profile
	stockData, err := s.getCompanyProfile(stock)
	if err != nil{
		return nil, err
	}
	//get company latest news
	currentDate := time.Now().Format("2006-01-02")
	response2, err := http.Get(s.externalAPIUrl + "/company-news?symbol=" + stock.Ticker + "&from=2025-01-01" + "&to=" +currentDate + "&token=" + s.externalAPIKey)
	if err != nil{
		return stockData, nil
	}
	defer response2.Body.Close()
	payload2, err := io.ReadAll(response2.Body)
	if err != nil{
		return stockData, nil
	}
	news := []InternalNewsDTO{}
	if err := json.Unmarshal(payload2, &news); err != nil{
		return stockData, nil
	}
	newsDTO := make([]domain.NewsDTO, len(news))
	for i, news := range news{
		newsDTO[i] = domain.NewsDTO{
			Date: time.Unix(news.Datetime, 0),
			Headline: news.Headline,
			Image: news.Image,
			Source: news.Source,
			Summary: news.Summary,
		}
	}
	stockData.News = newsDTO
	return stockData, nil
}


/*
	Internal method to convert models to DTO
*/
func (s *StocksServiceImpl) stockModelToDTO(model *domain.StockModel) *domain.StockDTO{
	return &domain.StockDTO{
		ID: model.ID,
		Ticker: model.Ticker,
		TargetFrom: model.TargetFrom,
		TargetTo: model.TargetTo,
		Company: model.Company,
		Action: model.Action,
		Brokerage: model.Brokerage,
		RatingFrom: model.RatingFrom,
		RatingTo: model.RatingTo,
		Time: model.Time,
	}
}