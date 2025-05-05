<template>
  <main>
    <RouterView></RouterView>
  </main>
  <footer class="relative mt-20 bg-gray-900 px-4 pt-20">
    <nav aria-label="Footer Navigation" class="mx-auto mb-10 flex max-w-lg flex-col gap-10 text-center sm:flex-row sm:text-left">
      <a href="#" class="font-medium text-white">Github Repository</a>
      <a href="#" class="font-medium text-white">Linkedin Profile</a>
      <a href="#" class="font-medium text-white">Ig Profile</a>
      <a href="#" class="font-medium text-white">Youtube</a>
    </nav>
    <p class="py-10 text-center text-gray-300">© 2025 Juan Diego Bermúdez Jiménez</p>
  </footer>

</template>

<script lan="ts" setup>
import {RouterView} from "vue-router";
</script>