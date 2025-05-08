<template>
<div class="bg-white border border-gray-200 flex flex-col items-center justify-center px-4 md:px-8 lg:px-24 py-8 rounded-xl shadow-2xl text-center">
  <AlertTriangle class="w-16 h-16 text-yellow-500 mb-4 animate-pulse" />
  <p class="text-6xl md:text-7xl lg:text-9xl font-extrabold tracking-tight text-gray-300">
    {{ ErrorCode }}
  </p>
  <p class="mt-4 text-lg md:text-xl text-gray-700 font-medium max-w-md">
    {{ ErrorMessage }}.
  </p>
  <p class="mt-2 text-sm text-gray-500">
    Please try again later
  </p>
</div>

</template>


<script lang="ts" setup>
import { AlertTriangle } from 'lucide-vue-next'

defineProps({
    ErrorCode: Number,
    ErrorMessage: String
})
</script>