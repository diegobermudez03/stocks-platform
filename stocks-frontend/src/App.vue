<template>
  <main>
    <RouterView></RouterView>
  </main>
  <footer class="relative mt-20 bg-gray-900 px-4 pt-20">
    <nav aria-label="Footer Navigation" class="mx-auto mb-10 flex max-w-lg flex-col gap-6 text-center sm:flex-row sm:text-left text-white">
    <a href="https://github.com/diegobermudez03/stocks-platform" class="flex items-center justify-center sm:justify-start gap-2 hover:text-gray-300 transition">
      <Github class="w-5 h-5" />
        Github Repository
    </a>
    <a href="https://www.linkedin.com/in/dbermu/" class="flex items-center justify-center sm:justify-start gap-2 hover:text-gray-300 transition">
      <Linkedin class="w-5 h-5" />
      LinkedIn Profile
    </a>
    <a href="https://www.instagram.com/diego_bermudezjimenez/" class="flex items-center justify-center sm:justify-start gap-2 hover:text-gray-300 transition">
      <Instagram class="w-5 h-5" />
      IG Profile
    </a>
    <a href="https://www.youtube.com/@johndb7861" class="flex items-center justify-center sm:justify-start gap-2 hover:text-gray-300 transition">
      <Youtube class="w-5 h-5" />
      YouTube
    </a>
  </nav>
  <p class="py-10 text-center text-gray-300">© 2025 Juan Diego Bermúdez Jiménez</p>
  </footer>

</template>

<script lan="ts" setup>
import { RouterView } from 'vue-router'

import { Github , Linkedin, Instagram, Youtube } from 'lucide-vue-next'
</script>