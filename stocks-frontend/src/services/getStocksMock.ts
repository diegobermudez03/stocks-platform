import { StockAction, StockRating, type StockModel } from '@/models/StockModel'

export const mockStocks: StockModel[] = [
  {
    ticker: "AAPL",
    targetFrom: 150,
    targetTo: 180,
    company: "Apple Inc.",
    action: StockAction.Upgrade,
    brokerage: "Goldman Sachs",
    ratingFrom: StockRating.Hold,
    ratingTo: StockRating.Buy,
    time: new Date("2025-04-29T10:00:00Z")
  },
  {
    ticker: "GOOGL",
    targetFrom: 2700,
    targetTo: 2950,
    company: "Alphabet Inc.",
    action: StockAction.Reiterated,
    brokerage: "Morgan Stanley",
    ratingFrom: StockRating.Buy,
    ratingTo: StockRating.Buy,
    time: new Date("2025-04-28T15:30:00Z")
  },
  {
    ticker: "MSFT",
    targetFrom: 310,
    targetTo: 330,
    company: "Microsoft Corp.",
    action: StockAction.Initiated,
    brokerage: "Barclays",
    ratingFrom: StockRating.Neutral,
    ratingTo: StockRating.Outperform,
    time: new Date("2025-04-27T09:00:00Z")
  },
  {
    ticker: "TSLA",
    targetFrom: 700,
    targetTo: 750,
    company: "Tesla Inc.",
    action: StockAction.TargetRaised,
    brokerage: "JP Morgan",
    ratingFrom: StockRating.Buy,
    ratingTo: StockRating.Buy,
    time: new Date("2025-04-26T14:15:00Z")
  },
  {
    ticker: "AMZN",
    targetFrom: 3300,
    targetTo: 3400,
    company: "Amazon.com Inc.",
    action: StockAction.Upgrade,
    brokerage: "Wells Fargo",
    ratingFrom: StockRating.Neutral,
    ratingTo: StockRating.Buy,
    time: new Date("2025-04-25T13:00:00Z")
  },
  {
    ticker: "NFLX",
    targetFrom: 500,
    targetTo: 480,
    company: "Netflix Inc.",
    action: StockAction.TargetLowered,
    brokerage: "Credit Suisse",
    ratingFrom: StockRating.Outperform,
    ratingTo: StockRating.Outperform,
    time: new Date("2025-04-24T16:45:00Z")
  },
  {
    ticker: "NVDA",
    targetFrom: 600,
    targetTo: 650,
    company: "NVIDIA Corp.",
    action: StockAction.Reiterated,
    brokerage: "UBS",
    ratingFrom: StockRating.Buy,
    ratingTo: StockRating.Buy,
    time: new Date("2025-04-23T11:30:00Z")
  },
  {
    ticker: "BABA",
    targetFrom: 200,
    targetTo: 220,
    company: "Alibaba Group",
    action: StockAction.Upgrade,
    brokerage: "Jefferies",
    ratingFrom: StockRating.Sell,
    ratingTo: StockRating.Neutral,
    time: new Date("2025-04-22T10:20:00Z")
  },
  {
    ticker: "INTC",
    targetFrom: 60,
    targetTo: 55,
    company: "Intel Corp.",
    action: StockAction.TargetLowered,
    brokerage: "Deutsche Bank",
    ratingFrom: StockRating.Buy,
    ratingTo: StockRating.Hold,
    time: new Date("2025-04-21T08:00:00Z")
  },
  {
    ticker: "PYPL",
    targetFrom: 270,
    targetTo: 290,
    company: "PayPal Holdings",
    action: StockAction.Initiated,
    brokerage: "HSBC",
    ratingFrom: StockRating.Neutral,
    ratingTo: StockRating.Buy,
    time: new Date("2025-04-20T17:10:00Z")
  }
]
