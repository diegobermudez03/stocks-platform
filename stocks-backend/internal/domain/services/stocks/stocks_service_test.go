package stocks

import (
	"errors"
	"testing"
	"time"

	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
	mocks "github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain/services/stocks/stocksmock"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
	Tests for GET STOCKS domain method, checks:
	- 	Success case with filters
	-	Success case with no filter
	-	Error handling from getstocks repo
	-	Error handling from getcount repo
*/
func TestStocksServiceImpl_GetStocks(t *testing.T) {
	stockTime := time.Date(2023, 10, 26, 10, 0, 0, 0, time.UTC)
	stockID1 := uuid.New()
	stockID2 := uuid.New()

	/*
		Test for the Get Stocks method, it checks the corrrect parsing of filter (specifically the order attribute)
		As well as the calculation of the variation percentage which is calculated
	*/
	t.Run("success - full filter", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "", externalAPIMock)

		inputFilter := domain.GetStocksFilter{
			Page:           1,
			Size:           10,
			TextSearch:     "Ap",
			RatingFromList: nil,
			RatingToList:   nil,
			TimeStart:      "2023-01-01",
			TimeEnd:        "2025-01-31",
			Sort:           "CLOSEST_DATE", 
			ActionList:     nil,
			TargetStart:    150.0,
			TargetEnd:      200.0,
		}

		expectedRepoStocks := []domain.StockModel{
			{ID: stockID1, Ticker: "AAPL", TargetFrom: 160, TargetTo: 176, Time: stockTime, Company: "Apple Inc."}, 
			{ID: stockID2, Ticker: "MSFT", TargetFrom: 169, TargetTo: 200, Time: stockTime, Company: "Api"},  
		}
		expectedCount := int64(2)
		repoMock.On("GetStocks",  mock.MatchedBy(func(f domain.GetStocksFilterModel) bool {
			return f.TextSearch == "Ap" &&
				   f.Sort == "time DESC" &&
				   f.Size == 10 &&
				   f.Page == 0
		})).Return(expectedRepoStocks, nil).Once()
		repoMock.On("GetCountWithFilter", mock.MatchedBy(func(f domain.GetStocksFilterModel) bool {
			return f.TextSearch == "Ap" &&
				   f.Sort == "time DESC" &&
				   f.Size == 10 &&
				   f.Page == 0
		})).Return(expectedCount, nil).Once()

		result, err := service.GetStocks(inputFilter)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedCount, result.Count)
		assert.Len(t, result.Stocks, 2)
		assert.Equal(t, "AAPL", result.Stocks[0].Ticker)
		//validate correct percentgae calculation
		assert.Equal(t, 10.0, result.Stocks[0].Percentage) 

		repoMock.AssertExpectations(t)
	})

	/*
		Test for GetStocks but with no filter
	*/
	t.Run("success - minimal filter and zero TargetFrom for percentage", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "",externalAPIMock)

		inputFilter := domain.GetStocksFilter{
			Page: 1,
			Size: 5,
		}

		expectedRepoStocks := []domain.StockModel{
			{ID: stockID1, Ticker: "ZERO", TargetFrom: 1, TargetTo: 10, Time: stockTime, Company: "Zero Corp"}, 
		}
		expectedCount := int64(1)

		expectedFilterModel := domain.GetStocksFilterModel{
			Page: 0,
			Size: 5,
		}

		repoMock.On("GetStocks", expectedFilterModel).Return(expectedRepoStocks, nil).Once()
		repoMock.On("GetCountWithFilter", expectedFilterModel).Return(expectedCount, nil).Once()

		result, err := service.GetStocks(inputFilter)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedCount, result.Count)
		assert.Len(t, result.Stocks, 1)
		assert.Equal(t, "ZERO", result.Stocks[0].Ticker)
		assert.Equal(t, float64(900), result.Stocks[0].Percentage)
		repoMock.AssertExpectations(t)
	})


	/*
		Test for validate the correct error handling insuide the method GetStocks when the repository fails
	*/
	t.Run("error from repo.GetStocks", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "",externalAPIMock)

		inputFilter := domain.GetStocksFilter{Page: 1, Size: 10}
		repoError := errors.New("database connection error")
		repoMock.On("GetStocks", mock.AnythingOfType("domain.GetStocksFilterModel")).Return(nil, repoError).Once()

		result, err := service.GetStocks(inputFilter)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "an error ocurred, try again later")

		repoMock.AssertExpectations(t)
		//Get count shouldnt have been called
		repoMock.AssertNotCalled(t, "GetCountWithFilter", mock.AnythingOfType("domain.GetStocksFilterModel"))
	})


	/*
		Test for validate the error when the get count returns an error
	*/
	t.Run("error from repo.GetCountWithFilter", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "",externalAPIMock)

		inputFilter := domain.GetStocksFilter{Page: 1, Size: 10}
		expectedRepoStocks := []domain.StockModel{
			{ID: stockID1, Ticker: "ANY", TargetFrom: 50, TargetTo: 55, Time: stockTime},
		}
		repoError := errors.New("error counting records")

		repoMock.On("GetStocks", mock.AnythingOfType("domain.GetStocksFilterModel")).Return(expectedRepoStocks, nil).Once()
		repoMock.On("GetCountWithFilter", mock.AnythingOfType("domain.GetStocksFilterModel")).Return(int64(0), repoError).Once()

		result, err := service.GetStocks(inputFilter)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "an error ocurred, try again later")
		repoMock.AssertExpectations(t)
	})
}


