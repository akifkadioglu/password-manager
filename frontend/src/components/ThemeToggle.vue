<template>
    <button @click="toggleTheme"
        class="relative inline-flex items-center h-8 rounded-full w-14 transition-colors focus:outline-none"
        :class="isDark ? 'bg-primary-600' : 'bg-zinc-200'">
        <span class="inline-block flex items-center justify-center w-8 h-8 transform transition-transform rounded-full bg-white shadow-lg"
            :class="[isDark ? 'translate-x-6' : 'translate-x-0']">
            <mdicon v-if="!isDark" name="weather-sunny" class="text-accent-400" />
            <mdicon v-else name="weather-night" class="text-primary-600"/>
        </span>
    </button>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const isDark = ref(false)

// Sayfa yüklendiğinde mevcut tema durumunu kontrol et
onMounted(() => {
    const currentTheme = document.documentElement.getAttribute('data-theme')
    isDark.value = currentTheme === 'dark'
    
    // Eğer tema atanmamışsa, sistem tercihine bak
    if (!currentTheme) {
        const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
        isDark.value = prefersDark
        document.documentElement.setAttribute('data-theme', prefersDark ? 'dark' : 'light')
    }
})

const toggleTheme = () => {
    isDark.value = !isDark.value
    const newTheme = isDark.value ? 'dark' : 'light'
    document.documentElement.setAttribute('data-theme', newTheme)
    
    // localStorage'a kaydet
    localStorage.setItem('theme', newTheme)
}
</script>
