<template>
    <button @click="toggleTheme"
        class="relative inline-flex items-center h-8 rounded-full w-14 transition-colors focus:outline-none bg-gray-200 dark:bg-gray-700"
        :class="{ 'bg-primary-600': isDark }">
        <span class="inline-block flex items-center justify-center w-8 h-8 transform transition-transform rounded-full bg-white shadow-lg"
            :class="[isDark ? 'translate-x-6' : 'translate-x-0']">
            <mdicon v-if="!isDark" name="weather-sunny" class="text-yellow-600" />
            <mdicon v-else name="weather-night" class="text-indigo-600"/>
        </span>
    </button>
</template>

<script setup>
import { ref } from 'vue'

const isDark = ref(false)

const toggleTheme = () => {
    isDark.value = !isDark.value
    document.documentElement.classList.toggle('dark')
}
</script>