/*
	Tests for GeActions domain method, checks:
	-	Success case
	-	Error case (from repository)
*/
func TestStocksServiceImpl_GetActions(t *testing.T) {
	t.Run("success - get actions", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "", externalAPIMock)

		expectedRepoActions := []domain.ParamValueModel{
			{Value: "Buy", Count: 10},
			{Value: "Sell", Count: 5},
		}

		repoMock.On("GetActions").Return(expectedRepoActions, nil).Once()

		result, err := service.GetActions()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, 2)

		assert.Equal(t, "Buy", result[0].Action)
		assert.Equal(t, int64(10), result[0].Count)
		assert.Equal(t, "Sell", result[1].Action)
		assert.Equal(t, int64(5), result[1].Count)

		repoMock.AssertExpectations(t)
	})

	t.Run("error - get actions from repo", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "", externalAPIMock)

		repoError := errors.New("database error fetching actions")
		repoMock.On("GetActions").Return(nil, repoError).Once()

		result, err := service.GetActions()

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, repoError, err)

		repoMock.AssertExpectations(t)
	})
}


/*
	Tests for GetRatings domain method, checks:
	-	Success case
	-	Error case (from repository)
*/
func TestStocksServiceImpl_GetRatings(t *testing.T) {
	t.Run("success - get ratings", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "", externalAPIMock)

		expectedRepoRatings := []domain.ParamValueModel{
			{Value: "Positive", Count: 15},
			{Value: "Neutral", Count: 8},
		}

		repoMock.On("GetRatings").Return(expectedRepoRatings, nil).Once()

		result, err := service.GetRatings()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, 2)

		assert.Equal(t, "Positive", result[0].Rating)
		assert.Equal(t, int64(15), result[0].Count)
		assert.Equal(t, "Neutral", result[1].Rating)
		assert.Equal(t, int64(8), result[1].Count)

		repoMock.AssertExpectations(t)
	})

	t.Run("error - get ratings from repo", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "", externalAPIMock)

		repoError := errors.New("database error fetching ratings")
		repoMock.On("GetRatings").Return(nil, repoError).Once()

		result, err := service.GetRatings()

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, repoError, err)

		repoMock.AssertExpectations(t)
	})
}


