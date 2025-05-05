import {defineStore} from 'pinia'
import { computed, ref } from 'vue'
import { getActions, getRatings, getStocks, type Result } from '@/services/stockServices'
import type { StockModel } from '@/models/StockModel'
import type { ParamModel } from '@/models/ParamModel'

export const stocksStore = defineStore('stocks', ()=>{
    const stocks = ref<StockModel[]>([])
    const expandedStock = ref<string>("")
    const totalRecords = ref(0)
    const loading = ref(true)
    const page = ref(1)
    const size = ref(12)
    const actions = ref<ParamModel[]>([])
    const ratings = ref<ParamModel[]>([])
    const errorMessage = ref<string | null>(null)
    const filterError = ref<string | null>(null)
    const loadingFilter = ref(true)
    const range = computed(()=>{
        if(totalRecords.value/size.value > 9) return 9
        else return totalRecords.value/size.value
    })
    const pages = computed (()=>{
        const firstPage: number = page.value > (range.value!/2) ? page.value - (range.value!/2): 1;
        var lastPage: number = page.value > (range.value!/2) ? page.value + (range.value!/2): range.value!;
        if(lastPage > (totalRecords.value/size.value)){
            lastPage = (totalRecords.value/size.value)
        }
        return Array.from({ length: Math.floor(lastPage) - Math.ceil(firstPage) + 1 }, (_, i) => i + Math.ceil(firstPage))
    })
    

    async function retrieveStocks(){
        loading.value = true
        const response = await getStocks({
            page : null,
            size : null,
            textSearch : null,
            targetStart : null,
            targetEnd : null,
            ratingFrom: [],
            ratingTo: [],
            action : [],
            timeStart : null,
            timeEnd : null,
            sort : null

        })
        loading.value = false
        if(response.ok){
            stocks.value = response.data.stocks
            totalRecords.value = response.data.count
        }else{
            errorMessage.value = response.error 
        }
    } 

    async function changePage(pageNumber:number){
        if(pageNumber < (totalRecords.value/size.value) && pageNumber>=1){
            expandedStock.value=""
            page.value=pageNumber
        }
    }

    async function getParams(){
        console.log("getting params")
        const actionsRes = await getActions()
        const ratingsRes = await getRatings()
        console.log("after all")
        if(actionsRes.ok){
            actions.value = actionsRes.data
        }else{
            filterError.value = actionsRes.error
        }
        if(ratingsRes.ok){
            ratings.value = ratingsRes.data
        }else{
            filterError.value = ratingsRes.error
        }
        console.log("got params")
        loadingFilter.value = false
    }

    async function expandStock(id :string ){
        expandedStock.value=id
    }

    async function closeStock(id :string ){
        expandedStock.value = ""
    }

    return {
        stocks, size, page, loading, errorMessage, retrieveStocks, 
        pages, changePage, getParams, loadingFilter, actions, ratings, 
        filterError, expandedStock, expandStock, closeStock
    }
})