<template>
    <h1 class=" z-20 bg-gray-900 text-white text-center py-3 text-xl font-semibold shadow-md">
      Our Top Recommendations
    </h1>
    <!--when loading-->
    <div v-if="store.loading && !store.errorMessage" class="w-full sticky top-0 z-10 bg-gray-900 h-32">
      <div class="flex space-x-2 justify-center items-center h-full">
        <div class="w-4 h-4 bg-white rounded-full animate-bounce [animation-delay:-0.3s]"></div>
        <div class="w-4 h-4 bg-white rounded-full animate-bounce [animation-delay:-0.15s]"></div>
        <div class="w-4 h-4 bg-white rounded-full animate-bounce"></div>
        <div class="w-4 h-4 bg-white rounded-full animate-bounce [animation-delay:0.15s]"></div>
      </div>
    </div>
    <!--when error-->
    <div v-else-if="store.errorMessage" class=" bg-gray-900 h-full items-center justify-center">
      <h3 class=" text-white text-center justify-center text-2xl h-full">{{'Unable to fetch recommendations: ' + store.errorMessage}}</h3>
    </div>
    <!--when the recommendations-->
    <div v-else ref="scrollContainer" @mouseenter="pauseScroll" @mouseleave="startScroll" class=" overflow-x-auto w-full sticky top-0 z-10">
        <div class=" flex flex-row spacex-4 min-w-max">
            <button v-if="store.recommendations.length > 0" v-for="recommendation in store.recommendations" :key="recommendation.id"  @click="store.openStock(recommendation.id)"
              class="flex flex-col bg-gray-900 text-white px-6 py-4 border border-gray-700 w-[350px] hover:bg-gray-800 hover:shadow-lg transition-all duration-300 hover:scale-[1.02] transform">
                <div class=" flex flex-row justify-between">
                  <h2 class=" text-2xl font-bold">{{ recommendation.ticker }}</h2>
                  <div class=" flex flex-col ">
                    <div class="flex items-center gap-1" :class="{'text-green-500': recommendation.percentageIncrease > 0, 'text-red-500': recommendation.percentageIncrease <= 0}">
                      <TrendingUp class="w-4 h-4" />
                      <span class="font-semibold">{{ (recommendation.percentageIncrease > 0 ? '+' : '') + recommendation.percentageIncrease.toFixed(2) + '%' }}</span>
                    </div>
                    <div class="flex items-center gap-1">
                      <DollarSign class="w-4 h-4 text-gray-300" />
                      <span>{{ recommendation.targetTo.toFixed(2) }}</span>
                    </div>

                  </div>
                </div>
                <div class=" flex flex-row items-center justify-between">
                  <div class=" flex flex-col">
                    <div class="flex items-center gap-1 text-gray-400">
                      <Gauge class="w-4 h-4" />
                      <span>Score</span>
                    </div>
                    <h3 class="text-lg font-medium" :class="{
                      'text-green-500' : recommendation.recommendationScore >= 0.5, 
                      'text-yellow-500': recommendation.recommendationScore >= 0.3 && recommendation.recommendationScore < 0.5,
                      'text-orange-500' : recommendation.recommendationScore >= 0.1 && recommendation.recommendationScore < 0.3,
                      'text-red-600' : recommendation.recommendationScore < 0.1
                    }">{{recommendation.recommendationScore.toFixed(2) }}</h3>
                  </div>
                  <div class=" flex flex-col">
                    <div class="flex items-center gap-1 text-gray-400">
                      <Users class="w-4 h-4" />
                      <span>Insider Sentiment</span>
                    </div>
                    <h3 class="text-lg font-medium" :class="{
                      'text-green-500' : recommendation.avrgSentiment >= 60, 
                      'text-yellow-500': recommendation.avrgSentiment >= 30 && recommendation.avrgSentiment < 60,
                      'text-orange-500' : recommendation.avrgSentiment >= 10 && recommendation.avrgSentiment < 30,
                      'text-red-600' : recommendation.avrgSentiment < 10
                    }">{{recommendation.avrgSentiment.toFixed(2) }}</h3>
                  </div>
                </div>
            </button>
        </div>
    </div>
</template>


<script lang="ts" setup>
import { recommendationsStore } from '@/stores/RecommendationsStore';
import { TrendingUp, DollarSign, Gauge, Users } from 'lucide-vue-next'
import { onBeforeUnmount, onMounted, ref } from 'vue';
import { watch, nextTick } from 'vue';
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

//watch and only start animation until our component has items (which means, is rendered)
watch(
  () => store.recommendations.length,
  async (newVal) => {
    if (newVal) {
      await nextTick() 
      startScroll()
    }
  },
  { immediate: true } 
)

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