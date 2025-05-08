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