<script setup>
import { onMounted, ref } from 'vue'
import ThemeToggle from './components/ThemeToggle.vue'

onMounted(() => {
  if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
    document.documentElement.classList.add('dark')
  }
})
const toggleValue = ref(false)
</script>

<template>
  <div class="container mx-auto p-8">
    <div class="flex justify-end mb-4">
      <ThemeToggle v-model="toggleValue" />
    </div>

    <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow">
      <h1 class="text-2xl font-bold mb-4 text-gray-900 dark:text-white">Password Manager</h1>
      <p class="text-gray-600 dark:text-gray-400">
        Dark mode destekli içerik
      </p>
    </div>
  </div>
</template>

<style>
:root {
  --bg-transition-time: 200ms;
}

* {
  transition: background-color var(--bg-transition-time) ease-in-out;
}
</style>
