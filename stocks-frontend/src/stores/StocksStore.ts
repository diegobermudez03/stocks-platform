import {defineStore} from 'pinia'
import { computed, ref } from 'vue'
import { getActions, getRatings, getStocks, type Result } from '@/services/stockServices'
import type { StockModel } from '@/models/StockModel'
import type { ParamModel } from '@/models/ParamModel'
import router from '@/router'
import type { ErrorModel } from '@/models/ErrorModel'

export const stocksStore = defineStore('stocks', ()=>{
    const greenRatings : string[] = ["Positive", "Speculative Buy", "Overweigh", "Market Outperform", "Outperform", "Sector Outperform", "Buy", "Strong-Buy"]
    const yellowRatings: string[]=["Sector Perform", "Hold", "Market Perform", "In-Line", "Sector Weight", "Equal Weight", "Peer Perform", "Neutral", "Unchanged"]
    const orangeRatings: string[]=["Cautious", "Sector Underperform", "Underperform", "Under Perform", "Underweight"]
    const redRatings: string[] =["Negative", "Sell", "Unchanged", "Reduce", "Neutral", "Overweight"]

    //for sort
    interface sortOption{
        value:string,
        label:string
    }
    const openSort = ref(false)
    const defaultEmptySort: sortOption= {
        value:'',
        label: 'Sort by'
    }
    const sortOptions = [
        {value:'CLOSEST_DATE', label:'Sort by Closest Date'},
        {value:'DISTANT_DATE', label: 'Sort by Distant Date'},
        defaultEmptySort
    ]
    const selectedSort = ref<sortOption>(defaultEmptySort)

    //for filters
    const searchQueryTmp = ref<string| null>(null)

    const searchQuery = ref<string | null>(null)
    const fromPriceTmp = ref<number | null>(null)
    const toPriceTmp = ref<number | null>(null)
    interface priceRange{
        from: number,
        to: number
    }
    const priceRange = ref<priceRange | null>(null)

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

    //debounce
    const timeout = ref<number | null>(null)

    //errors
    const errorMessage = ref<ErrorModel | null>(null)
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
        if(timeout.value){
            clearTimeout(timeout.value)
        }
        loading.value = true
        openFilter.value= null
        const response = await getStocks({
            page : page.value,
            size : size.value,
            textSearch : searchQuery.value,
            targetStart : priceRange.value?.from??null,
            targetEnd :priceRange.value?.to??null,
            ratingFrom: selectedRatingsFrom.value,
            ratingTo: selectedRatingsTo.value,
            action : selectedActions.value,
            timeStart : null,
            timeEnd : null,
            sort : selectedSort.value.value

        })
        loading.value = false
        if(response.ok){
            stocks.value = response.data.stocks
            totalRecords.value = response.data.count
        }else{
            errorMessage.value = {
                message: response.error,
                code: response.code
            }
        }
        if( page.value !== 1 && page.value > (totalRecords.value/size.value)){
            page.value=1
            retrieveStocks()
        }
    } 

    async function selectSortType(opt: sortOption){
        selectedSort.value = opt
        openSort.value = false
        retrieveStocks()
    }

    async function switchSortMenu(){
        if(openSort.value){
            openSort.value = false
        }else{
            openSort.value = true
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

    async function submitSearchFilter(){
        searchQuery.value = searchQueryTmp.value
        retrieveStocks()
    }

    async function submitPriceRange(){
        if(fromPriceTmp.value && toPriceTmp.value){
            priceRange.value = {
                from: fromPriceTmp.value,
                to: toPriceTmp.value
            }

            retrieveStocks()
        }
    }

    const activeFilters = computed(()=>{
        const filters: { label: string; onRemove: () => void }[] = [];
        if (searchQuery.value) {
            filters.push({
                label: `Search: ${searchQuery.value}`,
                onRemove: () => {searchQueryTmp.value = "" ; searchQuery.value = ''}
            });
        }

        if (priceRange.value) {
            filters.push({
            label: `From \$${priceRange.value.from.toFixed(2)} To: \$${priceRange.value.to.toFixed(2)}`,
                onRemove: () =>{
                    priceRange.value = null 
                    fromPriceTmp.value = null
                    toPriceTmp.value = null
                }
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

    async function openStock(id: string){
        router.push({name:'stockDetail', params:{
            id: id
        }})
    }

    //debounce implemented correctly
    async function debounce(){
        if (timeout.value) {
            clearTimeout(timeout.value);
        }
        timeout.value = setTimeout(()=>{
            submitSearchFilter()
        }, 500)
    }
    return {
        stocks, size, page, loading, errorMessage, retrieveStocks, 
        pages, changePage, getParams, loadingFilter, actions, ratings, 
        filterError, expandedStock, expandStock, closeStock,
        searchQuery, fromPriceTmp, toPriceTmp, selectedRatingsFrom, selectedRatingsTo, selectedActions,
        toggleFilter, openFilter, activeFilters, greenRatings, yellowRatings, orangeRatings, redRatings,openStock,
        totalRecords, openSort, sortOptions, selectedSort, selectSortType, switchSortMenu, submitSearchFilter,
        searchQueryTmp, submitPriceRange, debounce
    }
})