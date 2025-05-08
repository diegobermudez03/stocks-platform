<template>
    <div class=" bg-gray-700 grid grid-cols-1 md:grid-cols-2 gap-3 h-fit items-center">
        <!--search bar-->
        <div class=" flex items-center justify-center min-h-[10vh] px-5 bg-transparent">
            <div class=" flex rounded-full bg-white px-2 w-full max-w-[600px]">
                <input 
                    type="text"
                    class=" w-full  flex bg-transparent pl-2 text-black outline-0"
                    placeholder="Search for a stock or company"
                    v-model="store.searchQueryTmp"
                    @keydown.enter="store.submitSearchFilter"
                >
                <button type="submit" class=" relative p-2 bg-white rounded-full"  @click="store.submitSearchFilter">
                    <Search class="w-6 h-6 text-gray-800" />
                </button>
            </div>
        </div>
        <!--FILTERS-->
        <div class=" flex flex-wrap justify-center p-3 gap-10">
            <!-- Target price filter -->
            <div class="group relative">
                <button @click="store.toggleFilter('targetPrice')" 
                :class="store.openFilter === 'targetPrice' ? 'ring-2 ring-white' : ''"
                class="flex items-center gap-2 border font-bold border-gray-300 pb-1 transition-colors  [&::-webkit-details-marker]:hidden rounded-lg px-3 py-1 text-white shadow-md">
                    <SlidersHorizontal class="w-4 h-4 text-white" />
                    <span class="text-sm">Target Price Range</span>
                </button>
                <div v-if="store.openFilter === 'targetPrice'" 
                    class="absolute left-0 top-full mt-2 z-50 w-64 rounded-xl border border-gray-200 bg-white shadow-xl" >
                    <div class="flex items-center gap-3 p-3">
                        <label for="FromPrice">
                        <span class="text-sm text-gray-700"> From </span>
                        <input
                            type="number"
                            id="FromPrice"
                            value="0"
                            class="mt-0.5 w-full rounded-lg border-gray-600 border-[1px] shadow-md sm:text-sm px-2 py-1"
                            v-model="store.fromPriceTmp"
                        />
                        </label>
                        <label for="ToPrice">
                        <span class="text-sm text-gray-700"> To </span>
                        <input
                            type="number"
                            id="ToPrice"
                            value="10"
                            class="mt-0.5 w-full rounded-lg border-gray-600 border-[1px] shadow-md sm:text-sm px-2 py-1"
                            v-model="store.toPriceTmp"
                        />
                        </label>
                    </div>
                    <button v-if="store.fromPriceTmp && store.toPriceTmp" class=" px-5 py-2 bg-gray-700 text-white hover:bg-gray-900 rounded-lg m-3" @click="store.submitPriceRange">
                        Apply
                    </button>
                </div>
            </div>

             <!-- Rating from filter -->
            <div v-if="store.ratings.length > 0" class="group relative">
                    <button @click="store.toggleFilter('from-rating')" 
                    :class="store.openFilter === 'from-rating' ? 'ring-2 ring-white' : ''"
                    class="flex items-center gap-2 border font-bold border-gray-300 pb-1 transition-colors  [&::-webkit-details-marker]:hidden rounded-lg px-3 py-1 text-white shadow-md">
                        <Star class="w-4 h-4 text-white" />
                        <span class="text-sm">Rating From</span>
                    </button>
                    <div v-if="store.openFilter==='from-rating'"
                        class="absolute left-0 top-full mt-2 z-50 w-64 rounded-xl border border-gray-200 bg-white shadow-xl">
                        <fieldset class="p-3">
                            <div class="flex flex-col justify-center items-start gap-3">
                                <label v-for="rating in store.ratings" :key="rating.name" :for="rating.name" class="inline-flex items-center gap-3">
                                    <input type="checkbox" class="size-5 rounded-lg border-gray-300 shadow-sm" :id="rating.name"  :value="rating.name" v-model="store.selectedRatingsFrom"
                                    @change="store.retrieveStocks"/>
                                    <span class="text-sm font-medium text-gray-700"> {{rating.name + ' (' +rating.count + ')' }} </span>
                                </label>
                            </div>
                        </fieldset>
                    </div>
                </div>

             <!-- Rating to filter -->
             <div v-if="store.ratings.length > 0" class="group relative">
                <button @click="store.toggleFilter('to-rating')" 
                :class="store.openFilter === 'to-rating' ? 'ring-2 ring-white' : ''"
                class="flex items-center gap-2 border font-bold border-gray-300 pb-1 transition-colors  [&::-webkit-details-marker]:hidden rounded-lg px-3 py-1 text-white shadow-md">
                    <Star class="w-4 h-4 text-white" />
                    <span class="text-sm">Rating To</span>
                </button>
                <div v-if="store.openFilter==='to-rating'"
                    class="absolute left-0 top-full mt-2 z-50 w-64 rounded-xl border border-gray-200 bg-white shadow-xl">
                    <fieldset class="p-3">
                        <div class="flex flex-col justify-center items-start gap-3">
                            <label v-for="rating in store.ratings" :key="rating.name" :for="rating.name" class="inline-flex items-center gap-3">
                                <input type="checkbox" class="size-5 rounded-lg border-gray-300 shadow-sm" :id="rating.name" :value="rating.name" v-model="store.selectedRatingsTo"
                                @change="store.retrieveStocks"/>
                                <span class="text-sm font-medium text-gray-700"> {{rating.name + ' (' +rating.count + ')'}} </span>
                            </label>
                        </div>
                    </fieldset>
                </div>
            </div>

            <!-- Action filter -->
            <div v-if="store.actions.length> 0" class="group relative">
                <button @click="store.toggleFilter('actions')" 
                :class="store.openFilter === 'actions' ? 'ring-2 ring-white' : ''"
                class="flex items-center gap-2 border font-bold border-gray-300 pb-1 transition-colors  [&::-webkit-details-marker]:hidden rounded-lg px-3 py-1 text-white shadow-md">
                    <ArrowDownUp class="w-4 h-4 text-white" />
                    <span class="text-sm">Action</span>
                </button>
                <div v-if="store.openFilter==='actions'"
                    class="absolute left-0 top-full mt-2 z-50 w-64 rounded-xl border border-gray-200 bg-white shadow-xl">
                    <fieldset class="p-3">
                        <div class="flex flex-col justify-center items-start gap-3">
                            <label v-for="action in store.actions" :key="action.name" :for="action.name" class="inline-flex items-center gap-3" >
                                <input type="checkbox" class="size-5 rounded border-gray-300 shadow-sm" 
                                :id="action.name" :value="action.name" v-model="store.selectedActions" @change="store.retrieveStocks"/>
                                <span class="text-sm font-medium text-gray-700"> {{action.name  + ' (' +action.count + ')' }} </span>
                            </label>
                        </div>
                    </fieldset>
                </div>
            </div>
        </div>
    </div>
    <!--SELECTED FILTERS-->
    <div class=" w-full flex flex-wrap justify-center gap-x-5 gap-y-0 items-center bg-gray-700">
            <button class=" flex flex-row my-3 items-center bg-slate-900 text-white border-gray-500 border-2 shadow-md rounded-2xl px-5 py-2"
            v-for="filter in store.activeFilters" @click="()=>{filter.onRemove(); store.retrieveStocks()}">
                <Trash2 class="w-5 h-5 mr-2 text-red-400 hover:text-red-600" />
                {{ filter.label }}
            </button>
    </div>
</template>


<script lan="ts" setup>
import { stocksStore } from '@/stores/StocksStore';
import { Search, SlidersHorizontal, Filter, ArrowDownUp, Star, Trash2 } from 'lucide-vue-next'

const store = stocksStore()
store.getParams()
</script>