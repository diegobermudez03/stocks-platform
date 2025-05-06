import type { FullStockModel, NewsModel } from "@/models/FullStockModel";
import type { ParamModel } from "@/models/ParamModel";
import type { RecommendationModel } from "@/models/RecommendationModel";
import type { StockModel, StocksWithCount } from "@/models/StockModel"
import {recommendationsMock} from '@/services/mocks'

const API_BASE_URL = import.meta.env.VITE_API_URL;

export type Result<T> = {ok:true, data:T} | {ok:false, error:string} 

interface GetStocksFilter {
    page: number | null,
    size: number | null,
    textSearch : string | null,
    targetStart : number | null,
    targetEnd : number | null,
    ratingFrom : string[],
    ratingTo : string[],
    action : string[],
    timeStart : string | null,
    timeEnd : string | null,
    sort : string | null,
}

export async function getStocks(filter: GetStocksFilter): Promise<Result<StocksWithCount>> {
    try {
        const params = new URLSearchParams()
        if (filter.page != null) params.append('page', filter.page.toString())
        if (filter.size != null) params.append('size', filter.size.toString())
        if (filter.textSearch) params.append('text_search', filter.textSearch)
        if (filter.targetStart != null) params.append('target_from', filter.targetStart.toString())
        if (filter.targetEnd != null) params.append('target_to', filter.targetEnd.toString())
        if (filter.ratingFrom && filter.ratingFrom.length) filter.ratingFrom.forEach(val => params.append('rating_from', val))
        if (filter.ratingTo && filter.ratingTo.length) filter.ratingTo.forEach(val => params.append('rating_to', val))
        if (filter.action && filter.action.length) filter.action.forEach(val => params.append('action', val))
        if (filter.timeStart) params.append('time_start', filter.timeStart)
        if (filter.timeEnd) params.append('time_end', filter.timeEnd)
        if (filter.sort) params.append('sort', filter.sort)
        const url = `${API_BASE_URL}/stocks?${params.toString()}`
        const res = await fetch(url)
        if (!res.ok) {
            const errorText = await res.text()
            return { ok: false, error: errorText }
        }
        const json = await res.json()
        const mappedStocks: StockModel[] = json.stocks.map((s: any) => ({
            id: s.id,
            percentage: s.percentage,
            ticker: s.ticker,
            targetFrom: s.target_from,
            targetTo: s.target_to,
            company: s.company,
            action: s.action,
            brokerage: s.brokerage,
            ratingFrom: s.rating_from,
            ratingTo: s.rating_to,
            time: new Date(s.time),
        }))
        return {
            ok: true,
            data: {
                count: json.count,
                stocks: mappedStocks
            }
        }
    } catch (err) {
        return {
            ok: false,
            error: (err as Error).message
        }
    }
}

export async function getActions(): Promise<Result<ParamModel[]>>{
    try {
        const url = `${API_BASE_URL}/actions`
        const response = await fetch(url)
        if (!response.ok) {
            throw new Error(`Failed to fetch actions: ${response.statusText}`)
        }

        const json: { action: string; count: number }[] = await response.json()
        const data: ParamModel[] = json.map(actionDTO => ({
            name: actionDTO.action,
            count: actionDTO.count
        }))
        return {
            ok: true,
            data
        }
    } catch (error: any) {
        return {
            ok: false,
            error: error.message
        }
    }
}


export async function getRatings(): Promise<Result<ParamModel[]>> {
    try {
        const url = `${API_BASE_URL}/ratings`
        const response = await fetch(url)
        if (!response.ok) {
            throw new Error(`Failed to fetch ratings: ${response.statusText}`)
        }

        const json: { rating: string; count: number }[] = await response.json()
        const data: ParamModel[] = json.map(ratingDto => ({
            name: ratingDto.rating,
            count: ratingDto.count
        }))
        return {
            ok: true,
            data
        }
    } catch (error: any) {
        return {
            ok: false,
            error: error.message
        }
    }
}


