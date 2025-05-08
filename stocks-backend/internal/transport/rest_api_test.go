package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/transport/transportmock"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
	Tests for the get actions endpoint, should simply succesfully return either the actions
	or the error with respective internal server error code
*/
func TestRestAPIServer_getActions(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockService := transportmock.NewStocksService(t)
		server := NewRestAPIServer(ServerConfig{}, mockService)

		expectedActions := []domain.ActionDTO{
			{Action: "Buy", Count: 10},
			{Action: "Sell", Count: 5},
		}
		mockService.On("GetActions").Return(expectedActions, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/actions", nil)
		rr := httptest.NewRecorder()

		server.getActions(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var responseActions []domain.ActionDTO
		err := json.Unmarshal(rr.Body.Bytes(), &responseActions)
		assert.NoError(t, err)
		assert.Equal(t, expectedActions, responseActions)
		mockService.AssertExpectations(t)
	})

	t.Run("service error", func(t *testing.T) {
		mockService := transportmock.NewStocksService(t)
		server := NewRestAPIServer(ServerConfig{}, mockService)

		serviceErr := errors.New("failed to get actions")
		mockService.On("GetActions").Return(nil, serviceErr).Once()

		req := httptest.NewRequest(http.MethodGet, "/actions", nil)
		rr := httptest.NewRecorder()

		server.getActions(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		expectedBody := `{"error":"failed to get actions"}`
		assert.JSONEq(t, expectedBody, rr.Body.String())
		mockService.AssertExpectations(t)
	})
}


/*
	Tests for the get RATINGS endpoint, should simply succesfully return either the RATINGS
	or the error with respective internal server error code
*/
func TestRestAPIServer_getRatings(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockService := transportmock.NewStocksService(t)
		server := NewRestAPIServer(ServerConfig{}, mockService)

		expectedRatings := []domain.RatingDTO{
			{Rating: "Positive", Count: 100},
			{Rating: "Negative", Count: 50},
		}
		mockService.On("GetRatings").Return(expectedRatings, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/ratings", nil)
		rr := httptest.NewRecorder()

		server.getRatings(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var responseRatings []domain.RatingDTO
		err := json.Unmarshal(rr.Body.Bytes(), &responseRatings)
		assert.NoError(t, err)
		assert.Equal(t, expectedRatings, responseRatings)
		mockService.AssertExpectations(t)
	})

	t.Run("service error", func(t *testing.T) {
		mockService := transportmock.NewStocksService(t)
		server := NewRestAPIServer(ServerConfig{}, mockService)

		serviceErr := errors.New("failed to get ratings")
		mockService.On("GetRatings").Return(nil, serviceErr).Once()

		req := httptest.NewRequest(http.MethodGet, "/ratings", nil)
		rr := httptest.NewRecorder()

		server.getRatings(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		expectedBody := `{"error":"failed to get ratings"}`
		assert.JSONEq(t, expectedBody, rr.Body.String())
		mockService.AssertExpectations(t)
	})
}


/*
	Test for the GetStocks endpoint, should succesfully read the query, and then either
	return the stocks returned by the service, or return the error
*/
func TestRestAPIServer_getStocks(t *testing.T) {
	t.Run("success with filters", func(t *testing.T) {
		mockService := transportmock.NewStocksService(t)
		server := NewRestAPIServer(ServerConfig{}, mockService)

		expectedStocksReturn := &domain.StocksReturnDTO{
			Stocks: []domain.StockDTO{{ID: uuid.New(), Ticker: "AAPL"}},
			Count:  1,
		}

		expectedFilter := domain.GetStocksFilter{
			Page:           1,
			Size:           10,
			TextSearch:     "Apple",
			RatingFromList: []string{"Positive"},
			Sort:           "CLOSEST_DATE",
			TargetStart:    100.5,
		}

		mockService.On("GetStocks", mock.MatchedBy(func(filter domain.GetStocksFilter) bool {
			// We check a few key fields to ensure parsing logic is working
			return filter.Page == expectedFilter.Page &&
				filter.Size == expectedFilter.Size &&
				filter.TextSearch == expectedFilter.TextSearch &&
				len(filter.RatingFromList) == 1 && filter.RatingFromList[0] == "Positive" &&
				filter.Sort == expectedFilter.Sort &&
				filter.TargetStart == expectedFilter.TargetStart
		})).Return(expectedStocksReturn, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/stocks?page=1&size=10&text_search=Apple&rating_from=Positive&sort=CLOSEST_DATE&target_start=100.5", nil)
		rr := httptest.NewRecorder()

		server.getStocks(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var responseStocks domain.StocksReturnDTO
		err := json.Unmarshal(rr.Body.Bytes(), &responseStocks)
		assert.NoError(t, err)
		assert.Equal(t, expectedStocksReturn.Count, responseStocks.Count)
		assert.Equal(t, expectedStocksReturn.Stocks[0].Ticker, responseStocks.Stocks[0].Ticker)
		mockService.AssertExpectations(t)
	})

	t.Run("service error", func(t *testing.T) {
		mockService := transportmock.NewStocksService(t)
		server := NewRestAPIServer(ServerConfig{}, mockService)

		serviceErr := errors.New("failed to get stocks")
		mockService.On("GetStocks", mock.AnythingOfType("domain.GetStocksFilter")).Return(nil, serviceErr).Once()

		req := httptest.NewRequest(http.MethodGet, "/stocks", nil)
		rr := httptest.NewRecorder()

		server.getStocks(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		expectedBody := `{"error":"failed to get stocks"}`
		assert.JSONEq(t, expectedBody, rr.Body.String())
		mockService.AssertExpectations(t)
	})
}


/*
	Tests for the get stock full data endpoint, should succesfully extract the 
	stock ID from the url, and then return the data returned by the service or the error.
	There's a third case to test if the id is not passed or invalid
*/
func TestRestAPIServer_getStockData(t *testing.T) {
	stockID := uuid.New()

	//to create url param
	newRequestWithChiCtx := func(method, target string, stockIDParam string) *http.Request {
		req := httptest.NewRequest(method, target, nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", stockIDParam)
		ctx := context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		return req.WithContext(ctx)
	}

	t.Run("success", func(t *testing.T) {
		mockService := transportmock.NewStocksService(t)
		server := NewRestAPIServer(ServerConfig{}, mockService)

		expectedStockData := &domain.StockDataDTO{
			Stock: domain.StockDTO{ID: stockID, Ticker: "MSFT"},
			CompanyProfile: domain.CompanyProfileDTO{Name: "Microsoft"},
			News:           []domain.NewsDTO{{Headline: "MSFT News"}},
		}
		mockService.On("GetStockFullData", stockID).Return(expectedStockData, nil).Once()

		req := newRequestWithChiCtx(http.MethodGet, fmt.Sprintf("/stocks/%s", stockID.String()), stockID.String())
		rr := httptest.NewRecorder()

		server.getStockData(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var responseStockData domain.StockDataDTO
		err := json.Unmarshal(rr.Body.Bytes(), &responseStockData)
		assert.NoError(t, err)
		assert.Equal(t, expectedStockData.Stock.Ticker, responseStockData.Stock.Ticker)
		assert.Equal(t, expectedStockData.CompanyProfile.Name, responseStockData.CompanyProfile.Name)
		mockService.AssertExpectations(t)
	})

	t.Run("service error", func(t *testing.T) {
		mockService := transportmock.NewStocksService(t)
		server := NewRestAPIServer(ServerConfig{}, mockService)

		serviceErr := errors.New("failed to get stock data")
		mockService.On("GetStockFullData", stockID).Return(nil, serviceErr).Once()

		req := newRequestWithChiCtx(http.MethodGet, fmt.Sprintf("/stocks/%s", stockID.String()), stockID.String())
		rr := httptest.NewRecorder()

		server.getStockData(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		expectedBody := `{"error":"failed to get stock data"}`
		assert.JSONEq(t, expectedBody, rr.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("invalid stock id format", func(t *testing.T) {
		mockService := transportmock.NewStocksService(t) //Not called
		server := NewRestAPIServer(ServerConfig{}, mockService)

		req := newRequestWithChiCtx(http.MethodGet, "/stocks/not-a-uuid", "not-a-uuid")
		rr := httptest.NewRecorder()

		server.getStockData(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		expectedBody := `{"error":"invalid stock id"}`
		assert.JSONEq(t, expectedBody, rr.Body.String())
		mockService.AssertExpectations(t)
	})
}


/*
	Test for the get recommendations endpoint, should simply return the recommendations
	or error received from the domain
*/
func TestRestAPIServer_GetRecommendations(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockService := transportmock.NewStocksService(t)
		server := NewRestAPIServer(ServerConfig{}, mockService)

		stockTime := time.Now()
		expectedRecommendations := []domain.RecommendationDTO{
			{
				StockDTO: domain.StockDTO{
					ID:         uuid.New(),
					Ticker:     "REC1",
					TargetFrom: 10,
					TargetTo:   12,
					Company:    "Reco Corp 1",
					Action:     "Buy",
					Brokerage:  "Broker X",
					RatingFrom: "Neutral",
					RatingTo:   "Positive",
					Time:       stockTime,
					Percentage: 20.0,
				},
				RecommendationScore: 0.75,
				AvrgSentiment:       0.6,
			},
		}
		mockService.On("GetRecommendations").Return(expectedRecommendations, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/recommendations", nil)
		rr := httptest.NewRecorder()

		server.GetRecommendations(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var responseRecs []domain.RecommendationDTO
		err := json.Unmarshal(rr.Body.Bytes(), &responseRecs)
		assert.NoError(t, err)
		assert.Len(t, responseRecs, 1)
		assert.Equal(t, expectedRecommendations[0].StockDTO.Ticker, responseRecs[0].StockDTO.Ticker)
		assert.Equal(t, expectedRecommendations[0].RecommendationScore, responseRecs[0].RecommendationScore)

		//bECAUSE time can not be exact for comparisoon, so we check ian interval
		assert.WithinDuration(t, expectedRecommendations[0].StockDTO.Time, responseRecs[0].StockDTO.Time, time.Second)
		
		responseRecs[0].StockDTO.Time = expectedRecommendations[0].StockDTO.Time
		assert.Equal(t, expectedRecommendations, responseRecs)

		mockService.AssertExpectations(t)
	})

	t.Run("service error", func(t *testing.T) {
		mockService := transportmock.NewStocksService(t)
		server := NewRestAPIServer(ServerConfig{}, mockService)

		serviceErr := errors.New("failed to get recommendations")
		mockService.On("GetRecommendations").Return(nil, serviceErr).Once()

		req := httptest.NewRequest(http.MethodGet, "/recommendations", nil)
		rr := httptest.NewRecorder()

		server.GetRecommendations(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		expectedBody := `{"error":"failed to get recommendations"}`
		assert.JSONEq(t, expectedBody, rr.Body.String())
		mockService.AssertExpectations(t)
	})
}