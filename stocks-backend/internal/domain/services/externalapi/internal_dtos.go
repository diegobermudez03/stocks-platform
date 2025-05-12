package externalapi

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

type InternalSuscribeDTO struct {
	Type   string `json:"type"`
	Symbol string `json:"symbol"`
}

type InternalStockPriceUpdate struct {
	Data []struct {
		P float64 `json:"p"`
		S string  `json:"s"`
		T int64   `json:"t"`
		V float64 `json:"v"`
	} `json:"data"`
	Type string `json:"type"`
}