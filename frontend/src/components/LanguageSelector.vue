<template>
  <div class="relative inline-block text-left">
    <button
      @click="isOpen = !isOpen"
      class="inline-flex items-center px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm bg-white dark:bg-gray-800 text-sm font-medium text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-all duration-200"
    >
      <span class="mr-2 text-lg">🌐</span>
      {{ getCurrentLanguageName() }}
      <svg 
        class="-mr-1 ml-2 h-5 w-5 transform transition-transform duration-200" 
        :class="{ 'rotate-180': isOpen }"
        xmlns="http://www.w3.org/2000/svg" 
        viewBox="0 0 20 20" 
        fill="currentColor"
      >
        <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
      </svg>
    </button>

    <!-- Backdrop/Overlay -->
    <div 
      v-if="isOpen"
      @click="isOpen = false"
      class="fixed inset-0 z-40"
    ></div>

    <!-- Dropdown Menu -->
    <transition
      enter-active-class="transition duration-200 ease-out"
      enter-from-class="transform scale-95 opacity-0"
      enter-to-class="transform scale-100 opacity-100"
      leave-active-class="transition duration-150 ease-in"
      leave-from-class="transform scale-100 opacity-100"
      leave-to-class="transform scale-95 opacity-0"
    >
      <div
        v-show="isOpen"
        class="origin-top-left absolute left-0 mt-2 w-auto min-w-40 rounded-xl shadow-2xl bg-white dark:bg-gray-800 ring-1 ring-black ring-opacity-5 z-50 backdrop-blur-sm"
        style="box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);"
      >
        <div class="py-2">
          <button
            v-for="(lang, index) in languages"
            :key="lang.code"
            @click="changeLanguage(lang.code)"
            class="w-full text-left px-4 py-3 text-sm text-gray-700 dark:text-gray-300 hover:bg-blue-50 dark:hover:bg-blue-900/20 flex items-center transition-all duration-150 hover:scale-[1.01]"
            :class="{ 
              'bg-blue-50 dark:bg-blue-900/20 text-blue-600 dark:text-blue-400 font-medium': currentLocale === lang.code,
              'rounded-t-lg': index === 0,
              'rounded-b-lg': index === languages.length - 1
            }"
          >
            <span class="mr-3 text-lg">{{ lang.flag }}</span>
            <span class="flex-1">{{ lang.name }}</span>
            <span v-if="currentLocale === lang.code" class="ml-2 text-blue-500">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
              </svg>
            </span>
          </button>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useI18n } from 'vue-i18n'

const { locale, t } = useI18n()
const isOpen = ref(false)

const languages = [
  { code: 'en', name: 'English', flag: '🇺🇸' },
  { code: 'tr', name: 'Türkçe', flag: '🇹🇷' }
]

const currentLocale = computed(() => locale.value)

const getCurrentLanguageName = () => {
  const currentLang = languages.find(lang => lang.code === locale.value)
  return currentLang ? currentLang.name : 'English'
}

const changeLanguage = (langCode) => {
  locale.value = langCode
  localStorage.setItem('locale', langCode)
  isOpen.value = false
}

// Close dropdown when clicking outside or pressing Escape
const handleClickOutside = (event) => {
  if (!event.target.closest('.relative')) {
    isOpen.value = false
  }
}

const handleKeydown = (event) => {
  if (event.key === 'Escape') {
    isOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  document.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
  document.removeEventListener('keydown', handleKeydown)
})
</script>
