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
                <div v-if="store.tabSelected==='stock'" class="mt-5  space-y-3">
                    <h2 class="text-2xl font-bold">Stock Details</h2>
                    <p><span class="font-semibold">Ticker:</span> {{ store.stockInfo!.stock.ticker }}</p>
                    <p><span class="font-semibold">Company:</span> {{ store.stockInfo!.stock.company }}</p>
                    <p><span class="font-semibold">Brokerage:</span> {{ store.stockInfo!.stock.brokerage }}</p>
                    <p><span class="font-semibold">Action:</span> {{ store.stockInfo!.stock.action }}</p>
                    <p><span class="font-semibold">Rating:</span> {{ store.stockInfo!.stock.ratingFrom }} → {{ store.stockInfo?.stock.ratingTo }}</p>
                    <p><span class="font-semibold">Target Price:</span> ${{ store.stockInfo!.stock.targetFrom.toFixed(2) }} → ${{ store.stockInfo?.stock.targetTo.toFixed(2) }}</p>
                    <p><span class="font-semibold">Date:</span> {{ new Date(store.stockInfo!.stock.time).toLocaleDateString() }}</p>
                    <p><span class="font-semibold">Percentage change:</span> 
                        <span :class="{
                        'text-green-500': store.stockInfo!.stock.percentage > 0,
                        'text-red-500': store.stockInfo!.stock.percentage <= 0
                        }">
                        {{ (store.stockInfo!.stock.percentage > 0 ? '+' : '') + store.stockInfo!.stock.percentage.toFixed(2) + '%' }}
                        </span>
                    </p>
                </div>
                <div v-if="store.tabSelected === 'company' && store.stockInfo?.companyProfile" class="mt-5  space-y-3">
                    <h2 class="text-2xl font-bold mb-4">Company Profile</h2>

                    <p><span class="font-semibold">Name:</span> {{ store.stockInfo!.companyProfile.name }}</p>
                    <p><span class="font-semibold">Country:</span> {{ store.stockInfo!.companyProfile.country }}</p>
                    <p><span class="font-semibold">Currency:</span> {{ store.stockInfo!.companyProfile.currency }}</p>
                    <p><span class="font-semibold">Exchange:</span> {{ store.stockInfo!.companyProfile.exchange }}</p>
                    <p><span class="font-semibold">Industry:</span> {{ store.stockInfo!.companyProfile.industry }}</p>
                    <p><span class="font-semibold">IPO Date:</span> {{ new Date(store.stockInfo!.companyProfile.ipo).toLocaleDateString() }}</p>
                    <p><span class="font-semibold">Market Capitalization:</span> ${{ (store.stockInfo!.companyProfile.marketCapital / 1_000_000_000).toFixed(2) }}B</p>
                    <p><span class="font-semibold">Shares Outstanding:</span> {{ store.stockInfo?.companyProfile.shareOurstanding.toLocaleString() }}</p>
                    <p><span class="font-semibold">Phone:</span> {{ store.stockInfo?.companyProfile.phone }}</p>
                    <p>
                        <span class="font-semibold">Website:</span>
                        <a :href="store.stockInfo?.companyProfile.webUrl" class="text-blue-400 underline" target="_blank">
                        {{ store.stockInfo?.companyProfile.webUrl }}
                        </a>
                    </p>
                </div>
                <div v-if="store.tabSelected==='news'">
                    <div class=" border-gray-400 border-2 rounded-lg" >
                        <div v-for="news in store.stockInfo?.news || []"
                        class=" flex flex-col p-5 items-start">
                            <hr class=" border-2 border-gray-400 w-full my-2">
                            <img class=" w-36" :src="news.image"/>
                            <p>{{ news.source }}</p>
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