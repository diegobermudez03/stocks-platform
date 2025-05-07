<template>
    <h1 class=" bg-gray-900 text-gray-300 text-center pt-2 font-semibold text-lg">Our top recommendations</h1>
    <div ref="scrollContainer" @mouseenter="pauseScroll" @mouseleave="startScroll" class=" overflow-x-auto w-full sticky top-0">
        <div class=" flex flex-row spacex-4 min-w-max">
            <button v-if="store.recommendations.length > 0" v-for="recommendation in store.recommendations" :key="recommendation.id"  @click="store.openStock(recommendation.id)"
                class=" flex flex-col bg-gray-900 text-white px-10 border-x-[1px] border-gray-500 w-[400px] py-3 hover:bg-gray-950">
                <div class=" flex flex-row justify-between">
                  <h2 class=" text-2xl font-bold">{{ recommendation.ticker }}</h2>
                  <div class=" flex flex-col ">
                    <h2 class= " text-xl font-semibold" :class="{'text-green-600': recommendation.percentage_increase>0, 'text-red-500': recommendation.percentage_increase<=0}"> {{ (recommendation.percentage_increase>0 ? '+': '') + recommendation.percentage_increase.toFixed(2) + '%' }}</h2>
                    <h2 :class="{'text-green-600': recommendation.percentage_increase>0, 'text-red-500': recommendation.percentage_increase<=0}" >{{ '$' + recommendation.target_to.toFixed(2) }}</h2>
                  </div>
                </div>
                <div class=" flex flex-row items-center justify-between">
                  <div class=" flex flex-col">
                    <h3>Score (calculated)</h3>
                    <h3 class="text-lg" :class="{
                      'text-green-500' : recommendation.recommendation_score >= 0.85, 
                      'text-yellow-500': recommendation.recommendation_score >= 0.65 && recommendation.recommendation_score < 0.85,
                      'text-orange-500' : recommendation.recommendation_score >= 0.40 && recommendation.recommendation_score < 0.65,
                      'text-red-600' : recommendation.recommendation_score < 0.40
                    }">{{recommendation.recommendation_score }}</h3>
                  </div>
                  <div class=" flex flex-col">
                    <h3>Insider sentiment</h3>
                    <h3 class=" text-lg" :class="{
                      'text-green-500' : recommendation.recommendation_score >= 85, 
                      'text-yellow-500': recommendation.recommendation_score >= 65 && recommendation.recommendation_score < 85,
                      'text-orange-500' : recommendation.recommendation_score >= 40 && recommendation.recommendation_score < 65,
                      'text-red-600' : recommendation.recommendation_score < 40
                    }">{{recommendation.avrg_sentiment }}</h3>
                  </div>
                </div>
            </button>
        </div>
    </div>
</template>


<script lang="ts" setup>
import { recommendationsStore } from '@/stores/RecommendationsStore';
import { onBeforeUnmount, onMounted, ref } from 'vue';
const store = recommendationsStore()
store.retrieveRecommendations()

const scrollContainer = ref<HTMLElement | null>(null)
let scrollInterval: number | null = null

function pauseScroll(){
  if(scrollInterval)clearInterval(scrollInterval)
}

function startScroll(){
  if (!scrollContainer.value) return
  scrollInterval = setInterval(() => {
    const el = scrollContainer.value
    el!.scrollLeft += 1
    if (el!.scrollLeft >= el!.scrollWidth - el!.clientWidth) {
      el!.scrollLeft = 0 
    }
  }, 20) 
}

onMounted(startScroll)
onBeforeUnmount(pauseScroll)

</script>


<style>
@keyframes scroll {
  0% {
    transform: translateX(0%);
  }
  100% {
    transform: translateX(-50%);
  }
}

.animate-scroll {
  animation: scroll 30s linear infinite;
}
</style>