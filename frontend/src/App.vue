<script setup>
import { ref } from 'vue'
import ThemeToggle from './components/ThemeToggle.vue'
import LanguageSelector from './components/LanguageSelector.vue'

const toggleValue = ref(false)
</script>

<template>
  <div>
    <div class="flex justify-between items-center p-3 z-50">
      <LanguageSelector />
      <ThemeToggle v-model="toggleValue" />
    </div>
    <router-view />
  </div>
</template>
