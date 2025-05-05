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

    return {
        backToHome, loadStock, stockInfo, errorMessage, loading
    }
})