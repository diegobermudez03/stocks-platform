package domain

import (
	"time"

	"github.com/google/uuid"
)

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