export interface StockModel{
    ticker:         string,
    targetFrom:     number,
    targetTo:       number,
    company:        string,
    action:         StockAction,
    brokerage:      string,
    ratingFrom :   StockRating,
    ratingTo :     StockRating,
    time:           Date
}

export enum StockAction{
    Upgrade,
    Downgrade,
    Reiterated,
    Initiated,
    TargetRaised,
    TargetLowered
}

export enum StockRating{
    Buy,
    Sell,
    Neutral,
    Outperform,
    Overweight,
    MarketPerform,
    EqualWeight,
    Hold
}