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
	TargetFrom float64	`json:"target_to"`
	TargetTo float64	`json:"target_from"`
	Company string	`json:"company"`
	Action string	`json:"action"`
	Brokerage string	`json:"brokerage"`
	RatingFrom string	`json:"rating_from"`
	RatingTo string		`json:"rating_to"`
	Time time.Time		`json:"time"`
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
	ID	 uuid.UUID	`json:"id"`
	Ticker string	`json:"ticker"`
	CompanyProfile CompanyProfileDTO `json:"company_profile"`
	News []NewsDTO `json:"news"`
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
