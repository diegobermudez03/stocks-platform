<template>
    <div class=" bg-gray-700 flex flex-col gap-3 h-fit">
        <div class=" flex items-center justify-center min-h-[10vh] w-screen px-5 bg-transparent">
            <div class=" flex rounded-full bg-white px-2 w-full max-w-[600px]">
                <input 
                    type="text"
                    class=" w-full  flex bg-transparent pl-2 text-black outline-0"
                    placeholder="Search for a stock or company"
                >
                <button type="submit" class=" relative p-2 bg-white rounded-full" >
                    <svg width="30px" height="30px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <g id="SVGRepo_bgCarrier" stroke-width="0"/>
                    <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"/>
                    <g id="SVGRepo_iconCarrier"> <path d="M14.9536 14.9458L21 21M17 10C17 13.866 13.866 17 10 17C6.13401 17 3 13.866 3 10C3 6.13401 6.13401 3 10 3C13.866 3 17 6.13401 17 10Z" stroke="#353535" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/> </g>
                    </svg>
                </button>
            </div>
        </div>

        <div class=" flex flex-row justify-center p-3 gap-10">
            <!-- Target price filter -->
            <details class="group relative">
                <summary class="flex items-center gap-2 border font-bold border-gray-300 pb-1 transition-colors  [&::-webkit-details-marker]:hidden rounded px-3 py-1 text-white shadow-md">
                    <span class="text-sm "> Target Price Range </span>
                    <span class="transition-transform group-open:-rotate-180">
                        <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke-width="1.5"
                        stroke="currentColor"
                        class="size-4">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
                        </svg>
                    </span>
                </summary>
                <div class=" flex flex-col justify-center z-auto w-64 divide-y divide-gray-300 rounded border border-gray-300 bg-gray-50 shadow-sm group-open:absolute group-open:start-0 group-open:top-8" >
                    <div class="flex items-center gap-3 p-3">
                        <label for="MinPrice">
                        <span class="text-sm text-gray-700"> From </span>
                        <input
                            type="number"
                            id="MinPrice"
                            value="0"
                            class="mt-0.5 w-full rounded border-gray-300 shadow-md sm:text-sm px-2 py-1"
                        />
                        </label>
                        <label for="MaxPrice">
                        <span class="text-sm text-gray-700"> To </span>
                        <input
                            type="number"
                            id="MaxPrice"
                            value="10"
                            class="mt-0.5 w-full rounded border-gray-300 shadow-md sm:text-sm px-2 py-1"
                        />
                        </label>
                    </div>
                    <button class=" px-5 py-2 bg-gray-700 text-white hover:bg-gray-900 rounded-md m-3">
                        Apply
                    </button>
                </div>
            </details>

             <!-- Rating from filter -->
            <details v-if="store.ratings.length > 0" class="group relative">
                    <summary class="flex items-center gap-2 border font-bold border-gray-300 pb-1 transition-colors  [&::-webkit-details-marker]:hidden rounded px-3 py-1 text-white shadow-md">
                        <span class="text-sm font-medium"> Rating From </span>
                        <span class="transition-transform group-open:-rotate-180">
                            <svg
                            xmlns="http://www.w3.org/2000/svg"
                            fill="none"
                            viewBox="0 0 24 24"
                            stroke-width="1.5"
                            stroke="currentColor"
                            class="size-4"
                            >
                            <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
                            </svg>
                        </span>
                    </summary>
                    <div class="z-auto w-64 divide-y divide-gray-300 rounded border border-gray-300 bg-white shadow-sm group-open:absolute group-open:start-0 group-open:top-8">
                        <fieldset class="p-3">
                            <div class="flex flex-col justify-center items-start gap-3">
                                <label v-for="rating in store.ratings" :key="rating.name" :for="rating.name" class="inline-flex items-center gap-3">
                                    <input type="checkbox" class="size-5 rounded border-gray-300 shadow-sm" :id="rating.name" />

                                    <span class="text-sm font-medium text-gray-700"> {{rating.name}} </span>
                                </label>
                                <button class=" px-5 py-2 bg-gray-700 text-white hover:bg-gray-900 rounded-md m-3">
                                    Apply
                                </button>
                            </div>
                        </fieldset>
                    </div>
            </details>

             <!-- Rating to filter -->
             <details v-if="store.ratings.length > 0" class="group relative">
                <summary class="flex items-center gap-2 border font-bold border-gray-300 pb-1 transition-colors  [&::-webkit-details-marker]:hidden rounded px-3 py-1 text-white shadow-md">
                    <span class="text-sm font-medium"> Rating To </span>
                    <span class="transition-transform group-open:-rotate-180">
                        <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke-width="1.5"
                        stroke="currentColor"
                        class="size-4"
                        >
                        <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
                        </svg>
                    </span>
                </summary>
                <div class="z-auto w-64 divide-y divide-gray-300 rounded border border-gray-300 bg-white shadow-sm group-open:absolute group-open:start-0 group-open:top-8">
                    <fieldset class="p-3">
                        <div class="flex flex-col justify-center items-start gap-3">
                            <label v-for="rating in store.ratings" :key="rating.name" :for="rating.name" class="inline-flex items-center gap-3">
                                <input type="checkbox" class="size-5 rounded border-gray-300 shadow-sm" :id="rating.name" />
                                <span class="text-sm font-medium text-gray-700"> {{rating.name}} </span>
                            </label>
                            <button class=" px-5 py-2 bg-gray-700 text-white hover:bg-gray-900 rounded-md m-3">
                                Apply
                            </button>
                        </div>
                    </fieldset>
                </div>
            </details>

            <!-- Action filter -->
            <details v-if="store.actions.length> 0" class="group relative">
                <summary class="flex items-center gap-2 border font-bold border-gray-300 pb-1 transition-colors  [&::-webkit-details-marker]:hidden rounded px-3 py-1 text-white shadow-md">
                    <span class="text-sm font-medium"> Action </span>
                    <span class="transition-transform group-open:-rotate-180">
                        <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke-width="1.5"
                        stroke="currentColor"
                        class="size-4"
                        >
                        <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
                        </svg>
                    </span>
                </summary>
                <div class="z-auto w-64 divide-y divide-gray-300 rounded border border-gray-300 bg-white shadow-sm group-open:absolute group-open:start-0 group-open:top-8">
                    <fieldset class="p-3">
                        <div class="flex flex-col justify-center items-start gap-3">
                            <label v-for="action in store.actions" :key="action.name" :for="action.name" class="inline-flex items-center gap-3">
                                <input type="checkbox" class="size-5 rounded border-gray-300 shadow-sm" :id="action.name" />
                                <span class="text-sm font-medium text-gray-700"> {{action.name}} </span>
                            </label>
                            <button class=" px-5 py-2 bg-gray-700 text-white hover:bg-gray-900 rounded-md m-3">
                                Apply
                            </button>
                        </div>
                    </fieldset>
                </div>
            </details>
        </div>
    </div>
</template>


<script lan="ts" setup>
import { stocksStore } from '@/stores/StocksStore';
const store = stocksStore()
store.getParams()
</script>