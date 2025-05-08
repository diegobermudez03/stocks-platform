package externalapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
)

type ExternalAPIServiceImpl struct {
	externalAPIUrl string
	externalAPIKey string
}

func NewExternalAPIService(externalAPIUrl,externalAPIKey string ) domain.ExternalApiService{
	return &ExternalAPIServiceImpl{
		externalAPIUrl: externalAPIUrl,
		externalAPIKey: externalAPIKey,
	}
}


/*
	Method to retrieve the company info of the stock
*/
func (s *ExternalAPIServiceImpl) GetCompanyProfile(symbol string)(*domain.CompanyProfileDTO, error){
	//get the company profile
	response, err := http.Get(s.externalAPIUrl + "/stock/profile2?symbol=" + symbol + "&token=" + s.externalAPIKey)
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
	companyProfileDTO :=  &domain.CompanyProfileDTO{
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
	}
	return companyProfileDTO, nil
}


/*
	Method to retrieve the latest news related with the stock symbol
*/
func (s *ExternalAPIServiceImpl) GetLatestNews(symbol string)([]domain.NewsDTO, error){
	currentDate := time.Now().Format("2006-01-02")
	response2, err := http.Get(s.externalAPIUrl + "/company-news?symbol=" + symbol + "&from=2025-01-01" + "&to=" +currentDate + "&token=" + s.externalAPIKey)
	if err != nil{
		return nil, err
	}
	defer response2.Body.Close()
	payload2, err := io.ReadAll(response2.Body)
	if err != nil{
		return nil, err
	}
	news := []InternalNewsDTO{}
	if err := json.Unmarshal(payload2, &news); err != nil{
		return nil, err
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
	return newsDTO, nil
}


/*
	Method to retrieve the latest stock sentiment of the symbol (list of latest months insider sentiment)
*/ 
func (s *ExternalAPIServiceImpl) GetStockSentiment(symbol string)(*domain.InternalSentimentDTO, error){
	currentDate := time.Now().Format("2006-01-02")
	response, err := http.Get(s.externalAPIUrl + "/stock/insider-sentiment?symbol=" + symbol + "&from=2024-01-01" + "&to=" +currentDate + "&token=" + s.externalAPIKey)
	if err != nil{
		return nil, err
	}
	defer response.Body.Close()
	payload, err := io.ReadAll(response.Body)
	if err != nil{
		return nil, err
	}
	sentiment := domain.InternalSentimentDTO{}
	if err := json.Unmarshal(payload, &sentiment); err != nil{
		return nil, err
	}
	return &sentiment, nil
}

