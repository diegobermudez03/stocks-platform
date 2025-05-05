import type { RecommendationModel } from "@/models/RecommendationModel";
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

    return {
        recommendations, 
        loading, 
        errorMessage, 
        retrieveRecommendations
    }
})