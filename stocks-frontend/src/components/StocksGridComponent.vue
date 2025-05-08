<template>
    <div class=" flex flex-col bg-gray-100 min-h-screen">
        <FiltersComponent/>
        <RecommendationsComponent/>
        <!--print stocks-->
        <div class=" flex flex-col flex-grow max-w-200 items-center">
            <!--print found and sort by options-->
            <div v-if="!store.loading && !store.errorMessage" class=" flex flex-row mt-6 mb-1  max-w-[1200px] justify-between w-full">
                <div class=" flex gap-1 text-gray-700 text-lg">
                    <p class=" font-semibold">Found </p>
                    <p>{{store.totalRecords}} stocks</p>
                </div>
                <div class=" group relative">
                    <button  @click="store.switchSortMenu"  class="px-4 py-2 rounded-md bg-white border border-gray-300 shadow-sm hover:bg-gray-100 text-gray-700 font-medium">
                        {{ store.selectedSort.label }}
                    </button>
                    <div v-if="store.openSort"
                     class="absolute right-5 top-full mt-2 z-50 w-64 divide-y divide-gray-300 rounded-lg border border-gray-300 bg-gray-50 shadow-md">
                        <div class="flex flex-col justify-center items-start gap-3">
                            <button  v-for="sort in store.sortOptions" :key="sort.value" class="w-full px-4 py-2 text-left text-gray-700 hover:bg-gray-100 transition-all" 
                            @click="store.selectSortType(sort)"
                            >{{ sort.label }}</button>
                        </div>
                    </div>
                </div>
            </div>
            <!--print error-->
            <ErrorComponent v-if="store.errorMessage" :ErrorMessage="store.errorMessage.message" :ErrorCode="store.errorMessage.code" />
            <!--print skeleton for when loading-->
            <div v-if="store.loading" class=" duration-100 animate-pulse mt-3 grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5 clear-both">
                <div class=" bg-gray-300 min-h-[280px] md:w-[400px] w-[250px] rounded-xl border border-gray-300 shadow animate-pulse " 
                v-for="index in 9" :key="index" >
                </div>
            </div>
            <!--print stocks-->
            <div v-if="!store.errorMessage && !store.loading" class="mt-3 grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
                <div 
                class="flex flex-col min-h-[150px] md:w-[400px] w-[250px] bg-white border border-gray-300 shadow-sm rounded-2xl transition hover:bg-gray-50 hover:shadow-lg transform  hover:scale-[1.02]" 
                v-for="stock in store.stocks" :key="stock.id"  @click="store.openStock(stock.id)">
                    <div class="p-4 md:p-5">
                        <div class=" flex flex-row justify-between">
                            <h3 class=" text-3xl font-bold text-gray-800">{{ stock.ticker }}</h3>
                            <div class=" flex flex-row">
                                <svg v-if="stock.percentage > 0" width="30" height="30" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" fill="none" stroke="green" stroke-width="4" stroke-linecap="round" stroke-linejoin="round">
                                    <line x1="12" y1="19" x2="12" y2="5"/>
                                    <polyline points="5 12 12 5 19 12"/>
                                </svg>
                                <svg v-else width="30" height="30" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" fill="none" stroke="red" stroke-width="4" stroke-linecap="round" stroke-linejoin="round">
                                    <line x1="12" y1="5" x2="12" y2="19"/>
                                    <polyline points="5 12 12 19 19 12"/>
                                </svg>
                                <h2 class= " text-2xl font-semibold" :class="{'text-green-600': stock.percentage>0, 'text-red-500': stock.percentage<=0}"> {{ (stock.percentage>0 ? '+': '') + stock.percentage.toFixed(2) + '%' }}</h2>
                            </div>
                        </div>
                        <div class="mt-3 space-y-2 text-sm text-gray-600">
                            <div class="flex items-center gap-2">
                                <Building2 class="w-4 h-4 text-gray-500" />
                                <span class="font-medium text-gray-700">Company:</span> {{ stock.company }}
                            </div>
                            <div class="flex items-center gap-2">
                                <Activity class="w-4 h-4 text-gray-500" />
                                <span class="font-medium text-gray-700">Action:</span> {{ stock.action }}
                            </div>
                            <div class="flex items-center gap-2">
                                <Landmark class="w-4 h-4 text-gray-500" />
                                <span class="font-medium text-gray-700">Brokerage:</span> {{ stock.brokerage }}
                            </div>
                            <div class="flex items-center gap-2">
                                <Clock3 class="w-4 h-4 text-gray-500" />
                                <span class="font-medium text-gray-700">Time:</span>
                                {{ stock.time.toLocaleString('en-US', { dateStyle: 'medium', timeStyle: 'short' }) }}
                            </div>
                        </div>

                        <hr class=" border-[1px] border-gray-400 mt-4"/>
                        <div class=" flex flex-row items-center justify-between gap-6 mt-4">
                            <div class=" flex flex-col items-center text-sm md:text-base font-medium text-gray-700 space-y-1">
                                <h3 :class="{
                                    'text-green-600' : store.greenRatings.includes(stock.ratingFrom),
                                    'text-yellow-600' : store.yellowRatings.includes(stock.ratingFrom),
                                    'text-orange-600' : store.orangeRatings.includes(stock.ratingFrom),
                                    'text-red-600' : store.redRatings.includes(stock.ratingFrom),
                                }">{{ stock.ratingFrom }}</h3>
                                <h3 class=" text-2xl font-semibold">{{ '$' + stock.targetFrom.toFixed(2) }}</h3>
                            </div>
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" width="30" height="30">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14M13 6l6 6-6 6" />
                            </svg>
                            <div class=" flex flex-col items-center text-sm md:text-base font-medium text-gray-700 space-y-1">
                                <h3 :class="{
                                    'text-green-600' : store.greenRatings.includes(stock.ratingTo),
                                    'text-yellow-600' : store.yellowRatings.includes(stock.ratingTo),
                                    'text-orange-600' : store.orangeRatings.includes(stock.ratingTo),
                                    'text-red-600' : store.redRatings.includes(stock.ratingTo),
                                }">{{ stock.ratingTo }}</h3>
                                <h3 class=" text-2xl font-semibold">{{ '$' + stock.targetTo.toFixed(2) }}</h3>
                            </div>
                        </div>
                      
                        <button  class=" w-full mt-4 py-2 rounded-lg border border-gray-300 text-gray-700 font-medium bg-white hover:bg-gray-100 transition" @click="store.openStock(stock.id)">
                            See more
                        </button>
                    </div>
                </div>
            </div>
            <!--pagination-->
            <div class=" bg-transparent pt-10 pb-6 mt-6 border-t border-gray-300 shadow-sm">
                <ul class=" flex items-center justify-center gap-0 md:gap-2">
                    <li>
                        <button @click="store.changePage(store.page-1)" class=" flex h-10 min-w-10 items-center justify-center rounded-lg border border-stroke border-gray-400 bg-white px-2 text-base font-medium hover:bg-gray-300">
                            <span>
                                <svg width="20" height="21" viewBox="0 0 20 21" fill="none" xmlns="http://www.w3.org/2000/svg">
                                    <path
                                        d="M17.5 9.8125H4.15625L9.46875 4.40625C9.75 4.125 9.75 3.6875 9.46875 3.40625C9.1875 3.125 8.75 3.125 8.46875 3.40625L2 9.96875C1.71875 10.25 1.71875 10.6875 2 10.9688L8.46875 17.5312C8.59375 17.6562 8.78125 17.75 8.96875 17.75C9.15625 17.75 9.3125 17.6875 9.46875 17.5625C9.75 17.2812 9.75 16.8438 9.46875 16.5625L4.1875 11.2188H17.5C17.875 11.2188 18.1875 10.9062 18.1875 10.5312C18.1875 10.125 17.875 9.8125 17.5 9.8125Z"
                                        fill="currentColor" />
                                </svg>
                            </span>
                        </button>
                    </li>
                    <li v-for="num in store.pages" :key="num">
                        <button @click="store.changePage(num)" v-if="num!==store.page" class="flex h-10 w-10 items-center justify-center rounded-md border border-gray-300 bg-white text-gray-700 hover:bg-gray-100 transition">
                            {{ num }}
                        </button>
                        <button @click="store.changePage(num)" v-else class="flex h-10 min-w-10 items-center justify-center rounded-lg border border-stroke px-2 text-base font-medium border-gray-400 text-dark hover:bg-gray-1 bg-gray-400">
                            {{ num }}
                        </button>
                    </li>
                    <li>
                        <button @click="store.changePage(store.page+1)" class="flex h-10 min-w-10 items-center justify-center rounded-lg border border-stroke border-gray-400 bg-white px-2 text-base font-medium text-dark hover:bg-gray-300">
                            <svg width="20" height="21" viewBox="0 0 20 21" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path
                                d="M18 10L11.5312 3.4375C11.25 3.15625 10.8125 3.15625 10.5312 3.4375C10.25 3.71875 10.25 4.15625 10.5312 4.4375L15.7812 9.78125H2.5C2.125 9.78125 1.8125 10.0937 1.8125 10.4688C1.8125 10.8438 2.125 11.1875 2.5 11.1875H15.8437L10.5312 16.5938C10.25 16.875 10.25 17.3125 10.5312 17.5938C10.6562 17.7188 10.8437 17.7812 11.0312 17.7812C11.2187 17.7812 11.4062 17.7188 11.5312 17.5625L18 11C18.2812 10.7187 18.2812 10.2812 18 10Z"
                                fill="currentColor" />
                            </svg>
                        </button>
                    </li>
                </ul>
            </div>
        </div>
    </div>

</template>


<script lan="ts" setup>
import FiltersComponent from './FiltersComponent.vue';
import { stocksStore } from '@/stores/StocksStore';
import RecommendationsComponent from './RecommendationsComponent.vue';
import ErrorComponent from './ErrorComponent.vue';
import { ArrowUpRight, ArrowDownRight, Building2, Activity, Landmark, Clock3, ArrowRight } from 'lucide-vue-next'

const store = stocksStore()
store.retrieveStocks()
</script>