import {defineStore} from 'pinia'
import { computed, ref } from 'vue'
import { getActions, getRatings, getStocks, type Result } from '@/services/stockServices'
import type { StockModel } from '@/models/StockModel'
import type { ParamModel } from '@/models/ParamModel'

export const stocksStore = defineStore('stocks', ()=>{

    //for filters
    const searchQuery = ref<string| null>(null)
    const fromPrice = ref<number | null>(null)
    const toPrice = ref<number | null>(null)
    const selectedRatingsFrom = ref<string[]>([])
    const selectedRatingsTo = ref<string[]>([])
    const selectedActions = ref<string[]>([]) 
    const openFilter = ref<string | null>(null)
    //
    const stocks = ref<StockModel[]>([])
    const expandedStock = ref<string>("")
    const totalRecords = ref(0)
    const loading = ref(true)
    const page = ref(1)
    const size = ref(21)
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
        openFilter.value= null
        const response = await getStocks({
            page : page.value,
            size : size.value,
            textSearch : searchQuery.value,
            targetStart : fromPrice.value,
            targetEnd : fromPrice.value,
            ratingFrom: selectedRatingsFrom.value,
            ratingTo: selectedRatingsTo.value,
            action : selectedActions.value,
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
        if( page.value !== 1 && page.value > (totalRecords.value/size.value)){
            page.value=1
            retrieveStocks()
        }
    } 

    async function changePage(pageNumber:number){
        if(pageNumber <= Math.ceil(totalRecords.value/size.value) && pageNumber>=1){
            expandedStock.value=""
            page.value=pageNumber
            retrieveStocks()
        }
    }

    async function toggleFilter(filter: string){
        if(openFilter.value === filter){
            openFilter.value = null
            return
        }
        openFilter.value = filter
    }
    async function getParams(){
        const actionsRes = await getActions()
        const ratingsRes = await getRatings()
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
        loadingFilter.value = false
    }

    async function expandStock(id :string ){
        expandedStock.value=id
    }

    async function closeStock(id :string ){
        expandedStock.value = ""
    }


    const activeFilters = computed(()=>{
        const filters: { label: string; onRemove: () => void }[] = [];
        if (searchQuery.value) {
            filters.push({
                label: `Search: ${searchQuery.value}`,
                onRemove: () => (searchQuery.value = '')
            });
        }

        if (fromPrice.value) {
            filters.push({
            label: `From price: ${fromPrice.value}`,
                onRemove: () => (fromPrice.value = null)
            });
        }

        if (toPrice.value) {
            filters.push({
            label: `To price: ${toPrice.value}`,
            onRemove: () => (toPrice.value = null)
            });
        }

        selectedActions.value.forEach(action => {
            filters.push({
            label: `Action: ${action}`,
            onRemove: () => {
                selectedActions.value = selectedActions.value.filter(a => a !== action);
            }
            });
        });

        selectedRatingsFrom.value.forEach(rating => {
            filters.push({
            label: `Rating From: ${rating}`,
            onRemove: () => {
                selectedRatingsFrom.value = selectedRatingsFrom.value.filter(r => r !== rating);
            }
            });
        });

        selectedRatingsTo.value.forEach(rating => {
            filters.push({
            label: `Rating To: ${rating}`,
            onRemove: () => {
                selectedRatingsTo.value = selectedRatingsTo.value.filter(r => r !== rating);
            }
            });
        });

        return filters;
    })

    return {
        stocks, size, page, loading, errorMessage, retrieveStocks, 
        pages, changePage, getParams, loadingFilter, actions, ratings, 
        filterError, expandedStock, expandStock, closeStock,
        searchQuery, fromPrice, toPrice, selectedRatingsFrom, selectedRatingsTo, selectedActions,
        toggleFilter, openFilter, activeFilters
    }
})