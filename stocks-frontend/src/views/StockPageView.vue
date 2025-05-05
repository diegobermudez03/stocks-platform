<template>
    <div class=" bg-gray-900">
        <button class=" text-white text-2xl px-10 pt-5" @click="store.backToHome"> Back to homepage</button>
    </div>
    <RecommendationsComponent/>
    <div v-if="store.loading">
        LOADING
    </div>
    <div v-else-if="store.errorMessage">
        ERROR 
    </div>
    <div v-else class=" flex flex-col p-10">
        <div class=" flex flex-col md:flex-row gap-5">
            <img class="w-full md:w-1/3 rounded-lg" :src="store.stockInfo?.companyProfile.logo"/>
            <div class=" flex flex-col">
                <h1 class=" text-6xl font-bold">{{ store.stockInfo?.stock.ticker }}</h1>
                <h1 class=" text-6xl font-bold">{{ store.stockInfo?.stock.company }}</h1>
            </div>
        </div>
        <div class="mt-10">
            <h2 class=" font-bold text-2xl">Latest news for ticker</h2>
            <div class=" border-gray-400 border-2 rounded-lg" 
                 style="max-height: calc(100vh - 2rem);">
                <div v-for="news in store.stockInfo?.news || []"
                class=" flex flex-col p-5 items-start">
                    <hr class=" border-2 border-gray-400 w-full my-2">
                    <img class=" w-36" :src="news.image"/>
                    <div class=" flex flex-col">
                        <h3 class="text-lg font-bold">{{ news.headline }}</h3>
                        <p>{{ news.summary }}</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import RecommendationsComponent from '@/components/RecommendationsComponent.vue';
import {fullStockStore} from '@/stores/FullStockStore';
import { onMounted, watch } from 'vue';
import { useRoute } from 'vue-router';

const store = fullStockStore();
const route = useRoute()

onMounted(() => {
    const stockId = Array.isArray(route.params.id) ? route.params.id[0] : route.params.id;
    store.loadStock(stockId)
})

watch(() => route.params.id, (newId) => {
    const stockId = Array.isArray(route.params.id) ? route.params.id[0] : route.params.id;
    store.loadStock(stockId)
})
</script>