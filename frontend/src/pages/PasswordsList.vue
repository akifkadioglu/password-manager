<template>
  <div class="min-h-screen py-8">
    <div class="container mx-auto px-4">
      <div class="max-w-6xl mx-auto">
        <!-- Header -->
        <div class="text-center mb-8">
          <div class="flex justify-center mb-6">
            <div class="p-4 bg-gradient-to-br from-primary-600 to-primary-700 rounded-2xl shadow-2xl">
              <svg class="w-12 h-12 text-white" fill="currentColor" viewBox="0 0 24 24">
                <path :d="icons.shieldKey" />
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
                  <path :d="icons.plus" />
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
                <path :d="icons.refresh" />
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
                  <path :d="icons.shield" />
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
                  <path :d="icons.check" />
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
                  <path :d="icons.web" />
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
                  <path :d="icons.magnify" />
                </svg>
                <input
                  v-model="searchQuery"
                  type="text"
                  :placeholder="$t('passwords.searchPlaceholder')"
                  class="w-full pl-10 pr-4 py-3 border border-zinc-300 dark:border-zinc-600 rounded-lg bg-zinc-50 dark:bg-zinc-700 text-zinc-900 dark:text-zinc-100 placeholder-zinc-500 dark:placeholder-zinc-400 focus:ring-2 focus:ring-primary-500 focus:border-primary-500 transition-colors"
                />
              </div>
            </div>
            <!-- Custom Sort Dropdown -->
            <div class="relative inline-block text-left">
              <button
                @click="sortDropdownOpen = !sortDropdownOpen"
                class="inline-flex items-center justify-between w-full px-4 py-3 border border-zinc-300 dark:border-zinc-600 rounded-lg bg-zinc-50 dark:bg-zinc-700 text-zinc-900 dark:text-zinc-100 focus:ring-2 focus:ring-primary-500 focus:border-primary-500 transition-colors min-w-[160px]"
              >
                <span>{{ getSortOptionText(sortBy) }}</span>
                <svg 
                  class="ml-2 h-5 w-5 transform transition-transform duration-200" 
                  :class="{ 'rotate-180': sortDropdownOpen }"
                  fill="currentColor" 
                  viewBox="0 0 20 20"
                >
                  <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
                </svg>
              </button>

              <!-- Backdrop -->
              <div 
                v-if="sortDropdownOpen"
                @click="sortDropdownOpen = false"
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
                  v-show="sortDropdownOpen"
                  class="origin-top-right absolute right-0 mt-2 w-full min-w-[160px] rounded-xl shadow-2xl bg-white dark:bg-zinc-800 ring-1 ring-black ring-opacity-5 z-50"
                >
                  <div class="py-2">
                    <button
                      v-for="option in sortOptions"
                      :key="option.value"
                      @click="setSortBy(option.value)"
                      class="w-full text-left px-4 py-3 text-sm text-zinc-700 dark:text-zinc-300 hover:bg-primary-50 dark:hover:bg-primary-900/20 flex items-center transition-all duration-150"
                      :class="{ 
                        'bg-primary-50 dark:bg-primary-900/20 text-primary-600 dark:text-primary-400 font-medium': sortBy === option.value
                      }"
                    >
                      <span class="flex-1">{{ option.text }}</span>
                      <span v-if="sortBy === option.value" class="ml-2 text-primary-500">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                          <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                        </svg>
                      </span>
                    </button>
                  </div>
                </div>
              </transition>
            </div>
          </div>
        </div>

        <!-- Passwords List -->
        <div v-if="loading" class="text-center py-12">
          <div class="inline-flex items-center space-x-2">
            <svg class="w-6 h-6 animate-spin text-primary-600" fill="currentColor" viewBox="0 0 24 24">
              <path :d="icons.loading" />
            </svg>
            <span class="text-lg text-zinc-600 dark:text-zinc-300">{{ $t('passwords.loading') }}</span>
          </div>
        </div>

        <div v-else-if="filteredPasswords.length === 0" class="text-center py-12">
          <div class="mb-6">
            <div class="w-24 h-24 bg-zinc-100 dark:bg-zinc-700 rounded-full flex items-center justify-center mx-auto mb-4">
              <svg class="w-12 h-12 text-zinc-400" fill="currentColor" viewBox="0 0 24 24">
                <path :d="icons.shieldKey" />
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
                    <path :d="icons.dotsVertical" />
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
                      <path :d="icons.contentCopy" />
                    </svg>
                    <span>{{ $t('passwords.copyPassword') }}</span>
                  </button>
                  <button
                    @click="deletePassword(password.id); closeDropdown()"
                    class="flex items-center space-x-3 w-full px-4 py-2 text-left text-sm text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 transition-colors"
                  >
                    <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                      <path :d="icons.delete" />
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
                  <path :d="icons.email" />
                </svg>
                <span class="text-sm text-zinc-600 dark:text-zinc-300 truncate">{{ password.email }}</span>
              </div>
              
              <div v-if="password.username" class="flex items-center space-x-3">
                <svg class="w-4 h-4 text-zinc-400 flex-shrink-0" fill="currentColor" viewBox="0 0 24 24">
                  <path :d="icons.account" />
                </svg>
                <span class="text-sm text-zinc-600 dark:text-zinc-300 truncate">{{ password.username }}</span>
              </div>

              <div v-if="password.description" class="flex items-start space-x-3">
                <svg class="w-4 h-4 text-zinc-400 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 24 24">
                  <path :d="icons.noteText" />
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
                    <path :d="showPasswords[password.id] ? icons.eyeOff : icons.eye" />
                  </svg>
                  <span>{{ showPasswords[password.id] ? $t('passwords.hide') : $t('passwords.show') }}</span>
                </button>
                
                <button
                  @click="copyPassword(password.password)"
                  class="flex items-center space-x-2 bg-primary-600 hover:bg-primary-700 text-white px-4 py-2 rounded-lg text-sm transition-colors"
                >
                  <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                    <path :d="icons.contentCopy" />
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
import {
  mdiShieldKey,
  mdiPlus,
  mdiRefresh,
  mdiShield,
  mdiCheck,
  mdiWeb,
  mdiMagnify,
  mdiDotsVertical,
  mdiContentCopy,
  mdiDelete,
  mdiEmail,
  mdiAccount,
  mdiNoteText,
  mdiEye,
  mdiEyeOff,
  mdiLoading
} from '@mdi/js'

