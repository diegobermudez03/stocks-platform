import type { FullStockModel } from "@/models/FullStockModel";
import router from "@/router";
import { getFullStock } from "@/services/stockServices";
import { defineStore } from "pinia";
import { ref } from "vue";
import { useRoute } from "vue-router";


export const fullStockStore = defineStore('fullStock',()=>{
    const loading = ref(true)
    const errorMessage = ref<string | null>(null)
    const stockInfo = ref<FullStockModel | null>(null)
    const tabSelected = ref<string>("stock")

    async function backToHome(){
        router.push('/')
    }


    async function loadStock(id:string){
        loading.value=true
        console.log("loading")
        const resp = await getFullStock(id);
        if(resp.ok){
            stockInfo.value = resp.data
        }else{
            errorMessage.value=resp.error
        }
        loading.value = false
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
        backToHome, loadStock, stockInfo, errorMessage, loading, tabSelected, changeTab, tabs
    }
})