export async function getFullStock(id: string): Promise<Result<FullStockModel>>{
    await new Promise(resolve => setTimeout(resolve, 1000));
    const json = {
        "stock": {
          "id": "d613dc33-65ed-4b72-923a-b7e4bb46cf89",
          "ticker": "MODG",
          "target_from": 12,
          "target_to": 7,
          "company": "Topgolf Callaway Brands",
          "action": "target lowered by",
          "brokerage": "Truist Financial",
          "rating_from": "Buy",
          "rating_to": "Buy",
          "time": "2025-04-14T19:30:13.351058-05:00",
          "percentage": -41.66666666666667
        },
        "company_profile": {
          "country": "US",
          "currency": "USD",
          "exchange": "NEW YORK STOCK EXCHANGE, INC.",
          "industry": "Leisure Products",
          "ipo": "1992-02-28",
          "logo": "https://static2.finnhub.io/file/publicdatany/finnhubimage/stock_logo/ELY.png",
          "market_capital": 1247.657923440876,
          "name": "Topgolf Callaway Brands Corp",
          "phone": "17609311771",
          "web_url": "https://www.topgolfcallawaybrands.com/",
          "share_outstanding": 183.75
        },
        "news": [
          {
            "date": "2025-05-05T11:19:06-05:00",
            "headline": "Topgolf Callaway Brands to Release First Quarter 2025 Financial Results",
            "image": "",
            "source": "Finnhub",
            "summary": "CARLSBAD, Calif., May 5, 2025 /PRNewswire/ -- Topgolf Callaway Brands Corp. announced today that it intends to release its first quarter 2025 financial results on Monday, May 12, 2025, after the..."
          },
          {
            "date": "2025-05-05T08:58:00-05:00",
            "headline": "Marriott to Post Q1 Earnings: What's in the Offing for the Stock?",
            "image": "https://s.yimg.com/rz/stage/p/yahoo_finance_en-US_h_p_finance_2.png",
            "source": "Yahoo",
            "summary": "MAR'S first-quarter 2025 results are likely to benefit from robust RevPAR and ADR growth."
          },
          {
            "date": "2025-04-30T23:35:12-05:00",
            "headline": "3 Unprofitable Stocks Walking a Fine Line",
            "image": "https://s.yimg.com/rz/stage/p/yahoo_finance_en-US_h_p_finance_2.png",
            "source": "Yahoo",
            "summary": "Unprofitable companies can burn through cash quickly, leaving investors exposed if they fail to turn things around."
          }
        ]
      }
      
      const fullStockModel: FullStockModel = {
        stock: {
          id: json.stock.id,
          ticker: json.stock.ticker,
          targetFrom: json.stock.target_from,
          targetTo: json.stock.target_to,
          company: json.stock.company,
          action: json.stock.action,
          brokerage: json.stock.brokerage,
          ratingFrom: json.stock.rating_from,
          ratingTo: json.stock.rating_to,
          time: new Date(json.stock.time),
          percentage: json.stock.percentage
        },
        companyProfile: {
          country: json.company_profile.country,
          currency: json.company_profile.currency,
          exchange: json.company_profile.exchange,
          industry: json.company_profile.industry,
          ipo: new Date(json.company_profile.ipo),
          logo: json.company_profile.logo,
          marketCapital: json.company_profile.market_capital,
          name: json.company_profile.name,
          phone: json.company_profile.phone,
          webUrl: json.company_profile.web_url,
          shareOurstanding: json.company_profile.share_outstanding
        },
        news: json.news.map((n): NewsModel => ({
          date: new Date(n.date),
          headline: n.headline,
          image: n.image,
          source: n.source,
          summary: n.summary
        }))
      }
    return {
        ok: true,
        data: fullStockModel
    }
}


export async function getRecommendations(): Promise<Result<RecommendationModel[]>>{
    await new Promise(resolve => setTimeout(resolve, 500));
    return {
        ok: true,
        data: recommendationsMock
    }
}

