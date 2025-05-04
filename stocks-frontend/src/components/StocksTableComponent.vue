<template>
    <RecommendationsComponent/>
    //CONTINUE HERE WITH PRINTING STOCKS
</template>


<script lan="ts" setup>
import RecommendationsComponent from './RecommendationsComponent.vue';
import { stocksStore } from '@/stores/StocksStore';
const store = stocksStore()
</script>