<template>
    <div class=" bg-gray-900">
        <button
            @click="store.backToHome"
            class="flex items-center gap-2 text-white text-lg font-semibold bg-gray-800 px-4 py-2 rounded-lg shadow hover:bg-gray-700 transition" >
            <ArrowLeft class="w-4 h-4" />
            Back to Search
        </button>
    </div>
    <RecommendationsComponent/>
    <!--FOR WHEN LOADING THE INFO-->
    <div v-if="store.loading" class=" flex flex-col p-10 justify-center animate-pulse">
        <div class="flex flex-col md:flex-row w-full max-w-[1400px] mx-auto px-4 items-start gap-4">
            <div class=" bg-gray-300 w-full md:w-2/6 min-h-[600px] shadow-2xs rounded-xl border-2"></div>
            <div class=" flex flex-col gap-5">
                <div class=" bg-gray-300 h-40 w-[45rem] shadow-2xs rounded-xl border-2"></div>
                <div class=" bg-gray-300 h-[30rem] w-[45rem] shadow-2xs rounded-xl border-2"></div>
            </div>
        </div>
    </div>
    <!--WHEN GETTING AN ERROR-->
    <ErrorComponent v-else-if="store.errorMessage" :ErrorMessage="store.errorMessage.message" :ErrorCode="store.errorMessage.code"/>
    <!--RENDERING ACTUAL STOCK CONTENT-->
    <div v-else class=" flex flex-col p-10 justify-center">
        <div class="flex flex-col md:flex-row w-full px-4 lg:px-20 xl:px-32 2xl:px-48 items-start gap-8">
            <img class="w-full md:w-2/6 h-auto  object-contain rounded-lg border" :src="store.stockInfo?.companyProfile.logo"/>
            <div class=" flex flex-col">
                <h1 class=" text-6xl font-bold">{{ store.stockInfo?.stock.ticker }}</h1>
                <!--TAB NAVIGATON MENU-->
                <div class=" flex flex-wrap gap-2 border-b-2 border-gray-400 w-full mt-5">
                    <button v-for="tab in store.tabs"
                        :key="tab.name"
                        class="px-4 py-1 text-sm font-medium rounded-t-md transition-all"
                        :class="{
                            'bg-gray-300 text-black border-b-4 border-black': store.tabSelected === tab.name,
                            'bg-gray-100 text-gray-700 hover:bg-gray-200': store.tabSelected !== tab.name
                        }"
                        @click="store.changeTab(tab)"
                        >
                        {{ tab.label }}
                    </button>
                </div>
                <!--CONTAINER OF TAB CHOOSEN-->
                <div v-if="store.tabSelected==='stock'" class="mt-5 space-y-3">
                    <h2 class="text-2xl font-bold mb-4">Stock Details</h2>
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-x-48 gap-y-4">
                        <p class="flex items-center gap-2">
                            <TrendingUp class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Ticker:</span> {{ store.stockInfo!.stock.ticker }}
                        </p>
                        <p class="flex items-center gap-2">
                            <Building2 class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Company:</span> {{ store.stockInfo!.stock.company }}
                        </p>
                        <p class="flex items-center gap-2">
                            <Landmark class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Brokerage:</span> {{ store.stockInfo!.stock.brokerage }}
                        </p>
                        <p class="flex items-center gap-2">
                            <BarChart2 class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Action:</span> {{ store.stockInfo!.stock.action }}
                        </p>
                        <p class="flex items-center gap-2">
                            <BarChart2 class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Rating:</span>
                            {{ store.stockInfo!.stock.ratingFrom }} → {{ store.stockInfo?.stock.ratingTo }}
                        </p>
                            <p class="flex items-center gap-2">
                            <Banknote class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Target Price:</span>
                            ${{ store.stockInfo!.stock.targetFrom.toFixed(2) }} → ${{ store.stockInfo?.stock.targetTo.toFixed(2) }}
                        </p>
                        <p class="flex items-center gap-2">
                            <Calendar class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Date:</span>
                            {{ new Date(store.stockInfo!.stock.time).toLocaleDateString() }}
                        </p>
                        <p class="flex items-center gap-2">
                            <Percent class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Percentage change:</span>
                            <span :class="{
                                'text-green-500': store.stockInfo!.stock.percentage > 0,
                                'text-red-500': store.stockInfo!.stock.percentage <= 0
                            }">
                                {{ (store.stockInfo!.stock.percentage > 0 ? '+' : '') + store.stockInfo!.stock.percentage.toFixed(2) + '%' }}
                            </span>
                        </p>
                    </div>
                </div>
                <div v-if="store.tabSelected === 'company' && store.stockInfo?.companyProfile" class="mt-5 space-y-3">
                    <h2 class="text-2xl font-bold mb-4">Company Profile</h2>
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-x-48 gap-y-4">
                        <p class="flex items-center gap-2">
                            <Building2 class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Name:</span> {{ store.stockInfo!.companyProfile.name }}
                        </p>
                        <p class="flex items-center gap-2">
                            <Globe class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Country:</span> {{ store.stockInfo!.companyProfile.country }}
                        </p>
                        <p class="flex items-center gap-2">
                            <DollarSign class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Currency:</span> {{ store.stockInfo!.companyProfile.currency }}
                        </p>
                        <p class="flex items-center gap-2">
                            <Landmark class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Exchange:</span> {{ store.stockInfo!.companyProfile.exchange }}
                        </p>
                        <p class="flex items-center gap-2">
                            <BarChart2 class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Industry:</span> {{ store.stockInfo!.companyProfile.industry }}
                        </p>
                        <p class="flex items-center gap-2">
                            <Calendar class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">IPO Date:</span> {{ new Date(store.stockInfo!.companyProfile.ipo).toLocaleDateString() }}
                        </p>
                        <p class="flex items-center gap-2">
                            <Banknote class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Market Capitalization:</span>
                            ${{ (store.stockInfo!.companyProfile.marketCapital / 1_000_000_000).toFixed(2) }}B
                        </p>
                        <p class="flex items-center gap-2">
                            <Users class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Shares Outstanding:</span> {{ store.stockInfo?.companyProfile.shareOurstanding.toLocaleString() }}
                        </p>
                        <p class="flex items-center gap-2">
                            <Phone class="w-4 h-4 text-gray-600" />
                            <span class="font-semibold">Phone:</span> {{ store.stockInfo?.companyProfile.phone }}
                        </p>
                        <p class="flex items-center gap-2">
                            <LinkIcon class="w-4 h-4 text-blue-400" />
                            <span class="font-semibold">Website:</span>
                            <a :href="store.stockInfo?.companyProfile.webUrl" class="text-blue-400 underline" target="_blank">
                                {{ store.stockInfo?.companyProfile.webUrl }}
                            </a>
                        </p>
                    </div>
                </div>
                <div v-if="store.tabSelected==='news'">
                    <div class=" border-gray-400 border-2 rounded-lg" >
                        <div v-for="news in store.stockInfo?.news || []"
                            class="flex flex-col md:flex-row gap-5 p-5 border-t border-gray-300"
                            >
                            <img v-if="news.image" class="w-36 h-24 object-cover rounded-lg" :src="news.image" />
                            <div class="flex flex-col gap-2">
                                <div class="flex items-center gap-2 text-xs text-gray-500">
                                    <span>{{ news.source }}</span>
                                    <span>•</span>
                                    <span>{{ new Date(news.date).toLocaleString('en-US', { dateStyle: 'medium', timeStyle: 'short' }) }}</span>
                                </div>
                                <h3 class="text-lg font-semibold text-gray-800">{{ news.headline }}</h3>
                                <p class="text-sm text-gray-700">{{ news.summary }}</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import ErrorComponent from '@/components/ErrorComponent.vue';
import RecommendationsComponent from '@/components/RecommendationsComponent.vue';
import {fullStockStore} from '@/stores/FullStockStore';
import { onMounted, watch } from 'vue';
import { useRoute } from 'vue-router';
import {
  Building2,
  Banknote,
  TrendingUp,
  Landmark,
  BarChart2,
  Users,
  Calendar,
  DollarSign,
  Percent,
  Globe,
  Phone,
  ArrowLeft ,
  Link as LinkIcon
} from 'lucide-vue-next'


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