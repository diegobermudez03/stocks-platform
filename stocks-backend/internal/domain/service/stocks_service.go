package service

import (
	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/repository"
)

type StocksServiceImpl struct {
	repo repository.StocksRepo
	apiUrl string
	apiToken string
}

func NewStocksService(repo repository.StocksRepo, apiUrl, apiToken string) domain.StocksService{
	return &StocksServiceImpl{
		repo : repo,
		apiUrl: apiUrl,
		apiToken: apiToken,
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