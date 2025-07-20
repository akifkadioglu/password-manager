<template>
  <div class="min-h-screen py-8">
    <div class="container mx-auto px-4">
      <div class="max-w-4xl mx-auto">
        <!-- Header -->
        <div class="text-center mb-8">
          <h1 class="text-3xl font-bold text-zinc-900 dark:text-zinc-100 mb-4">
            {{ $t('passwordGenerator.title') }}
          </h1>
          <p class="text-lg text-zinc-700 dark:text-zinc-300">
            {{ $t('passwordGenerator.subtitle') }}
          </p>
        </div>

          <div class="bg-white dark:bg-zinc-800 rounded-lg shadow-app-md border border-zinc-200 dark:border-zinc-700 p-6 mb-6">
          <div class="flex items-center space-x-3 mb-4">
            <div class="p-2 bg-blue-100 dark:bg-blue-900/30 rounded-lg">
              <svg class="w-5 h-5 text-blue-600 dark:text-blue-400" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12,1L21,5V9C21,16.55 15.84,23.74 9,23.74C2.16,23.74 -3,16.55 -3,9V5L12,1Z" />
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-zinc-900 dark:text-zinc-100">
              {{ $t('passwordGenerator.systemFingerprint') }}
            </h3>
          </div>
          <div class="flex items-center space-x-3">
            <code class="bg-zinc-100 dark:bg-zinc-700 px-4 py-3 rounded-lg text-sm font-mono text-zinc-800 dark:text-zinc-200 flex-1 border">
              {{ systemFingerprint || 'Yükleniyor...' }}
            </code>
            <button
              @click="refreshFingerprint"
              class="flex items-center justify-center w-12 h-12 bg-zinc-100 dark:bg-zinc-700 hover:bg-zinc-200 dark:hover:bg-zinc-600 rounded-lg transition-all duration-200 border"
              title="Yenile"
            >
              <svg class="w-5 h-5 text-zinc-600 dark:text-zinc-300" fill="currentColor" viewBox="0 0 24 24">
                <path d="M17.65,6.35C16.2,4.9 14.21,4 12,4A8,8 0 0,0 4,12A8,8 0 0,0 12,20C15.73,20 18.84,17.45 19.73,14H17.65C16.83,16.33 14.61,18 12,18A6,6 0 0,1 6,12A6,6 0 0,1 12,6C13.66,6 15.14,6.69 16.22,7.78L13,11H20V4L17.65,6.35Z" />
              </svg>
            </button>
          </div>
        </div>

        <!-- Password Generator -->
        <div class="bg-white dark:bg-zinc-800 rounded-lg shadow-app-md border border-zinc-200 dark:border-zinc-700 p-6">
          <!-- Generated Password Display -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2">
              {{ $t('passwordGenerator.generatedPassword') }}
            </label>
            <div class="flex items-center space-x-3">
              <input
                v-model="generatedPassword"
                :type="showPassword ? 'text' : 'password'"
                readonly
                class="flex-1 px-4 py-3 border border-zinc-300 dark:border-zinc-600 rounded-lg bg-zinc-50 dark:bg-zinc-700 text-zinc-900 dark:text-zinc-100 font-mono text-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                :placeholder="$t('passwordGenerator.noPasswordGenerated')"
              />
              <button
                @click="togglePasswordVisibility"
                class="flex items-center justify-center w-12 h-12 bg-zinc-100 dark:bg-zinc-700 hover:bg-zinc-200 dark:hover:bg-zinc-600 rounded-lg transition-all duration-200 border border-zinc-300 dark:border-zinc-600"
                :title="showPassword ? 'Gizle' : 'Göster'"
              >
                <svg class="w-5 h-5 text-zinc-600 dark:text-zinc-300" fill="currentColor" viewBox="0 0 24 24">
                  <path v-if="!showPassword" d="M12,9A3,3 0 0,0 9,12A3,3 0 0,0 12,15A3,3 0 0,0 15,12A3,3 0 0,0 12,9M12,17A5,5 0 0,1 7,12A5,5 0 0,1 12,7A5,5 0 0,1 17,12A5,5 0 0,1 12,17M12,4.5C7,4.5 2.73,7.61 1,12C2.73,16.39 7,19.5 12,19.5C17,19.5 21.27,16.39 23,12C21.27,7.61 17,4.5 12,4.5Z" />
                  <path v-else d="M11.83,9L15,12.16C15,12.11 15,12.05 15,12A3,3 0 0,0 12,9C11.94,9 11.89,9 11.83,9M7.53,9.8L9.08,11.35C9.03,11.56 9,11.77 9,12A3,3 0 0,0 12,15C12.22,15 12.44,14.97 12.65,14.92L14.2,16.47C13.53,16.8 12.79,17 12,17A5,5 0 0,1 7,12C7,11.21 7.2,10.47 7.53,9.8M2,4.27L4.28,6.55L4.73,7C3.08,8.3 1.78,10 1,12C2.73,16.39 7,19.5 12,19.5C13.55,19.5 15.03,19.2 16.38,18.66L16.81,19.09L19.73,22L21,20.73L3.27,3M12,7A5,5 0 0,1 17,12C17,12.64 16.87,13.26 16.64,13.82L19.57,16.75C21.07,15.5 22.27,13.86 23,12C21.27,7.61 17,4.5 12,4.5C10.6,4.5 9.26,4.75 8,5.2L10.17,7.35C10.76,7.13 11.37,7 12,7Z" />
                </svg>
              </button>
              <button
                @click="copyPassword"
                :disabled="!generatedPassword"
                class="flex items-center justify-center w-12 h-12 bg-primary-600 hover:bg-primary-700 disabled:bg-zinc-300 dark:disabled:bg-zinc-600 rounded-lg transition-all duration-200 disabled:cursor-not-allowed"
                title="Kopyala"
              >
                <svg class="w-5 h-5 text-white dark:text-zinc-100" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M19,21H8V7H19M19,5H8A2,2 0 0,0 6,7V21A2,2 0 0,0 8,23H19A2,2 0 0,0 21,21V7A2,2 0 0,0 19,5M16,1H4A2,2 0 0,0 2,3V17H4V3H16V1Z" />
                </svg>
              </button>
              <button
                @click="openSaveModal"
                :disabled="!generatedPassword"
                class="flex items-center justify-center w-12 h-12 bg-green-600 hover:bg-green-700 disabled:bg-zinc-300 dark:disabled:bg-zinc-600 rounded-lg transition-all duration-200 disabled:cursor-not-allowed"
                title="Kaydet"
              >
                <svg class="w-5 h-5 text-white dark:text-zinc-100" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M15,9H5V5H15M12,19A3,3 0 0,1 9,16A3,3 0 0,1 12,13A3,3 0 0,1 15,16A3,3 0 0,1 12,19M17,3H5C3.89,3 3,3.9 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V7L17,3Z" />
                </svg>
              </button>
            </div>
            <div v-if="copyMessage" class="text-sm text-green-600 dark:text-green-400 mt-2">
              {{ copyMessage }}
            </div>
          </div>

          <!-- Password Options -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- Length -->
            <div>
              <label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2">
                {{ $t('passwordGenerator.length') }}: {{ passwordOptions.length }}
              </label>
              <input
                v-model.number="passwordOptions.length"
                type="range"
                min="4"
                max="128"
                class="w-full h-2 bg-zinc-200 dark:bg-zinc-600 rounded-lg appearance-none cursor-pointer"
              />
            </div>

            <!-- Character Options -->
            <div class="space-y-3">
              <label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2">
                {{ $t('passwordGenerator.characterTypes') }}
              </label>
              
              <div class="space-y-2">
                <label class="flex items-center">
                  <input
                    v-model="passwordOptions.includeUppercase"
                    type="checkbox"
                    class="mr-3 h-4 w-4 text-primary-600 rounded border-zinc-300 dark:border-zinc-600"
                  />
                  <span class="text-sm text-zinc-700 dark:text-zinc-300">
                    {{ $t('passwordGenerator.uppercase') }} (A-Z)
                  </span>
                </label>

                <label class="flex items-center">
                  <input
                    v-model="passwordOptions.includeLowercase"
                    type="checkbox"
                    class="mr-3 h-4 w-4 text-primary-600 rounded border-zinc-300 dark:border-zinc-600"
                  />
                  <span class="text-sm text-zinc-700 dark:text-zinc-300">
                    {{ $t('passwordGenerator.lowercase') }} (a-z)
                  </span>
                </label>

                <label class="flex items-center">
                  <input
                    v-model="passwordOptions.includeNumbers"
                    type="checkbox"
                    class="mr-3 h-4 w-4 text-primary-600 rounded border-zinc-300 dark:border-zinc-600"
                  />
                  <span class="text-sm text-zinc-700 dark:text-zinc-300">
                    {{ $t('passwordGenerator.numbers') }} (0-9)
                  </span>
                </label>

                <label class="flex items-center">
                  <input
                    v-model="passwordOptions.includeSymbols"
                    type="checkbox"
                    class="mr-3 h-4 w-4 text-primary-600 rounded border-zinc-300 dark:border-zinc-600"
                  />
                  <span class="text-sm text-zinc-700 dark:text-zinc-300">
                    {{ $t('passwordGenerator.symbols') }} (!@#$%^&*)
                  </span>
                </label>

                <label class="flex items-center">
                  <input
                    v-model="passwordOptions.useFingerprint"
                    type="checkbox"
                    class="mr-3 h-4 w-4 text-primary-600 rounded border-zinc-300 dark:border-zinc-600"
                  />
                  <span class="text-sm text-zinc-700 dark:text-zinc-300">
                    {{ $t('passwordGenerator.useFingerprint') }}
                  </span>
                </label>
              </div>
            </div>
          </div>

          <!-- Generate Button -->
          <div class="mt-8 flex justify-center">
            <button
              @click="generatePassword"
              :disabled="isGenerating || !hasValidOptions"
              class="flex items-center space-x-3 bg-gradient-to-r from-primary-600 to-primary-700 hover:from-primary-700 hover:to-primary-800 disabled:from-zinc-400 disabled:to-zinc-500 text-white px-8 py-4 rounded-xl shadow-lg hover:shadow-xl transition-all duration-300 disabled:cursor-not-allowed text-lg font-semibold min-w-[200px] justify-center"
            >
              <svg 
                class="w-6 h-6" 
                :class="{ 'animate-spin': isGenerating }"
                fill="currentColor" 
                viewBox="0 0 24 24"
              >
                <path v-if="!isGenerating" d="M12,1L21,5V9C21,16.55 15.84,23.74 9,23.74C2.16,23.74 -3,16.55 -3,9V5L12,1M12,7A2,2 0 0,0 10,9A2,2 0 0,0 12,11A2,2 0 0,0 14,9A2,2 0 0,0 12,7Z" />
                <path v-else d="M12,4V2A10,10 0 0,0 2,12H4A8,8 0 0,1 12,4Z" />
              </svg>
              <span>
                {{ isGenerating ? $t('passwordGenerator.generating') : $t('passwordGenerator.generate') }}
              </span>
            </button>
          </div>

          <!-- Password Strength -->
          <div v-if="generatedPassword" class="mt-6">
            <label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2">
              {{ $t('passwordGenerator.strength') }}
            </label>
            <div class="w-full bg-zinc-200 dark:bg-zinc-600 rounded-full h-2">
              <div
                :class="strengthColor"
                class="h-2 rounded-full transition-all duration-300"
                :style="{ width: strengthPercentage + '%' }"
              ></div>
            </div>
            <div class="text-sm mt-1" :class="strengthTextColor">
              {{ strengthText }}
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Password Save Modal -->
    <PasswordSaveModal 
      :show="showSaveModal"
      :password="generatedPassword"
      @close="closeSaveModal"
      @saved="onPasswordSaved"
    />
  </div>
</template>

<script>
import { GeneratePassword, GetSystemFingerprint } from '../../wailsjs/go/main/App.js'
import PasswordSaveModal from '../components/PasswordSaveModal.vue'

export default {
  name: 'PasswordGenerator',
  components: {
    PasswordSaveModal
  },
  data() {
    return {
      generatedPassword: '',
      showPassword: false,
      copyMessage: '',
      isGenerating: false,
      systemFingerprint: '',
      showSaveModal: false,
      passwordOptions: {
        length: 16,
        includeUppercase: true,
        includeLowercase: true,
        includeNumbers: true,
        includeSymbols: true,
        useFingerprint: true
      }
    }
  },
  computed: {
    hasValidOptions() {
      return (
        this.passwordOptions.includeUppercase ||
        this.passwordOptions.includeLowercase ||
        this.passwordOptions.includeNumbers ||
        this.passwordOptions.includeSymbols
      )
    },
    passwordStrength() {
      if (!this.generatedPassword) return 0
      
      let score = 0
      const password = this.generatedPassword
      
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
    strengthPercentage() {
      return (this.passwordStrength / 5) * 100
    },
    strengthColor() {
      const strength = this.passwordStrength
      if (strength <= 2) return 'bg-red-500'
      if (strength <= 3) return 'bg-yellow-500'
      if (strength <= 4) return 'bg-blue-500'
      return 'bg-green-500'
    },
    strengthTextColor() {
      const strength = this.passwordStrength
      if (strength <= 2) return 'text-red-600 dark:text-red-400'
      if (strength <= 3) return 'text-yellow-600 dark:text-yellow-400'
      if (strength <= 4) return 'text-blue-600 dark:text-blue-400'
      return 'text-green-600 dark:text-green-400'
    },
    strengthText() {
      const strength = this.passwordStrength
      if (strength <= 2) return this.$t('passwordGenerator.strengthWeak')
      if (strength <= 3) return this.$t('passwordGenerator.strengthFair')
      if (strength <= 4) return this.$t('passwordGenerator.strengthGood')
      return this.$t('passwordGenerator.strengthStrong')
    }
  },
  async mounted() {
    await this.loadSystemFingerprint()
    await this.generatePassword()
  },
  methods: {
    async loadSystemFingerprint() {
      try {
        this.systemFingerprint = await GetSystemFingerprint()
      } catch (error) {
        console.error('Failed to get system fingerprint:', error)
        this.systemFingerprint = 'Error loading fingerprint'
      }
    },
    async refreshFingerprint() {
      await this.loadSystemFingerprint()
    },
    async generatePassword() {
      if (!this.hasValidOptions) return
      
      this.isGenerating = true
      try {
        this.generatedPassword = await GeneratePassword(this.passwordOptions)
        this.copyMessage = ''
      } catch (error) {
        console.error('Failed to generate password:', error)
        alert(this.$t('passwordGenerator.errorGenerating'))
      } finally {
        this.isGenerating = false
      }
    },
    togglePasswordVisibility() {
      this.showPassword = !this.showPassword
    },
    async copyPassword() {
      if (!this.generatedPassword) return
      
      try {
        await navigator.clipboard.writeText(this.generatedPassword)
        this.copyMessage = this.$t('passwordGenerator.copied')
        setTimeout(() => {
          this.copyMessage = ''
        }, 2000)
      } catch (error) {
        console.error('Failed to copy password:', error)
        // Fallback for older browsers
        const textArea = document.createElement('textarea')
        textArea.value = this.generatedPassword
        document.body.appendChild(textArea)
        textArea.select()
        try {
          document.execCommand('copy')
          this.copyMessage = this.$t('passwordGenerator.copied')
          setTimeout(() => {
            this.copyMessage = ''
          }, 2000)
        } catch (fallbackError) {
          console.error('Fallback copy failed:', fallbackError)
        }
        document.body.removeChild(textArea)
      }
    },
    openSaveModal() {
      if (!this.generatedPassword) return
      this.showSaveModal = true
    },
    closeSaveModal() {
      this.showSaveModal = false
    },
    onPasswordSaved(savedEntry) {
      console.log('Password saved:', savedEntry)
      // Redirect to passwords list
      this.$router.push('/passwords')
    }
  }
}
</script>

<style scoped>
/* Custom slider styles */
input[type="range"]::-webkit-slider-thumb {
  appearance: none;
  height: 20px;
  width: 20px;
  border-radius: 50%;
  background: #3B82F6;
  cursor: pointer;
}

input[type="range"]::-moz-range-thumb {
  height: 20px;
  width: 20px;
  border-radius: 50%;
  background: #3B82F6;
  cursor: pointer;
  border: none;
}
</style>
