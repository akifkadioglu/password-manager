<template>
  <div class="min-h-screen py-8">
    <div class="container mx-auto px-4">
      <div class="max-w-6xl mx-auto">
        <!-- Header -->
        <div class="text-center mb-8">
          <div class="flex justify-center mb-6">
            <div class="p-4 bg-gradient-to-br from-primary-600 to-primary-700 rounded-2xl shadow-2xl">
              <svg class="w-12 h-12 text-white" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12,1L21,5V9C21,16.55 15.84,23.74 9,23.74C2.16,23.74 -3,16.55 -3,9V5L12,1M12,7A2,2 0 0,0 10,9A2,2 0 0,0 12,11A2,2 0 0,0 14,9A2,2 0 0,0 12,7Z" />
              </svg>
            </div>
          </div>
          <h1 class="text-4xl font-bold bg-gradient-to-r from-primary-600 via-primary-700 to-primary-800 bg-clip-text text-transparent mb-4">
            {{ $t('passwords.title') }}
          </h1>
          <p class="text-lg text-zinc-600 dark:text-zinc-300 mb-6">
            {{ $t('passwords.subtitle') }}
          </p>
          
          <!-- Actions -->
          <div class="flex flex-col sm:flex-row items-center justify-center gap-4">
            <router-link to="/password-generator">
              <button class="flex items-center space-x-2 bg-gradient-to-r from-primary-600 to-primary-700 hover:from-primary-700 hover:to-primary-800 text-white px-6 py-3 rounded-xl shadow-lg hover:shadow-xl transition-all duration-300 transform hover:scale-105">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M19,13H13V19H11V13H5V11H11V5H13V11H19V13Z" />
                </svg>
                <span>{{ $t('passwords.addNew') }}</span>
              </button>
            </router-link>
            <button
              @click="refreshPasswords"
              :disabled="loading"
              class="flex items-center space-x-2 bg-zinc-100 dark:bg-zinc-700 hover:bg-zinc-200 dark:hover:bg-zinc-600 text-zinc-700 dark:text-zinc-300 px-6 py-3 rounded-xl transition-all duration-300"
            >
              <svg 
                class="w-5 h-5" 
                :class="{ 'animate-spin': loading }"
                fill="currentColor" 
                viewBox="0 0 24 24"
              >
                <path d="M17.65,6.35C16.2,4.9 14.21,4 12,4A8,8 0 0,0 4,12A8,8 0 0,0 12,20C15.73,20 18.84,17.45 19.73,14H17.65C16.83,16.33 14.61,18 12,18A6,6 0 0,1 6,12A6,6 0 0,1 12,6C13.66,6 15.14,6.69 16.22,7.78L13,11H20V4L17.65,6.35Z" />
              </svg>
              <span>{{ $t('passwords.refresh') }}</span>
            </button>
          </div>
        </div>

        <!-- Stats Cards -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
          <div class="bg-white dark:bg-zinc-800 rounded-2xl p-6 shadow-lg border border-zinc-200 dark:border-zinc-700">
            <div class="flex items-center space-x-4">
              <div class="p-3 bg-blue-100 dark:bg-blue-900/30 rounded-xl">
                <svg class="w-6 h-6 text-blue-600 dark:text-blue-400" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12,1L21,5V9C21,16.55 15.84,23.74 9,23.74C2.16,23.74 -3,16.55 -3,9V5L12,1Z" />
                </svg>
              </div>
              <div>
                <p class="text-2xl font-bold text-zinc-900 dark:text-zinc-100">{{ passwords.length }}</p>
                <p class="text-sm text-zinc-600 dark:text-zinc-400">{{ $t('passwords.totalPasswords') }}</p>
              </div>
            </div>
          </div>

          <div class="bg-white dark:bg-zinc-800 rounded-2xl p-6 shadow-lg border border-zinc-200 dark:border-zinc-700">
            <div class="flex items-center space-x-4">
              <div class="p-3 bg-green-100 dark:bg-green-900/30 rounded-xl">
                <svg class="w-6 h-6 text-green-600 dark:text-green-400" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z" />
                </svg>
              </div>
              <div>
                <p class="text-2xl font-bold text-zinc-900 dark:text-zinc-100">{{ secureCount }}</p>
                <p class="text-sm text-zinc-600 dark:text-zinc-400">{{ $t('passwords.securePasswords') }}</p>
              </div>
            </div>
          </div>

          <div class="bg-white dark:bg-zinc-800 rounded-2xl p-6 shadow-lg border border-zinc-200 dark:border-zinc-700">
            <div class="flex items-center space-x-4">
              <div class="p-3 bg-purple-100 dark:bg-purple-900/30 rounded-xl">
                <svg class="w-6 h-6 text-purple-600 dark:text-purple-400" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12,2C13.1,2 14,2.9 14,4C14,5.1 13.1,6 12,6C10.9,6 10,5.1 10,4C10,2.9 10.9,2 12,2M21,9V7L12,2L3,7V9C3,14.55 6.84,19.74 12,19.74C17.16,19.74 21,14.55 21,9M12,17C8.13,17 5,14.72 5,9H19C19,14.72 15.87,17 12,17Z" />
                </svg>
              </div>
              <div>
                <p class="text-2xl font-bold text-zinc-900 dark:text-zinc-100">{{ uniquePlatforms }}</p>
                <p class="text-sm text-zinc-600 dark:text-zinc-400">{{ $t('passwords.platforms') }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Search and Filter -->
        <div class="bg-white dark:bg-zinc-800 rounded-2xl p-6 shadow-lg border border-zinc-200 dark:border-zinc-700 mb-8">
          <div class="flex flex-col md:flex-row gap-4">
            <div class="flex-1">
              <div class="relative">
                <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-zinc-400" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M9.5,3A6.5,6.5 0 0,1 16,9.5C16,11.11 15.41,12.59 14.44,13.73L14.71,14H15.5L20.5,19L19,20.5L14,15.5V14.71L13.73,14.44C12.59,15.41 11.11,16 9.5,16A6.5,6.5 0 0,1 3,9.5A6.5,6.5 0 0,1 9.5,3M9.5,5C7,5 5,7 5,9.5C5,12 7,14 9.5,14C12,14 14,12 14,9.5C14,7 12,5 9.5,5Z" />
                </svg>
                <input
                  v-model="searchQuery"
                  type="text"
                  :placeholder="$t('passwords.searchPlaceholder')"
                  class="w-full pl-10 pr-4 py-3 border border-zinc-300 dark:border-zinc-600 rounded-lg bg-zinc-50 dark:bg-zinc-700 text-zinc-900 dark:text-zinc-100 placeholder-zinc-500 dark:placeholder-zinc-400 focus:ring-2 focus:ring-primary-500 focus:border-primary-500 transition-colors"
                />
              </div>
            </div>
            <select
              v-model="sortBy"
              class="px-4 py-3 border border-zinc-300 dark:border-zinc-600 rounded-lg bg-zinc-50 dark:bg-zinc-700 text-zinc-900 dark:text-zinc-100 focus:ring-2 focus:ring-primary-500 focus:border-primary-500 transition-colors"
            >
              <option value="newest">{{ $t('passwords.sortNewest') }}</option>
              <option value="oldest">{{ $t('passwords.sortOldest') }}</option>
              <option value="platform">{{ $t('passwords.sortPlatform') }}</option>
            </select>
          </div>
        </div>

        <!-- Passwords List -->
        <div v-if="loading" class="text-center py-12">
          <div class="inline-flex items-center space-x-2">
            <svg class="w-6 h-6 animate-spin text-primary-600" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12,4V2A10,10 0 0,0 2,12H4A8,8 0 0,1 12,4Z" />
            </svg>
            <span class="text-lg text-zinc-600 dark:text-zinc-300">{{ $t('passwords.loading') }}</span>
          </div>
        </div>

        <div v-else-if="filteredPasswords.length === 0" class="text-center py-12">
          <div class="mb-6">
            <div class="w-24 h-24 bg-zinc-100 dark:bg-zinc-700 rounded-full flex items-center justify-center mx-auto mb-4">
              <svg class="w-12 h-12 text-zinc-400" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12,1L21,5V9C21,16.55 15.84,23.74 9,23.74C2.16,23.74 -3,16.55 -3,9V5L12,1M12,7A2,2 0 0,0 10,9A2,2 0 0,0 12,11A2,2 0 0,0 14,9A2,2 0 0,0 12,7Z" />
              </svg>
            </div>
            <h3 class="text-xl font-semibold text-zinc-900 dark:text-zinc-100 mb-2">
              {{ searchQuery ? $t('passwords.noSearchResults') : $t('passwords.noPasswords') }}
            </h3>
            <p class="text-zinc-600 dark:text-zinc-300 mb-6">
              {{ searchQuery ? $t('passwords.tryDifferentSearch') : $t('passwords.getStartedMessage') }}
            </p>
            <router-link to="/password-generator" v-if="!searchQuery">
              <button class="bg-gradient-to-r from-primary-600 to-primary-700 hover:from-primary-700 hover:to-primary-800 text-white px-6 py-3 rounded-xl shadow-lg hover:shadow-xl transition-all duration-300">
                {{ $t('passwords.createFirst') }}
              </button>
            </router-link>
          </div>
        </div>

        <div v-else class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
          <div
            v-for="password in filteredPasswords"
            :key="password.id"
            class="group bg-white dark:bg-zinc-800 rounded-2xl p-6 shadow-lg border border-zinc-200 dark:border-zinc-700 hover:shadow-xl transition-all duration-300 hover:transform hover:scale-[1.02]"
          >
            <!-- Card Header -->
            <div class="flex items-start justify-between mb-4">
              <div class="flex items-center space-x-3">
                <div class="w-12 h-12 bg-gradient-to-br from-primary-500 to-primary-600 rounded-xl flex items-center justify-center">
                  <span class="text-white font-bold text-lg">{{ password.platform.charAt(0).toUpperCase() }}</span>
                </div>
                <div>
                  <h3 class="font-semibold text-zinc-900 dark:text-zinc-100 group-hover:text-primary-600 dark:group-hover:text-primary-400 transition-colors">
                    {{ password.platform }}
                  </h3>
                  <p class="text-sm text-zinc-500 dark:text-zinc-400">
                    {{ formatDate(password.createdAt) }}
                  </p>
                </div>
              </div>
              
              <!-- Actions Dropdown -->
              <div class="relative">
                <button
                  @click="toggleDropdown(password.id)"
                  class="p-2 rounded-lg hover:bg-zinc-100 dark:hover:bg-zinc-700 transition-colors opacity-0 group-hover:opacity-100"
                >
                  <svg class="w-5 h-5 text-zinc-600 dark:text-zinc-300" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M12,16A2,2 0 0,1 14,18A2,2 0 0,1 12,20A2,2 0 0,1 10,18A2,2 0 0,1 12,16M12,10A2,2 0 0,1 14,12A2,2 0 0,1 12,14A2,2 0 0,1 10,12A2,2 0 0,1 12,10M12,4A2,2 0 0,1 14,6A2,2 0 0,1 12,8A2,2 0 0,1 10,6A2,2 0 0,1 12,4Z" />
                  </svg>
                </button>
                
                <!-- Dropdown Menu -->
                <div
                  v-if="activeDropdown === password.id"
                  class="absolute right-0 top-10 w-48 bg-white dark:bg-zinc-700 rounded-lg shadow-xl border border-zinc-200 dark:border-zinc-600 py-2 z-10"
                >
                  <button
                    @click="copyPassword(password.password); closeDropdown()"
                    class="flex items-center space-x-3 w-full px-4 py-2 text-left text-sm text-zinc-700 dark:text-zinc-300 hover:bg-zinc-50 dark:hover:bg-zinc-600 transition-colors"
                  >
                    <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                      <path d="M19,21H8V7H19M19,5H8A2,2 0 0,0 6,7V21A2,2 0 0,0 8,23H19A2,2 0 0,0 21,21V7A2,2 0 0,0 19,5M16,1H4A2,2 0 0,0 2,3V17H4V3H16V1Z" />
                    </svg>
                    <span>{{ $t('passwords.copyPassword') }}</span>
                  </button>
                  <button
                    @click="deletePassword(password.id); closeDropdown()"
                    class="flex items-center space-x-3 w-full px-4 py-2 text-left text-sm text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 transition-colors"
                  >
                    <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                      <path d="M19,4H15.5L14.5,3H9.5L8.5,4H5V6H19M6,19A2,2 0 0,0 8,21H16A2,2 0 0,0 18,19V7H6V19Z" />
                    </svg>
                    <span>{{ $t('passwords.delete') }}</span>
                  </button>
                </div>
              </div>
            </div>

            <!-- Card Content -->
            <div class="space-y-3">
              <div v-if="password.email" class="flex items-center space-x-3">
                <svg class="w-4 h-4 text-zinc-400 flex-shrink-0" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M20,8L12,13L4,8V6L12,11L20,6M20,4H4C2.89,4 2,4.89 2,6V18A2,2 0 0,0 4,20H20A2,2 0 0,0 22,18V6C22,5.11 21.11,4 20,4Z" />
                </svg>
                <span class="text-sm text-zinc-600 dark:text-zinc-300 truncate">{{ password.email }}</span>
              </div>
              
              <div v-if="password.username" class="flex items-center space-x-3">
                <svg class="w-4 h-4 text-zinc-400 flex-shrink-0" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12,4A4,4 0 0,1 16,8A4,4 0 0,1 12,12A4,4 0 0,1 8,8A4,4 0 0,1 12,4M12,14C16.42,14 20,15.79 20,18V20H4V18C4,15.79 7.58,14 12,14Z" />
                </svg>
                <span class="text-sm text-zinc-600 dark:text-zinc-300 truncate">{{ password.username }}</span>
              </div>

              <div v-if="password.description" class="flex items-start space-x-3">
                <svg class="w-4 h-4 text-zinc-400 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z" />
                </svg>
                <span class="text-sm text-zinc-600 dark:text-zinc-300">{{ password.description }}</span>
              </div>

              <!-- Password Display -->
              <div class="mt-4 p-3 bg-zinc-50 dark:bg-zinc-700/50 rounded-lg border">
                <div class="flex items-center justify-between">
                  <span class="text-sm font-medium text-zinc-700 dark:text-zinc-300">
                    {{ $t('passwords.password') }}
                  </span>
                  <div class="flex items-center space-x-2">
                    <div :class="getPasswordStrengthColor(password.password)" class="w-2 h-2 rounded-full"></div>
                    <span class="text-xs text-zinc-500 dark:text-zinc-400">
                      {{ getPasswordStrengthText(password.password) }}
                    </span>
                  </div>
                </div>
                <div class="mt-2 font-mono text-sm text-zinc-900 dark:text-zinc-100">
                  {{ showPasswords[password.id] ? password.password : '••••••••••••' }}
                </div>
              </div>
            </div>

            <!-- Card Footer -->
            <div class="mt-4 pt-4 border-t border-zinc-200 dark:border-zinc-600">
              <div class="flex items-center justify-between">
                <button
                  @click="togglePasswordVisibility(password.id)"
                  class="flex items-center space-x-2 text-sm text-zinc-600 dark:text-zinc-300 hover:text-primary-600 dark:hover:text-primary-400 transition-colors"
                >
                  <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                    <path v-if="!showPasswords[password.id]" d="M12,9A3,3 0 0,0 9,12A3,3 0 0,0 12,15A3,3 0 0,0 15,12A3,3 0 0,0 12,9M12,17A5,5 0 0,1 7,12A5,5 0 0,1 12,7A5,5 0 0,1 17,12A5,5 0 0,1 12,17M12,4.5C7,4.5 2.73,7.61 1,12C2.73,16.39 7,19.5 12,19.5C17,19.5 21.27,16.39 23,12C21.27,7.61 17,4.5 12,4.5Z" />
                    <path v-else d="M11.83,9L15,12.16C15,12.11 15,12.05 15,12A3,3 0 0,0 12,9C11.94,9 11.89,9 11.83,9M7.53,9.8L9.08,11.35C9.03,11.56 9,11.77 9,12A3,3 0 0,0 12,15C12.22,15 12.44,14.97 12.65,14.92L14.2,16.47C13.53,16.8 12.79,17 12,17A5,5 0 0,1 7,12C7,11.21 7.2,10.47 7.53,9.8M2,4.27L4.28,6.55L4.73,7C3.08,8.3 1.78,10 1,12C2.73,16.39 7,19.5 12,19.5C13.55,19.5 15.03,19.2 16.38,18.66L16.81,19.09L19.73,22L21,20.73L3.27,3M12,7A5,5 0 0,1 17,12C17,12.64 16.87,13.26 16.64,13.82L19.57,16.75C21.07,15.5 22.27,13.86 23,12C21.27,7.61 17,4.5 12,4.5C10.6,4.5 9.26,4.75 8,5.2L10.17,7.35C10.76,7.13 11.37,7 12,7Z" />
                  </svg>
                  <span>{{ showPasswords[password.id] ? $t('passwords.hide') : $t('passwords.show') }}</span>
                </button>
                
                <button
                  @click="copyPassword(password.password)"
                  class="flex items-center space-x-2 bg-primary-600 hover:bg-primary-700 text-white px-4 py-2 rounded-lg text-sm transition-colors"
                >
                  <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M19,21H8V7H19M19,5H8A2,2 0 0,0 6,7V21A2,2 0 0,0 8,23H19A2,2 0 0,0 21,21V7A2,2 0 0,0 19,5M16,1H4A2,2 0 0,0 2,3V17H4V3H16V1Z" />
                  </svg>
                  <span>{{ $t('passwords.copy') }}</span>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { GetPasswordEntries, DeletePasswordEntry } from '../../wailsjs/go/main/App.js'

export default {
  name: 'PasswordsList',
  data() {
    return {
      passwords: [],
      loading: false,
      searchQuery: '',
      sortBy: 'newest',
      showPasswords: {},
      activeDropdown: null,
      copyMessage: ''
    }
  },
  computed: {
    filteredPasswords() {
      let filtered = this.passwords

      // Search filter
      if (this.searchQuery) {
        const query = this.searchQuery.toLowerCase()
        filtered = filtered.filter(password =>
          password.platform.toLowerCase().includes(query) ||
          (password.email && password.email.toLowerCase().includes(query)) ||
          (password.username && password.username.toLowerCase().includes(query)) ||
          (password.description && password.description.toLowerCase().includes(query))
        )
      }

      // Sort
      filtered.sort((a, b) => {
        switch (this.sortBy) {
          case 'newest':
            return new Date(b.createdAt) - new Date(a.createdAt)
          case 'oldest':
            return new Date(a.createdAt) - new Date(b.createdAt)
          case 'platform':
            return a.platform.localeCompare(b.platform)
          default:
            return 0
        }
      })

      return filtered
    },
    secureCount() {
      return this.passwords.filter(password => this.calculatePasswordStrength(password.password) >= 4).length
    },
    uniquePlatforms() {
      const platforms = new Set(this.passwords.map(password => password.platform.toLowerCase()))
      return platforms.size
    }
  },
  async mounted() {
    await this.loadPasswords()
    // Close dropdown when clicking outside
    document.addEventListener('click', this.handleClickOutside)
  },
  beforeUnmount() {
    document.removeEventListener('click', this.handleClickOutside)
  },
  methods: {
    async loadPasswords() {
      this.loading = true
      try {
        this.passwords = await GetPasswordEntries()
      } catch (error) {
        console.error('Failed to load passwords:', error)
        alert(this.$t('passwords.errorLoading'))
      } finally {
        this.loading = false
      }
    },
    async refreshPasswords() {
      await this.loadPasswords()
    },
    togglePasswordVisibility(id) {
      this.showPasswords = {
        ...this.showPasswords,
        [id]: !this.showPasswords[id]
      }
    },
    toggleDropdown(id) {
      this.activeDropdown = this.activeDropdown === id ? null : id
    },
    closeDropdown() {
      this.activeDropdown = null
    },
    handleClickOutside(event) {
      if (!event.target.closest('.relative')) {
        this.closeDropdown()
      }
    },
    async copyPassword(password) {
      try {
        await navigator.clipboard.writeText(password)
        // Show success message (you can implement a toast here)
        console.log('Password copied to clipboard')
      } catch (error) {
        console.error('Failed to copy password:', error)
        // Fallback
        const textArea = document.createElement('textarea')
        textArea.value = password
        document.body.appendChild(textArea)
        textArea.select()
        try {
          document.execCommand('copy')
          console.log('Password copied to clipboard (fallback)')
        } catch (fallbackError) {
          console.error('Fallback copy failed:', fallbackError)
        }
        document.body.removeChild(textArea)
      }
    },
    async deletePassword(id) {
      if (!confirm(this.$t('passwords.confirmDelete'))) return
      
      try {
        await DeletePasswordEntry(id)
        await this.loadPasswords() // Refresh list
        console.log('Password deleted successfully')
      } catch (error) {
        console.error('Failed to delete password:', error)
        alert(this.$t('passwords.errorDeleting'))
      }
    },
    formatDate(dateString) {
      const date = new Date(dateString)
      return date.toLocaleDateString(this.$i18n.locale, {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      })
    },
    calculatePasswordStrength(password) {
      if (!password) return 0
      
      let score = 0
      
      // Length factor
      if (password.length >= 8) score += 1
      if (password.length >= 12) score += 1
      if (password.length >= 16) score += 1
      
      // Character variety
      if (/[a-z]/.test(password)) score += 1
      if (/[A-Z]/.test(password)) score += 1
      if (/[0-9]/.test(password)) score += 1
      if (/[^a-zA-Z0-9]/.test(password)) score += 1
      
      return Math.min(score, 5)
    },
    getPasswordStrengthColor(password) {
      const strength = this.calculatePasswordStrength(password)
      if (strength <= 2) return 'bg-red-500'
      if (strength <= 3) return 'bg-yellow-500'
      if (strength <= 4) return 'bg-blue-500'
      return 'bg-green-500'
    },
    getPasswordStrengthText(password) {
      const strength = this.calculatePasswordStrength(password)
      if (strength <= 2) return this.$t('passwords.strengthWeak')
      if (strength <= 3) return this.$t('passwords.strengthFair')
      if (strength <= 4) return this.$t('passwords.strengthGood')
      return this.$t('passwords.strengthStrong')
    }
  }
}
</script>

<style scoped>
/* Custom scrollbar */
::-webkit-scrollbar {
  width: 6px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: rgba(156, 163, 175, 0.5);
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: rgba(156, 163, 175, 0.7);
}
</style>
