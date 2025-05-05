<template>
    <div class=" flex flex-col">
        <FiltersComponent/>
        <RecommendationsComponent/>
        <!--print stocks-->
        <div class=" flex flex-col flex-grow max-w-200 items-center">
            <!--print stocks-->
            <div class="mt-3 grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
                <div 
                class="flex flex-col min-h-[150px] w-[400px] bg-white border-gray-600 shadow-2xs rounded-xl border-2 hover:bg-gray-200 hover:shadow-xl" 
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
                        <h3>{{'Company: ' +stock.company }}</h3>
                        <h3>{{'Action: ' +stock.action }}</h3>
                        <h3>{{'Brokerage: ' +stock.brokerage }}</h3>
                        <h3>{{'Time: ' +stock.time.toDateString() + ' ' + stock.time.getHours() + ':' + stock.time.getMinutes() }}</h3>
                        <hr class=" border-[1px] border-gray-400 mt-4"/>
                        <div class=" flex flex-row  items-center justify-center gap-4">
                            <div class=" flex flex-col items-center font-semibold text-lg my-2">
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
                            <div class=" flex flex-col items-center font-semibold text-lg my-2">
                                <h3 :class="{
                                    'text-green-600' : store.greenRatings.includes(stock.ratingTo),
                                    'text-yellow-600' : store.yellowRatings.includes(stock.ratingTo),
                                    'text-orange-600' : store.orangeRatings.includes(stock.ratingTo),
                                    'text-red-600' : store.redRatings.includes(stock.ratingTo),
                                }">{{ stock.ratingTo }}</h3>
                                <h3 class=" text-2xl font-semibold">{{ '$' + stock.targetTo.toFixed(2) }}</h3>
                            </div>
                        </div>
                      
                        <button class=" bg-gray-100 text-gray-800 shadow-md border-[1px] border-gray-600 hover:bg-gray-200 rounded-lg py-1 w-full" @click="store.openStock(stock.id)">
                            See more
                        </button>
                    </div>
                </div>
            </div>
            <!--pagination-->
            <div class=" bg-white pt-10 text-center">
                <ul class=" flex items-center justify-center gap-2">
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
                        <button @click="store.changePage(num)" v-if="num!==store.page" class="flex h-10 min-w-10 items-center justify-center rounded-lg border border-stroke border-gray-400 bg-white px-2 text-base font-medium text-dark hover:bg-gray-300">
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
const store = stocksStore()
store.retrieveStocks()
</script>