import type { RecommendationModel } from "@/models/RecommendationModel";
import router from "@/router";
import { getRecommendations } from "@/services/stockServices";
import { defineStore } from "pinia";
import { ref } from "vue";


export const recommendationsStore = defineStore('recommendations', () => {
    const recommendations = ref<RecommendationModel[]>([])
    const loading = ref(true)
    const errorMessage = ref<string | null>(null)
    
    async function retrieveRecommendations(){
        const resp = await getRecommendations()
        if(resp.ok){
            recommendations.value = resp.data
        }else{
            errorMessage.value = resp.error
        }
        loading.value = false
    }

    async function openStock(id: string){
        router.push({name:'stockDetail', params:{
            id: id
        }})
    }

    return {
        recommendations, 
        loading, 
        errorMessage, 
        retrieveRecommendations,
        openStock
    }
})