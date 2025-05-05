<template>
    <h1>rec</h1>
    <div class=" overflow-x-auto w-full">
        <div class=" flex flex-row spacex-4 min-w-max">
            <div v-if="store.recommendations.length > 0" v-for="recommendation in store.recommendations" :key="recommendation.id"  
                class=" bg-gray-900 text-white px-10">
                <h2>{{ recommendation.action }}</h2>
            </div>
        </div>
    </div>
</template>


<script lang="ts" setup>
import { recommendationsStore } from '@/stores/RecommendationsStore';
const store = recommendationsStore()
store.retrieveRecommendations()
</script>

<style scoped>
.overflow-x-auto::-webkit-scrollbar {
  height: 8px;
}

.overflow-x-auto::-webkit-scrollbar-thumb {
  background-color: rgba(0, 0, 0, 0.3); 
  border-radius: 10px; 
}

.overflow-x-auto::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.1); 
  border-radius: 10px; 
}

.overflow-x-auto::-webkit-scrollbar-thumb:hover {
  background-color: rgba(0, 0, 0, 0.5); 
}
</style>