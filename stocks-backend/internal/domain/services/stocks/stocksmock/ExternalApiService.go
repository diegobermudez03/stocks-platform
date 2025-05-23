// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	domain "github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// ExternalApiService is an autogenerated mock type for the ExternalApiService type
type ExternalApiService struct {
	mock.Mock
}

// CloseLiveConnection provides a mock function with no fields
func (_m *ExternalApiService) CloseLiveConnection() {
	_m.Called()
}

// GetCompanyProfile provides a mock function with given fields: symbol
func (_m *ExternalApiService) GetCompanyProfile(symbol string) (*domain.CompanyProfileDTO, error) {
	ret := _m.Called(symbol)

	if len(ret) == 0 {
		panic("no return value specified for GetCompanyProfile")
	}

	var r0 *domain.CompanyProfileDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.CompanyProfileDTO, error)); ok {
		return rf(symbol)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.CompanyProfileDTO); ok {
		r0 = rf(symbol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CompanyProfileDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(symbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestNews provides a mock function with given fields: symbol
func (_m *ExternalApiService) GetLatestNews(symbol string) ([]domain.NewsDTO, error) {
	ret := _m.Called(symbol)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestNews")
	}

	var r0 []domain.NewsDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]domain.NewsDTO, error)); ok {
		return rf(symbol)
	}
	if rf, ok := ret.Get(0).(func(string) []domain.NewsDTO); ok {
		r0 = rf(symbol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.NewsDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(symbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStockSentiment provides a mock function with given fields: symbol
func (_m *ExternalApiService) GetStockSentiment(symbol string) (*domain.InternalSentimentDTO, error) {
	ret := _m.Called(symbol)

	if len(ret) == 0 {
		panic("no return value specified for GetStockSentiment")
	}

	var r0 *domain.InternalSentimentDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.InternalSentimentDTO, error)); ok {
		return rf(symbol)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.InternalSentimentDTO); ok {
		r0 = rf(symbol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.InternalSentimentDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(symbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartLiveConnection provides a mock function with no fields
func (_m *ExternalApiService) StartLiveConnection() (chan string, chan domain.StockPriceUpdate, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StartLiveConnection")
	}

	var r0 chan string
	var r1 chan domain.StockPriceUpdate
	var r2 error
	if rf, ok := ret.Get(0).(func() (chan string, chan domain.StockPriceUpdate, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() chan string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan string)
		}
	}

	if rf, ok := ret.Get(1).(func() chan domain.StockPriceUpdate); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan domain.StockPriceUpdate)
		}
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewExternalApiService creates a new instance of ExternalApiService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExternalApiService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExternalApiService {
	mock := &ExternalApiService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
