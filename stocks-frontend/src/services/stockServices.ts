import type { StockModel } from "@/models/StockModel"
import {mockStocks} from '@/services/getStocksMock'

type Result<T> = {ok:true, data:T} | {ok:false, error:string} 


interface GetStocksFilter {
    limit: number,
    page: number,
}

export async function getStocks(filter: GetStocksFilter): Promise<Result<StockModel[]>>{
    await new Promise(resolve => setTimeout(resolve, 500));
    return {
        ok:   true,
        data: mockStocks
    }
}