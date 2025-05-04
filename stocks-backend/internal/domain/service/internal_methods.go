package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
	"github.com/google/uuid"
)

/*
	This is the method that is called if the db isnt populated yet, its responsible from making the HTTP calls til fill all the db
*/
func (s *StocksServiceImpl) populateWithAPI() error {
	client := &http.Client{}
	var nextPage string
	for {
		//get the url with the param if we have one (all iterations except the first one)
		url := s.apiUrl 
		if nextPage != "" {
			url = url + "?next_page=" + nextPage
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return domain.ErrMakingHttpCall
		}
		//set the auth header and make the http call
		req.Header.Set("Authorization", "Bearer "+s.apiToken)
		resp, err := client.Do(req)
		if err != nil {
			return domain.ErrMakingHttpCall
		}
		defer resp.Body.Close()
		//unmarshal the received json
		body, _ := io.ReadAll(resp.Body)
		response := StocksAPIResponseDTO{}
		if err := json.Unmarshal(body, &response); err != nil{
			return domain.ErrInvalidPayloadInHttpCall
		}
		nextPage = response.NextPage
		fmt.Println(nextPage)
		//save the bunch of stocks in the DB
		if err := s.saveBunchOfRecords(response.Items); err !=nil{
			return err
		}
		//if we reach the end then we break
		if nextPage == ""{
			break
		}
	}
	return nil
}

/*
	This internal method receives the list of DTO's and saves them in the db (verifies that the data is valid)
*/
func (s *StocksServiceImpl) saveBunchOfRecords(records []StockAPIRecordDTO) error{
	for _, dto:=range records{
		dto.TargetFrom = strings.ReplaceAll(dto.TargetFrom,"$", "")
		dto.TargetFrom = strings.ReplaceAll(dto.TargetFrom,",", "")
		dto.TargetTo = strings.ReplaceAll(dto.TargetTo, "$", "")
		dto.TargetTo = strings.ReplaceAll(dto.TargetTo, ",", "")
		targetFrom, err := strconv.ParseFloat(dto.TargetFrom, 64)
		if err != nil{
			return domain.ErrInvalidPayloadInHttpCall
		}
		targetTo, err := strconv.ParseFloat(dto.TargetTo, 64)
		if err != nil{
			return domain.ErrInvalidPayloadInHttpCall
		}
		parsedTime, err := time.Parse(time.RFC3339Nano,  dto.Time)
		if err != nil{
			return domain.ErrInvalidPayloadInHttpCall
		}
		model := domain.StockModel{
			ID: uuid.New(),
			Ticker: dto.Ticker,
			TargetFrom: targetFrom,
			TargetTo: targetTo,
			Company: dto.Company,
			Action: dto.Action,
			Brokerage: dto.Brokerage,
			RatingFrom: dto.RatingFrom,
			RatingTo: dto.RatingTo,
			Time: parsedTime,
		}
		if err := s.repo.CreateStockRecord(model); err != nil{
			return domain.ErrInternalErrorWritingToDb
		}
	}
	return nil
}

/*
	Internal method to retrieve the company profile of a given stock
*/
func (s *StocksServiceImpl) getCompanyProfile(stock *domain.StockModel)(*domain.StockDataDTO, error){
	//get the company profile
	response, err := http.Get(s.externalAPIUrl + "/stock/profile2?symbol=" + stock.Ticker + "&token=" + s.externalAPIKey)
	if err != nil{
		return nil, domain.ErrInternalError
	}
	defer response.Body.Close()
	payload, err := io.ReadAll(response.Body)
	if err != nil{
		return nil, domain.ErrInternalError
	}
	companyProfile := InternalCompanyProfileDTO{}
	if err := json.Unmarshal(payload, &companyProfile); err != nil{
		return nil, domain.ErrInternalError
	}

	//construct the preview payload, since if the news fail, we will still return succesfully
	stockData := domain.StockDataDTO{
		ID: stock.ID,
		Ticker: stock.Ticker,
		CompanyProfile: domain.CompanyProfileDTO{
			Country: companyProfile.Country,
			Currency: companyProfile.Currency,
			Exchange: companyProfile.Exchange,
			Industry: companyProfile.FinnhubIndustry,
			Ipo: companyProfile.IPO,
			Logo: companyProfile.Logo,
			MarketCapital: companyProfile.MarketCapitalization,
			Name: companyProfile.Name,
			Phone: companyProfile.Phone,
			WebUrl: companyProfile.WebURL,
			ShareOutstanding: companyProfile.ShareOutstanding,
		},
		News: []domain.NewsDTO{},
	}
	return &stockData, nil
}