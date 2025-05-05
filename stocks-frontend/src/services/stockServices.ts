import { API_BASE_URL } from "@/config";
import type { ParamModel } from "@/models/ParamModel";
import type { RecommendationModel } from "@/models/RecommendationModel";
import type { StockModel, StocksWithCount } from "@/models/StockModel"
import {mockStocks, recommendationsMock} from '@/services/mocks'

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


export async function getRecommendations(): Promise<Result<RecommendationModel[]>>{
    await new Promise(resolve => setTimeout(resolve, 500));
    return {
        ok: true,
        data: recommendationsMock
    }
}