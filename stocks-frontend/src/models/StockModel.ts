export interface StockModel{
    id: string,
    ticker:         string,
    targetFrom:     number,
    targetTo:       number,
    company:        string,
    action:         string,
    brokerage:      string,
    ratingFrom :   string,
    ratingTo :     string,
    time:           Date
}

export interface StocksWithCount{
    count: number,
    stocks: StockModel[]
}

