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

export async function getStocks(filter: GetStocksFilter): Promise<Result<StocksWithCount>>{
    await new Promise(resolve => setTimeout(resolve, 500));
    return {
        ok:   true,
        data: {
            count : 300,
            stocks: mockStocks
        }
    }
}

export async function getActions(): Promise<Result<ParamModel[]>>{
    await new Promise(resolve => setTimeout(resolve, 500));
    return {
        ok: true,
        data:[
            {count: 5 , name: "Target lowered"},
            {count: 12 , name: "Target higher"},
            {count: 63 , name: "Equal"},
            {count: 51, name: "target reached"},
            {count: 10, name: "lost target"}
        ]
    }
}


export async function getRatings(): Promise<Result<ParamModel[]>>{
    await new Promise(resolve => setTimeout(resolve, 500));
    return {
        ok: true,
        data:[
            {count: 4 , name: "Risk sell"},
            {count: 102 , name: "Target reached"},
            {count: 623 , name: "Equal"},
            {count: 536, name: "target reached"},
            {count: 110, name: "Buy"},
            {count: 110, name: "Sell"},
            {count: 110, name: "BLANACE"},
            {count: 110, name: "Risk buy"}
        ]
    }
}


export async function getRecommendations(): Promise<Result<RecommendationModel[]>>{
    await new Promise(resolve => setTimeout(resolve, 500));
    return {
        ok: true,
        data: recommendationsMock
    }
}