export default {
  name: 'PasswordsList',
  data() {
    return {
      passwords: [],
      loading: false,
      searchQuery: '',
      sortBy: 'newest',
      sortDropdownOpen: false,
      showPasswords: {},
      activeDropdown: null,
      copyMessage: '',
      // MDI Icons
      icons: {
        shieldKey: mdiShieldKey,
        plus: mdiPlus,
        refresh: mdiRefresh,
        shield: mdiShield,
        check: mdiCheck,
        web: mdiWeb,
        magnify: mdiMagnify,
        dotsVertical: mdiDotsVertical,
        contentCopy: mdiContentCopy,
        delete: mdiDelete,
        email: mdiEmail,
        account: mdiAccount,
        noteText: mdiNoteText,
        eye: mdiEye,
        eyeOff: mdiEyeOff,
        loading: mdiLoading
      }
    }
  },
  computed: {
    sortOptions() {
      return [
        { value: 'newest', text: this.$t('passwords.sortNewest') },
        { value: 'oldest', text: this.$t('passwords.sortOldest') },
        { value: 'platform', text: this.$t('passwords.sortPlatform') }
      ]
    },
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
        this.sortDropdownOpen = false
      }
    },
    getSortOptionText(value) {
      const option = this.sortOptions.find(opt => opt.value === value)
      return option ? option.text : this.$t('passwords.sortNewest')
    },
    setSortBy(value) {
      this.sortBy = value
      this.sortDropdownOpen = false
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
