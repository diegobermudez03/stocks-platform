import type { StockModel } from "./StockModel";

export interface FullStockModel{
    stock : StockModel,
    companyProfile: {
        country : string,
        currency: string,
        exchange : string,
        industry : string,
        ipo: Date,
        logo: string,
        marketCapital : number,
        name: string,
        phone: string,
        webUrl: string,
        shareOurstanding: number
    },
    news: NewsModel[]
}

export interface NewsModel{
    date: Date,
    headline: string,
    image: string,
    source: string,
    summary: string,
}