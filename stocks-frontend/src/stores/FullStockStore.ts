import type { ErrorModel } from "@/models/ErrorModel";
import type { FullStockModel } from "@/models/FullStockModel";
import router from "@/router";
import { getFullStock, getLivePrice } from "@/services/stockServices";
import { defineStore } from "pinia";
import { ref } from "vue";
import { useRoute } from "vue-router";


export const fullStockStore = defineStore('fullStock',()=>{
    const loading = ref(true)
    const errorMessage = ref<ErrorModel | null>(null)
    const stockInfo = ref<FullStockModel | null>(null)
    const tabSelected = ref<string>("stock")
    const currentPrice = ref<number | null>(null)

    async function backToHome(){
        router.push('/')
    }


    async function loadStock(id:string){
        loading.value=true
        errorMessage.value = null
        const resp = await getFullStock(id);
        if(resp.ok){
            stockInfo.value = resp.data
        }else{
            errorMessage.value= {
                message: resp.error,
                code: resp.code
            }
        }
        loading.value = false
    }

    function getPrice(id:string): ()=>void{
        const closeCallback = getLivePrice(id, (price)=>{
            currentPrice.value = price
        })
        return closeCallback
    }

    interface Tab{
        name: string,
        label: string
    }
    async function changeTab(tab: Tab){
        tabSelected.value = tab.name
    }

    const tabs=[{name: 'stock', label: 'Stock Info'}, {name:'company',label:'Company Profile'},{name: 'news', label:"Related News"}]

    return {
        backToHome, loadStock, stockInfo, errorMessage, loading, tabSelected, changeTab, tabs, currentPrice, getPrice
    }
})