package domain

import (
	"time"

	"github.com/google/uuid"
)

type StockModel struct {
	ID     		uuid.UUID	`gorm:"type:uuid;primary_key"`
	Ticker 		string
	TargetFrom 	float64
	TargetTo 	float64
	Company 	string 
	Action 		string 
	Brokerage 	string 
	RatingFrom 	string 
	RatingTo 	string 
	Time 		time.Time
}

type ParamValueModel struct{
	Value 	string 
	Count 	int64
}


type GetStocksFilterModel struct{
	Page           int
	Size           int
	TextSearch     string
	RatingFromList []string
	RatingToList   []string
	TimeStart      *time.Time
	TimeEnd        *time.Time
	Sort           string
	ActionList     []string
	TargetStart    *float64
	TargetEnd      *float64
}