<script setup>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import ThemeToggle from './ThemeToggle.vue'
import LanguageSelector from './LanguageSelector.vue'
import { 
  mdiHome,
  mdiShieldKey,
  mdiFormatListBulleted,
  mdiMenu
} from '@mdi/js'

const route = useRoute()
const mobileMenuOpen = ref(false)
const toggleValue = ref(false)

// Navigation items configuration with MDI icons
const navItems = [
  {
    name: 'home',
    path: '/',
    routeName: 'Landing',
    icon: mdiHome
  },
  {
    name: 'passwordGenerator',
    path: '/password-generator',
    routeName: 'PasswordGenerator',
    icon: mdiShieldKey
  },
  {
    name: 'passwords',
    path: '/passwords',
    routeName: 'PasswordsList',
    icon: mdiFormatListBulleted
  }
]

const closeMobileMenu = () => {
  mobileMenuOpen.value = false
}

const isActiveRoute = (routeName) => {
  return route.name === routeName
}
</script>

<template>
  <nav class="bg-white/95 dark:bg-zinc-900/95 backdrop-blur-xl shadow-lg border-b border-zinc-200/50 dark:border-zinc-700/50 sticky top-0 z-50">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-16">
        <!-- Logo/Brand Section -->
        <div class="flex items-center">
          <router-link 
            to="/" 
            class="flex items-center space-x-3 hover:opacity-90 transition-all duration-200 group"
            @click="closeMobileMenu"
          >
            <!-- Logo Container -->
            <div class="relative">
              <div class="w-12 h-12 bg-gradient-to-br from-sky-500 via-sky-500 to-sky-600 rounded-xl shadow-lg group-hover:shadow-xl transition-all duration-300 flex items-center justify-center">
                <!-- Owl Icon -->
                <img src="../assets/images/icon.png" class="rounded-xl" alt="">
              </div>
              <!-- Status Indicator -->
              <div class="absolute -top-1 -right-1 w-4 h-4 bg-emerald-500 rounded-full border-2 border-white dark:border-zinc-900 animate-pulse shadow-sm"></div>
            </div>
            
            <!-- Brand Text -->
            <div class="hidden sm:block">
              <h1 class="text-xl font-bold bg-gradient-to-r from-sky-600 via-sky-600 to-sky-700 bg-clip-text text-transparent">
                OwlLock
              </h1>
              <p class="text-xs text-zinc-500 dark:text-zinc-400 -mt-1 font-medium">
                Secure & Vigilant
              </p>
            </div>
          </router-link>
        </div>

        <!-- Desktop Navigation -->
        <div class="hidden md:flex items-center space-x-1">
          <router-link 
            v-for="item in navItems"
            :key="item.name"
            :to="item.path"
            class="flex items-center space-x-2 px-4 py-2 rounded-xl text-sm font-medium transition-all duration-200 hover:scale-105"
            :class="[
              isActiveRoute(item.routeName)
                ? 'text-amber-700 dark:text-amber-400 bg-gradient-to-r from-amber-50 to-orange-50 dark:from-amber-900/30 dark:to-orange-900/30 shadow-md border border-amber-200/50 dark:border-amber-800/50'
                : 'text-zinc-600 dark:text-zinc-300 hover:text-amber-600 dark:hover:text-amber-400 hover:bg-zinc-50/80 dark:hover:bg-zinc-800/50'
            ]"
          >
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
              <path :d="item.icon" />
            </svg>
            <span>{{ $t(`navigation.${item.name}`) }}</span>
          </router-link>
        </div>

        <!-- Controls Section -->
        <div class="flex items-center space-x-3">
          <!-- Desktop Controls -->
          <div class="hidden md:flex items-center space-x-3">
            <div class="h-6 w-px bg-zinc-300 dark:bg-zinc-600"></div>
            <LanguageSelector />
            <ThemeToggle v-model="toggleValue" />
          </div>
          
          <!-- Mobile Menu Button -->
          <button
            @click="mobileMenuOpen = !mobileMenuOpen"
            class="md:hidden relative p-2 rounded-xl text-zinc-600 dark:text-zinc-300 hover:text-amber-600 dark:hover:text-amber-400 hover:bg-zinc-100/80 dark:hover:bg-zinc-800/80 transition-all duration-200"
            :class="{ 'text-amber-600 dark:text-amber-400': mobileMenuOpen }"
          >
            <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
              <path :d="mdiMenu" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Mobile Navigation Menu -->
    <div 
      class="md:hidden border-t border-zinc-200/50 dark:border-zinc-700/50 bg-white/98 dark:bg-zinc-900/98 backdrop-blur-xl"
      :class="mobileMenuOpen ? 'block' : 'hidden'"
    >
      <div class="max-w-7xl mx-auto px-4 py-4 space-y-2">
        <!-- Mobile Navigation Links -->
        <div class="space-y-1">
          <router-link 
            v-for="item in navItems"
            :key="item.name"
            :to="item.path"
            @click="closeMobileMenu"
            class="flex items-center space-x-3 px-4 py-3 rounded-xl text-sm font-medium transition-all duration-200"
            :class="[
              isActiveRoute(item.routeName)
                ? 'text-amber-700 dark:text-amber-400 bg-gradient-to-r from-amber-50 to-orange-50 dark:from-amber-900/30 dark:to-orange-900/30 shadow-md border border-amber-200/50 dark:border-amber-800/50'
                : 'text-zinc-600 dark:text-zinc-300 hover:text-amber-600 dark:hover:text-amber-400 hover:bg-zinc-50/80 dark:hover:bg-zinc-800/50'
            ]"
          >
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
              <path :d="item.icon" />
            </svg>
            <span>{{ $t(`navigation.${item.name}`) }}</span>
          </router-link>
        </div>
        
        <!-- Mobile Controls -->
        <div class="border-t border-zinc-200/50 dark:border-zinc-700/50 pt-4 mt-4">
          <div class="flex items-center justify-between px-4 py-2">
            <span class="text-sm font-medium text-zinc-500 dark:text-zinc-400">
              Settings
            </span>
            <div class="flex items-center space-x-3">
              <LanguageSelector />
              <ThemeToggle v-model="toggleValue" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<style scoped>
/* Additional animations and transitions */
.router-link-active {
  transition-property: all;
  transition-duration: 200ms;
}

/* Custom scrollbar for mobile menu if needed */
.mobile-menu::-webkit-scrollbar {
  width: 4px;
}

.mobile-menu::-webkit-scrollbar-track {
  background-color: transparent;
}

.mobile-menu::-webkit-scrollbar-thumb {
  background-color: #d4d4d8;
  border-radius: 9999px;
}

.dark .mobile-menu::-webkit-scrollbar-thumb {
  background-color: #52525b;
}
</style>