/*
	Tests for GetStockFullData method, it checks:
	-	Success case
	-	Error case due to repository failure
	-	Error case due to Get company profile failure
	-	Success case even with an error with the news fetching
*/
func TestStocksServiceImpl_GetStockFullData(t *testing.T) {
	stockID := uuid.New()
	stockTime := time.Date(2023, 11, 1, 12, 0, 0, 0, time.UTC)

	mockStockModel := &domain.StockModel{
		ID:         stockID,
		Ticker:     "TEST",
		TargetFrom: 100,
		TargetTo:   110,
		Company:    "Test Corp",
		Action:     "Buy",
		Brokerage:  "Test Broker",
		RatingFrom: "Neutral",
		RatingTo:   "Positive",
		Time:       stockTime,
	}

	mockCompanyProfile := &domain.CompanyProfileDTO{
		Country: "US",
		Name:    "Test Corp Profile",
	}

	mockNews := []domain.NewsDTO{
		{Headline: "Test News 1", Date: time.Now()},
	}

	t.Run("success - get full stock data", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "", externalAPIMock)

		externalAPIMock.On("GetStockSentiment", "TEST").Return(&domain.InternalSentimentDTO{}, nil).Once()
		repoMock.On("GetStockById", stockID).Return(mockStockModel, nil).Once()
		externalAPIMock.On("GetCompanyProfile", "TEST").Return(mockCompanyProfile, nil).Once()
		externalAPIMock.On("GetLatestNews", "TEST").Return(mockNews, nil).Once()

		result, err := service.GetStockFullData(stockID)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, stockID, result.Stock.ID)
		assert.Equal(t, "TEST", result.Stock.Ticker)
		assert.Equal(t, 10.0, result.Stock.Percentage) // (110-100)/100 * 100
		assert.Equal(t, "Test Corp Profile", result.CompanyProfile.Name)
		assert.Len(t, result.News, 1)
		assert.Equal(t, "Test News 1", result.News[0].Headline)

		repoMock.AssertExpectations(t)
		externalAPIMock.AssertExpectations(t)
	})

	t.Run("error - repo GetStockById fails", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "", externalAPIMock)

		repoError := errors.New("database error fetching stock by ID")
		repoMock.On("GetStockById", stockID).Return(nil, repoError).Once()

		result, err := service.GetStockFullData(stockID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, repoError, err)

		repoMock.AssertExpectations(t)
		externalAPIMock.AssertNotCalled(t, "GetCompanyProfile", mock.Anything)
		externalAPIMock.AssertNotCalled(t, "GetLatestNews", mock.Anything)
	})

	t.Run("error - externalAPI GetCompanyProfile fails", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "", externalAPIMock)

		repoMock.On("GetStockById", stockID).Return(mockStockModel, nil).Once()
		apiError := errors.New("external API error fetching company profile")
		externalAPIMock.On("GetCompanyProfile", "TEST").Return(nil, apiError).Once()

		result, err := service.GetStockFullData(stockID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, apiError, err)

		repoMock.AssertExpectations(t)
		externalAPIMock.AssertExpectations(t)
		externalAPIMock.AssertNotCalled(t, "GetLatestNews", mock.Anything)
	})

	t.Run("success - externalAPI GetLatestNews fails (should still return partial data)", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "", externalAPIMock)

		repoMock.On("GetStockById", stockID).Return(mockStockModel, nil).Once()
		externalAPIMock.On("GetStockSentiment", "TEST").Return(&domain.InternalSentimentDTO{}, nil).Once()
		externalAPIMock.On("GetCompanyProfile", "TEST").Return(mockCompanyProfile, nil).Once()
		apiError := errors.New("external API error fetching news")
		externalAPIMock.On("GetLatestNews", "TEST").Return(nil, apiError).Once()

		result, err := service.GetStockFullData(stockID)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, stockID, result.Stock.ID)
		assert.Equal(t, "Test Corp Profile", result.CompanyProfile.Name)
		assert.Empty(t, result.News) // News should be empty

		repoMock.AssertExpectations(t)
		externalAPIMock.AssertExpectations(t)
	})
}


