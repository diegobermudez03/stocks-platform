<template>
    <div class=" grid grid-cols-2 md:grid-cols-1 max-h-[80vh]">
        <h1>Filters</h1>
        <input placeholder="search">
        <h2>Actions</h2>
        <div v-if="store.loadingFilter">
            loading
        </div>
        <div v-else-if="store.filterError">
            error
        </div>
        <div v-else class=" flex flex-col items-start gap-3">
            <label v-for="action in store.actions" class=" inline-flex items-center gap-3">
                <input type="checkbox" class=" size-5 rounded border-gray-300 shadow-sm text-blue-500" :id="action.name">
                <span class=" font-medium text-gray-700">{{ action.name + " " +action.count }}</span>
            </label>
        </div>
        <div class=" flex flex-row gap-5">
            <button>
                Apply
            </button>
            <button>
                Reset
            </button>
        </div>
    </div>
</template>


<script lan="ts" setup>
import { stocksStore } from '@/stores/StocksStore';
const store = stocksStore()
store.getParams()
</script>