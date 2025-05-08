package stocks

import (
	"time"

	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/repository"
	"github.com/google/uuid"
)

type StocksServiceImpl struct {
	repo repository.StocksRepo
	apiUrl string
	apiToken string
	recommendationsCache []domain.RecommendationDTO
	lastCacheSaved time.Time
	externalAPI 	domain.ExternalApiService
}

func NewStocksService(repo repository.StocksRepo, apiUrl, apiToken string, externalAPI domain.ExternalApiService) domain.StocksService{
	return &StocksServiceImpl{
		repo : repo,
		apiUrl: apiUrl,
		apiToken: apiToken,
		externalAPI: externalAPI,
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
func (s *StocksServiceImpl) GetStocks(filter domain.GetStocksFilter) (*domain.StocksReturnDTO, error){
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
	filterModel := domain.GetStocksFilterModel{
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
	}
	stocks, err := s.repo.GetStocks(filterModel)
	if err != nil{
		return nil, domain.ErrInternalError
	}
	stocksDTO := make([]domain.StockDTO, len(stocks))
	for i, stock := range stocks{
		stocksDTO[i] = *s.stockModelToDTO(&stock)
	}
	count, err := s.repo.GetCountWithFilter(filterModel)
	if err != nil{
		return nil, domain.ErrInternalError
	}

	return &domain.StocksReturnDTO{
		Stocks: stocksDTO,
		Count: count,
	}, nil
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
	companyProfile, err := s.externalAPI.GetCompanyProfile(stock.Ticker)
	if err != nil{
		return nil, err
	}
	//construct the preview payload, since if the news fail, we will still return succesfully
	stockData := domain.StockDataDTO{
		Stock: *s.stockModelToDTO(stock),
		CompanyProfile: *companyProfile,
		News: []domain.NewsDTO{},
	}
	//get company latest news
	news, err := s.externalAPI.GetLatestNews(stock.Ticker)
	//even if we had an error with the news, we will still return the stock info with an empty array
	if err != nil{
		return &stockData, nil
	}
	stockData.News = news
	return &stockData, nil
}


/*
	Method to get the recommendations, it add sthe cache check layer
*/
func (s *StocksServiceImpl) GetRecommendations()([]domain.RecommendationDTO, error){
	//if it has passed 10 minutes since last cache, we re fetch and re store
	if s.lastCacheSaved.Before(time.Now().Add(time.Minute*-10)){
		recommendations, err := s.getRecommendationsInternal()
		if err != nil{
			return nil, err
		}
		s.recommendationsCache = recommendations
		s.lastCacheSaved = time.Now()
		return recommendations, nil
	}
	return s.recommendationsCache, nil
}


/*
	Internal method to convert models to DTO
*/
func (s *StocksServiceImpl) stockModelToDTO(model *domain.StockModel) *domain.StockDTO{
	variationPercentage := (model.TargetTo-model.TargetFrom)/model.TargetFrom*100
	return &domain.StockDTO{
		ID: model.ID,
		Ticker: model.Ticker,
		TargetFrom: model.TargetFrom,
		TargetTo: model.TargetTo,
		Company: model.Company,
		Action: model.Action,
		Brokerage: model.Brokerage,
		RatingFrom: model.RatingFrom,
		Percentage: variationPercentage,
		RatingTo: model.RatingTo,
		Time: model.Time,
	}
}