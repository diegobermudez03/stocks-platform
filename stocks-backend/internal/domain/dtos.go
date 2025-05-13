package domain

import (
	"time"

	"github.com/google/uuid"
)

var SortMap = map[string]string{
	"CLOSEST_DATE" : "time DESC",
	"DISTANT_DATE" : "time ASC",
}

type GetStocksFilter struct {
	Page           int
	Size           int
	TextSearch     string
	RatingFromList []string
	RatingToList   []string
	TimeStart      string
	TimeEnd        string
	Sort           string
	ActionList     []string
	TargetStart    float64
	TargetEnd      float64
}

type StockDTO struct {
	ID uuid.UUID	`json:"id"`
	Ticker string	`json:"ticker"`
	TargetFrom float64	`json:"target_from"`
	TargetTo float64	`json:"target_to"`
	Company string	`json:"company"`
	Action string	`json:"action"`
	Brokerage string	`json:"brokerage"`
	RatingFrom string	`json:"rating_from"`
	RatingTo string		`json:"rating_to"`
	Time time.Time		`json:"time"`
	Percentage float64 	`json:"percentage"`
}

type StocksReturnDTO struct{
	Count 	int64			`json:"count"`
	Stocks 	[]StockDTO		`json:"stocks"`
}

type RatingDTO struct {
	Rating string `json:"rating"`
	Count  int64  `json:"count"`
}

type ActionDTO struct {
	Action string `json:"action"`
	Count  int64  `json:"count"`
}


type StockDataDTO struct{
	Stock StockDTO `json:"stock"`
	CompanyProfile CompanyProfileDTO `json:"company_profile"`
	News []NewsDTO `json:"news"`
	RecommendationScore float64 `json:"recommendation_score"`
	AvrgSentiment	float64 	`json:"avrg_sentiment"`
}


type CompanyProfileDTO struct{
	Country string `json:"country"`
	Currency string `json:"currency"`
	Exchange string `json:"exchange"`
	Industry string `json:"industry"`
	Ipo string `json:"ipo"`
	Logo string `json:"logo"`
	MarketCapital float64 `json:"market_capital"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	WebUrl string `json:"web_url"`
	ShareOutstanding float64 `json:"share_outstanding"`
}


type NewsDTO struct{
	Date time.Time `json:"date"`
	Headline string `json:"headline"`
	Image string `json:"image"`
	Source string `json:"source"`
	Summary string `json:"summary"`
}


type RecommendationDTO struct{
	StockDTO `json:"stock"`
	RecommendationScore float64 `json:"recommendation_score"`
	AvrgSentiment	float64 	`json:"avrg_sentiment"`
}

type PriceUpdateDTO struct{
	Price 	float64	`json:"price"`
}

/*
	Only used for internal processment between services, declared at the domain
	top level in order to not couple the implementations of the services (both dependent only on the domain)
*/
type InternalSentimentDTO struct {
	Data []struct {
		Symbol string  `json:"symbol"`
		Year   int     `json:"year"`
		Month  int     `json:"Month"`
		Change float64 `json:"change"`
		Mspr   float64 `json:"mspr"`
	} `json:"data"`
}


