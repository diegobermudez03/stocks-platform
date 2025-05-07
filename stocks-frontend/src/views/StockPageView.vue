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
    <div v-else class=" flex flex-col p-10 justify-center">
        <div class=" flex flex-col md:flex-row w-full max-w-[1400px] mx-auto px-4 items-start gap-4">
            <img class="w-full md:w-2/6 h-auto  object-contain rounded-lg" :src="store.stockInfo?.companyProfile.logo"/>
            <div class=" flex flex-col">
                <h1 class=" text-6xl font-bold">{{ store.stockInfo?.stock.ticker }}</h1>
                <h1 class=" text-6xl font-bold">{{ store.stockInfo?.stock.company }}</h1>
                <!--TAB NAVIGATON MENU-->
                <div class=" flex flex-row border-b-2 border-gray-400 w-full mt-5">
                    <button v-for="tab in store.tabs" 
                    :key="tab.name" 
                    class=" hover:bg-gray-400 border-gray-700 rounded-sm px-3 py-1" 
                    :class="{
                        'border-black': store.tabSelected===tab.name, 
                        'bg-slate-400': store.tabSelected===tab.name, 
                        'border-b-4' : store.tabSelected===tab.name
                    }"
                    @click="store.changeTab(tab)"
                    >{{ tab.label }}</button>
                </div>
                <!--CONTAINER OF TAB CHOOSEN-->
                <div v-if="store.tabSelected==='news'">
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