/*
	Test for the GetRecommendations method, we check:
	-	Success with less than 10 stocks (the limit), returns all
	-	Success with more than 10 stocks, returns the first 10 sorted
	-	Error with top stocks retrieving method in repository
	-	Success, External API fails, still works with 0 sentiment value
*/
func TestStocksServiceImpl_GetRecommendations(t *testing.T) {
	//helper for creating stocks
	createMockStockModel := func(id uuid.UUID, ticker string, targetFrom, targetTo float64, ratingFrom, ratingTo, action string, hoursAgo int) domain.StockModel {
		return domain.StockModel{
			ID:         id,
			Ticker:     ticker,
			TargetFrom: targetFrom,
			TargetTo:   targetTo,
			Company:    ticker + " Company",
			Action:     action,
			Brokerage:  "Test Broker",
			RatingFrom: ratingFrom,
			RatingTo:   ratingTo,
			Time:       time.Now().Add(-time.Duration(hoursAgo) * time.Hour),
		}
	}

	/////		MOCK STOCKS
	//STOCK A: Highest score
	stockA_ID := uuid.New()
	stockA := createMockStockModel(stockA_ID, "STKA", 100, 150, "Sell", "Strong-Buy", "upgraded by", 24) // High variation, high rating change, good action, recent
	sentimentA := &domain.InternalSentimentDTO{Data: []struct {
		Symbol string  `json:"symbol"`
		Year   int     `json:"year"`
		Month  int     `json:"Month"`
		Change float64 `json:"change"`
		Mspr   float64 `json:"mspr"`
	}{{Mspr: 0.8}}} 

	//STOCK B: Medium score
	stockB_ID := uuid.New()
	stockB := createMockStockModel(stockB_ID, "STKB", 100, 120, "Neutral", "Positive", "target set by", 72) // Medium variation, medium rating change, neutral action, less recent
	sentimentB := &domain.InternalSentimentDTO{Data: []struct {
		Symbol string  `json:"symbol"`
		Year   int     `json:"year"`
		Month  int     `json:"Month"`
		Change float64 `json:"change"`
		Mspr   float64 `json:"mspr"`
	}{{Mspr: 0.5}}} 

	//STOCK C: Low score
	stockC_ID := uuid.New()
	stockC := createMockStockModel(stockC_ID, "STKC", 100, 105, "Hold", "Hold", "target lowered by", 240) // Low variation, no rating change, bad action, old
	sentimentC := &domain.InternalSentimentDTO{Data: []struct {
		Symbol string  `json:"symbol"`
		Year   int     `json:"year"`
		Month  int     `json:"Month"`
		Change float64 `json:"change"`
		Mspr   float64 `json:"mspr"`
	}{{Mspr: 0.1}}}

	t.Run("success - get recommendations (fewer than 10 available, returns all sorted)", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "", externalAPIMock)
		//to force calling the method (not cache used)
		service.(*StocksServiceImpl).lastCacheSaved = time.Time{}

		//unsorted stocks
		mockTopStocks := []domain.StockModel{stockB, stockA, stockC} 

		repoMock.On("GetBasicTopStocks").Return(mockTopStocks, nil).Once()
		externalAPIMock.On("GetStockSentiment", "STKA").Return(sentimentA, nil).Once()
		externalAPIMock.On("GetStockSentiment", "STKB").Return(sentimentB, nil).Once()
		externalAPIMock.On("GetStockSentiment", "STKC").Return(sentimentC, nil).Once()

		recommendations, err := service.GetRecommendations()

		assert.NoError(t, err)
		assert.NotNil(t, recommendations)
		assert.Len(t, recommendations, 3)

		assert.Equal(t, stockA_ID, recommendations[0].ID)
		assert.Equal(t, stockB_ID, recommendations[1].ID)
		assert.Equal(t, stockC_ID, recommendations[2].ID)

		assert.True(t, recommendations[0].RecommendationScore > recommendations[1].RecommendationScore)
		assert.True(t, recommendations[1].RecommendationScore > recommendations[2].RecommendationScore)
		assert.NotZero(t, recommendations[0].AvrgSentiment)
		assert.Equal(t, "STKA", recommendations[0].Ticker)
		assert.Equal(t, (stockA.TargetTo-stockA.TargetFrom)/stockA.TargetFrom*100, recommendations[0].Percentage)

		repoMock.AssertExpectations(t)
		externalAPIMock.AssertExpectations(t)
	})

	t.Run("success - get recommendations (more than 10 available, returns top 10 sorted)", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "", externalAPIMock)
		service.(*StocksServiceImpl).lastCacheSaved = time.Time{}

		var mockTopStocks []domain.StockModel
		for i := 0; i < 12; i++ {
			id := uuid.New()
			//Create stocks with varying scores; stockA should still be highest
			var stock domain.StockModel
			if i == 0 {
				stock = stockA
			} else {
				stock = createMockStockModel(id, "STK"+string(rune('D'+i)), 100, float64(110-i), "Neutral", "Neutral", "target set by", 24*(i+1))
			}
			mockTopStocks = append(mockTopStocks, stock)
			//Mock sentiment for all. For simplicity, give stockA higher sentiment.
			if stock.Ticker == "STKA" {
				externalAPIMock.On("GetStockSentiment", stock.Ticker).Return(sentimentA, nil).Once()
			} else {
				externalAPIMock.On("GetStockSentiment", stock.Ticker).Return(&domain.InternalSentimentDTO{Data: []struct {
					Symbol string  `json:"symbol"`
					Year   int     `json:"year"`
					Month  int     `json:"Month"`
					Change float64 `json:"change"`
					Mspr   float64 `json:"mspr"`
				}{{Mspr: 0.4 - float64(i)*0.01}}}, nil).Once()
			}
		}

		repoMock.On("GetBasicTopStocks").Return(mockTopStocks, nil).Once()

		recommendations, err := service.GetRecommendations()

		assert.NoError(t, err)
		assert.NotNil(t, recommendations)
		assert.Len(t, recommendations, 10) //Should return only top 10

		//Verify Stock A is likely the first due to its high parameters
		assert.Equal(t, stockA_ID, recommendations[0].ID)
		//Verify scores are generally decreasing
		for i := 0; i < len(recommendations)-1; i++ {
			assert.True(t, recommendations[i].RecommendationScore >= recommendations[i+1].RecommendationScore)
		}

		repoMock.AssertExpectations(t)
		externalAPIMock.AssertExpectations(t)
	})

	t.Run("error - repo GetBasicTopStocks fails", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "", externalAPIMock)
		service.(*StocksServiceImpl).lastCacheSaved = time.Time{}

		repoError := errors.New("database error fetching top stocks")
		repoMock.On("GetBasicTopStocks").Return(nil, repoError).Once()

		recommendations, err := service.GetRecommendations()

		assert.Error(t, err)
		assert.Nil(t, recommendations)
		assert.Equal(t, repoError, err)

		repoMock.AssertExpectations(t)
		externalAPIMock.AssertNotCalled(t, "GetStockSentiment", mock.Anything)
	})

	t.Run("success - externalAPI GetStockSentiment fails for some (graceful handling)", func(t *testing.T) {
		repoMock := mocks.NewStocksRepo(t)
		externalAPIMock := mocks.NewExternalApiService(t)
		service := NewStocksService(repoMock, "", "", externalAPIMock)
		service.(*StocksServiceImpl).lastCacheSaved = time.Time{}

		mockTopStocks := []domain.StockModel{stockA, stockB} //tock A (sentiment success), Stock B (sentiment fail)S

		repoMock.On("GetBasicTopStocks").Return(mockTopStocks, nil).Once()
		externalAPIMock.On("GetStockSentiment", "STKA").Return(sentimentA, nil).Once()
		externalAPIMock.On("GetStockSentiment", "STKB").Return(nil, errors.New("API error for STKB sentiment")).Once() // Fails for Stock B

		recommendations, err := service.GetRecommendations()

		assert.NoError(t, err) 
		assert.NotNil(t, recommendations)
		assert.Len(t, recommendations, 2)

		//Stock A should have a higher score and its sentiment
		assert.Equal(t, stockA_ID, recommendations[0].ID)
		assert.InDelta(t, 0.8, recommendations[0].AvrgSentiment, 0.001) // sentimentA.Data[0].Mspr

		//Stock B should have AvrgSentiment = 0 because of the API error
		assert.Equal(t, stockB_ID, recommendations[1].ID)
		assert.Equal(t, float64(0), recommendations[1].AvrgSentiment)

		assert.True(t, recommendations[0].RecommendationScore > recommendations[1].RecommendationScore)

		repoMock.AssertExpectations(t)
		externalAPIMock.AssertExpectations(t)
	})
}