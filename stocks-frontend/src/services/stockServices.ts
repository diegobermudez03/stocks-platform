import type { FullStockModel, NewsModel } from "@/models/FullStockModel";
import type { ParamModel } from "@/models/ParamModel";
import type { RecommendationModel } from "@/models/RecommendationModel";
import type { StockModel, StocksWithCount } from "@/models/StockModel"
import {recommendationsMock} from '@/services/mocks'

const API_BASE_URL = import.meta.env.VITE_API_URL;
const UNKNOWN_ERROR = "Internal Unknown Error"

export type Result<T> = {ok:true, data:T} | {ok:false, code:number, error:string} 

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
        if (filter.targetStart != null) params.append('target_start', filter.targetStart.toString())
        if (filter.targetEnd != null) params.append('target_end', filter.targetEnd.toString())
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
            return { 
                ok: false, 
                error: errorText ,
                code: res.status
            }
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
    } catch (err ) {
        return {
            ok: false,
            code: 0,
            error: UNKNOWN_ERROR
        }
    }
}

export async function getActions(): Promise<Result<ParamModel[]>>{
    try {
        const url = `${API_BASE_URL}/actions`
        const response = await fetch(url)
        if (!response.ok) {
            const errorText = await response.text()
            return {
                ok: false,
                error: errorText,
                code: response.status
            }
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
            code: 0,
            error: UNKNOWN_ERROR
        }
    }
}


export async function getRatings(): Promise<Result<ParamModel[]>> {
    try {
        const url = `${API_BASE_URL}/ratings`
        const response = await fetch(url)
        if (!response.ok) {
            const errorText = await response.text()
            return {
                ok: false,
                error: errorText,
                code: response.status
            }
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
            code: 0,
            error: UNKNOWN_ERROR
        }
    }
}


export async function getFullStock(id: string): Promise<Result<FullStockModel>>{
    try {
        const response = await fetch(`${API_BASE_URL}/stocks/${id}`);
        if (!response.ok) {
            const errorText = await response.text()
            return {
                ok: false,
                error: errorText,
                code: response.status
            }
        }
        const data =  await response.json()
        
        const result: FullStockModel = {
            recommendationScore: data.recommendation_score,
            avrgSentiment: data.avrg_sentiment,
            stock: {
                id: data.stock.id,
                percentage: data.stock.percentage,
                ticker: data.stock.ticker,
                targetFrom: data.stock.target_from,
                targetTo: data.stock.target_to,
                company: data.stock.company,
                action: data.stock.action,
                brokerage: data.stock.brokerage,
                ratingFrom: data.stock.rating_from,
                ratingTo: data.stock.rating_to,
                time: new Date(data.stock.time),
            },
            companyProfile: {
              country: data.company_profile.country,
              currency: data.company_profile.currency,
              exchange: data.company_profile.exchange,
              industry: data.company_profile.industry,
              ipo: new Date(data.company_profile.ipo),
              logo: data.company_profile.logo,
              marketCapital: data.company_profile.market_capital,
              name: data.company_profile.name,
              phone: data.company_profile.phone,
              webUrl: data.company_profile.web_url,
              shareOurstanding: data.company_profile.share_outstanding,
            },
            news: (data.news || []).map((n: any) => ({
              date: new Date(n.date),
              headline: n.headline,
              image: n.image,
              source: n.source,
              summary: n.summary,
            })),
        };
    
        return { ok: true, data: result };
      } catch (error:any) {
        return { 
            ok: false, 
            error: UNKNOWN_ERROR,
            code: 0
        };
      }
}


export async function getRecommendations(): Promise<Result<RecommendationModel[]>> {
    try {
      const res = await fetch(`${API_BASE_URL}/recommendations`)
      if (!res.ok) {
        const errorText = await res.text()
        return { 
            ok: false,
            code: res.status, 
            error: errorText 
        }
      }
  
      const rawData = await res.json()
  
      const recommendations: RecommendationModel[] = rawData.map((item: any) => ({
        id: item.stock.id,
        ticker: item.stock.ticker,
        targetFrom: item.stock.target_from,
        targetTo: item.stock.target_to,
        company: item.stock.company,
        action: item.stock.action,
        brokerage: item.stock.brokerage,
        ratingFrom: item.stock.rating_from,
        ratingTo: item.stock.rating_to,
        time: item.stock.time,
        percentageIncrease: item.stock.percentage,
        recommendationScore: item.recommendation_score,
        avrgSentiment: item.avrg_sentiment
      }))
  
      return { ok: true, data: recommendations }
    } catch (error) {
      return { 
        ok: false, 
        code: 0,
        error: UNKNOWN_ERROR
      }
    }
  }