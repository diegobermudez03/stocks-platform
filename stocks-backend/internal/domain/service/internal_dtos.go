package service

type StockAPIRecordDTO struct {
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type StocksAPIResponseDTO struct {
	Items    []StockAPIRecordDTO `json:"items"`
	NextPage string              `json:"next_page"`
}

type InternalCompanyProfileDTO struct {
	Country              string  `json:"country"`
	Currency             string  `json:"currency"`
	EstimateCurrency     string  `json:"estimateCurrency"`
	Exchange             string  `json:"exchange"`
	FinnhubIndustry      string  `json:"finnhubIndustry"`
	IPO                  string  `json:"ipo"`
	Logo                 string  `json:"logo"`
	MarketCapitalization float64 `json:"marketCapitalization"`
	Name                 string  `json:"name"`
	Phone                string  `json:"phone"`
	ShareOutstanding     float64 `json:"shareOutstanding"`
	Ticker               string  `json:"ticker"`
	WebURL               string  `json:"weburl"`
}

type InternalNewsDTO struct {
	Category string `json:"category"`
	Datetime int64  `json:"datetime"`
	Headline string `json:"headline"`
	ID       int64  `json:"id"`
	Image    string `json:"image"`
	Related  string `json:"related"`
	Source   string `json:"source"`
	Summary  string `json:"summary"`
	URL      string `json:"url"`
}

type InternalSentimentDTO struct {
	Data []struct {
		Symbol string  `json:"symbol"`
		Year   int     `json:"year"`
		Month  int     `json:"Month"`
		Change float64 `json:"change"`
		Mspr   float64 `json:"mspr"`
	} `json:"data"`
}