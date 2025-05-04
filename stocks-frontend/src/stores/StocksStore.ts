import {defineStore} from 'pinia'
import { ref } from 'vue'
import { getStocks } from '@/services/stockServices'
import type { StockModel } from '@/models/StockModel'

export const stocksStore = defineStore('stocks', ()=>{
    const stocks = ref<StockModel[]>([])
    const loading = ref(false)
    const errorMessage = ref<string | null>(null)

    async function retrieveStocks(){
        loading.value = true
        const response = await getStocks({limit:10, page:1})
        loading.value = false
        if(response.ok){
            stocks.value = response.data
        }else{
            errorMessage.value = response.error 
        }
    } 

    return {stocks, loading, errorMessage, retrieveStocks}
})