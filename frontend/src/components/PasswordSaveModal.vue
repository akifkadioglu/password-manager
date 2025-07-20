<template>
  <transition 
    name="modal" 
    enter-active-class="transition-all duration-300 ease-out" 
    enter-from-class="opacity-0" 
    enter-to-class="opacity-100"
    leave-active-class="transition-all duration-200 ease-in" 
    leave-from-class="opacity-100" 
    leave-to-class="opacity-0"
  >
    <div v-if="show" class="fixed inset-0 z-50 overflow-y-auto" @click="handleBackdropClick">
      <!-- Backdrop -->
      <div class="fixed inset-0 bg-black/60 backdrop-blur-sm"></div>
      
      <!-- Modal Container -->
      <div class="flex min-h-full items-center justify-center p-4">
        <transition
          name="modal-content"
          enter-active-class="transition-all duration-300 ease-out"
          enter-from-class="opacity-0 scale-95 translate-y-4"
          enter-to-class="opacity-100 scale-100 translate-y-0"
          leave-active-class="transition-all duration-200 ease-in"
          leave-from-class="opacity-100 scale-100 translate-y-0"
          leave-to-class="opacity-0 scale-95 translate-y-4"
        >
          <div 
            v-if="show"
            class="relative bg-white dark:bg-zinc-800 rounded-2xl shadow-2xl border border-zinc-200 dark:border-zinc-700 w-full max-w-2xl max-h-[90vh] overflow-hidden"
            @click.stop
          >
            <!-- Header -->
            <div class="bg-gradient-to-r from-primary-600 to-primary-700 px-6 py-4">
              <div class="flex items-center justify-between">
                <div class="flex items-center space-x-3">
                  <div class="p-2 bg-white/20 rounded-lg backdrop-blur-sm">
                    <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 24 24">
                      <path d="M12,1L21,5V9C21,16.55 15.84,23.74 9,23.74C2.16,23.74 -3,16.55 -3,9V5L12,1M12,7A2,2 0 0,0 10,9A2,2 0 0,0 12,11A2,2 0 0,0 14,9A2,2 0 0,0 12,7Z" />
                    </svg>
                  </div>
                  <div>
                    <h2 class="text-xl font-bold text-white">
                      {{ $t('passwordModal.title') }}
                    </h2>
                    <p class="text-primary-100 text-sm">
                      {{ $t('passwordModal.subtitle') }}
                    </p>
                  </div>
                </div>
                <button
                  @click="close"
                  class="p-2 rounded-lg hover:bg-white/20 transition-colors duration-200"
                >
                  <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M19,6.41L17.59,5L12,10.59L6.41,5L5,6.41L10.59,12L5,17.59L6.41,19L12,13.41L17.59,19L19,17.59L13.41,12L19,6.41Z" />
                  </svg>
                </button>
              </div>
            </div>

            <!-- Content -->
            <div class="p-6 space-y-6 max-h-[70vh] overflow-y-auto">
              <!-- Generated Password Display -->
              <div class="bg-zinc-50 dark:bg-zinc-700/50 rounded-xl p-4 border border-zinc-200 dark:border-zinc-600">
                <label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-3">
                  {{ $t('passwordModal.generatedPassword') }}
                </label>
                <div class="flex items-center space-x-3">
                  <input
                    :value="password"
                    :type="showPassword ? 'text' : 'password'"
                    readonly
                    class="flex-1 px-4 py-3 bg-white dark:bg-zinc-700 border border-zinc-300 dark:border-zinc-600 rounded-lg text-zinc-900 dark:text-zinc-100 font-mono text-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                  />
                  <button
                    @click="togglePasswordVisibility"
                    class="flex items-center justify-center w-12 h-12 bg-zinc-100 dark:bg-zinc-600 hover:bg-zinc-200 dark:hover:bg-zinc-500 rounded-lg transition-all duration-200"
                    :title="showPassword ? 'Gizle' : 'Göster'"
                  >
                    <svg class="w-5 h-5 text-zinc-600 dark:text-zinc-300" fill="currentColor" viewBox="0 0 24 24">
                      <path v-if="!showPassword" d="M12,9A3,3 0 0,0 9,12A3,3 0 0,0 12,15A3,3 0 0,0 15,12A3,3 0 0,0 12,9M12,17A5,5 0 0,1 7,12A5,5 0 0,1 12,7A5,5 0 0,1 17,12A5,5 0 0,1 12,17M12,4.5C7,4.5 2.73,7.61 1,12C2.73,16.39 7,19.5 12,19.5C17,19.5 21.27,16.39 23,12C21.27,7.61 17,4.5 12,4.5Z" />
                      <path v-else d="M11.83,9L15,12.16C15,12.11 15,12.05 15,12A3,3 0 0,0 12,9C11.94,9 11.89,9 11.83,9M7.53,9.8L9.08,11.35C9.03,11.56 9,11.77 9,12A3,3 0 0,0 12,15C12.22,15 12.44,14.97 12.65,14.92L14.2,16.47C13.53,16.8 12.79,17 12,17A5,5 0 0,1 7,12C7,11.21 7.2,10.47 7.53,9.8M2,4.27L4.28,6.55L4.73,7C3.08,8.3 1.78,10 1,12C2.73,16.39 7,19.5 12,19.5C13.55,19.5 15.03,19.2 16.38,18.66L16.81,19.09L19.73,22L21,20.73L3.27,3M12,7A5,5 0 0,1 17,12C17,12.64 16.87,13.26 16.64,13.82L19.57,16.75C21.07,15.5 22.27,13.86 23,12C21.27,7.61 17,4.5 12,4.5C10.6,4.5 9.26,4.75 8,5.2L10.17,7.35C10.76,7.13 11.37,7 12,7Z" />
                    </svg>
                  </button>
                  <button
                    @click="copyPassword"
                    class="flex items-center justify-center w-12 h-12 bg-primary-600 hover:bg-primary-700 rounded-lg transition-all duration-200"
                    title="Kopyala"
                  >
                    <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 24 24">
                      <path d="M19,21H8V7H19M19,5H8A2,2 0 0,0 6,7V21A2,2 0 0,0 8,23H19A2,2 0 0,0 21,21V7A2,2 0 0,0 19,5M16,1H4A2,2 0 0,0 2,3V17H4V3H16V1Z" />
                    </svg>
                  </button>
                </div>
                <div v-if="copyMessage" class="text-sm text-green-600 dark:text-green-400 mt-2">
                  {{ copyMessage }}
                </div>
              </div>

              <!-- Form Fields -->
              <form @submit.prevent="saveEntry" class="space-y-5">
                <!-- Platform -->
                <div>
                  <label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2">
                    <div class="flex items-center space-x-2">
                      <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M11,16.5L6.5,12L7.91,10.59L11,13.67L16.59,8.09L18,9.5L11,16.5Z" />
                      </svg>
                      <span>{{ $t('passwordModal.platform') }}</span>
                      <span class="text-red-500">*</span>
                    </div>
                  </label>
                  <input
                    v-model="form.platform"
                    type="text"
                    :placeholder="$t('passwordModal.platformPlaceholder')"
                    required
                    class="w-full px-4 py-3 border border-zinc-300 dark:border-zinc-600 rounded-lg bg-white dark:bg-zinc-700 text-zinc-900 dark:text-zinc-100 placeholder-zinc-500 dark:placeholder-zinc-400 focus:ring-2 focus:ring-primary-500 focus:border-primary-500 transition-colors"
                  />
                </div>

                <!-- Email & Username -->
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div>
                    <label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2">
                      <div class="flex items-center space-x-2">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                          <path d="M20,8L12,13L4,8V6L12,11L20,6M20,4H4C2.89,4 2,4.89 2,6V18A2,2 0 0,0 4,20H20A2,2 0 0,0 22,18V6C22,5.11 21.11,4 20,4Z" />
                        </svg>
                        <span>{{ $t('passwordModal.email') }}</span>
                      </div>
                    </label>
                    <input
                      v-model="form.email"
                      type="email"
                      :placeholder="$t('passwordModal.emailPlaceholder')"
                      class="w-full px-4 py-3 border border-zinc-300 dark:border-zinc-600 rounded-lg bg-white dark:bg-zinc-700 text-zinc-900 dark:text-zinc-100 placeholder-zinc-500 dark:placeholder-zinc-400 focus:ring-2 focus:ring-primary-500 focus:border-primary-500 transition-colors"
                    />
                  </div>

                  <div>
                    <label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2">
                      <div class="flex items-center space-x-2">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                          <path d="M12,4A4,4 0 0,1 16,8A4,4 0 0,1 12,12A4,4 0 0,1 8,8A4,4 0 0,1 12,4M12,14C16.42,14 20,15.79 20,18V20H4V18C4,15.79 7.58,14 12,14Z" />
                        </svg>
                        <span>{{ $t('passwordModal.username') }}</span>
                      </div>
                    </label>
                    <input
                      v-model="form.username"
                      type="text"
                      :placeholder="$t('passwordModal.usernamePlaceholder')"
                      class="w-full px-4 py-3 border border-zinc-300 dark:border-zinc-600 rounded-lg bg-white dark:bg-zinc-700 text-zinc-900 dark:text-zinc-100 placeholder-zinc-500 dark:placeholder-zinc-400 focus:ring-2 focus:ring-primary-500 focus:border-primary-500 transition-colors"
                    />
                  </div>
                </div>

                <!-- Description -->
                <div>
                  <label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2">
                    <div class="flex items-center space-x-2">
                      <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z" />
                      </svg>
                      <span>{{ $t('passwordModal.description') }}</span>
                    </div>
                  </label>
                  <textarea
                    v-model="form.description"
                    :placeholder="$t('passwordModal.descriptionPlaceholder')"
                    rows="3"
                    class="w-full px-4 py-3 border border-zinc-300 dark:border-zinc-600 rounded-lg bg-white dark:bg-zinc-700 text-zinc-900 dark:text-zinc-100 placeholder-zinc-500 dark:placeholder-zinc-400 focus:ring-2 focus:ring-primary-500 focus:border-primary-500 transition-colors resize-none"
                  ></textarea>
                </div>
              </form>
            </div>

            <!-- Footer -->
            <div class="bg-zinc-50 dark:bg-zinc-700/50 px-6 py-4 border-t border-zinc-200 dark:border-zinc-600">
              <div class="flex items-center justify-between">
                <div class="flex items-center space-x-2 text-sm text-zinc-500 dark:text-zinc-400">
                  <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M12,1L21,5V9C21,16.55 15.84,23.74 9,23.74C2.16,23.74 -3,16.55 -3,9V5L12,1M12,7A2,2 0 0,0 10,9A2,2 0 0,0 12,11A2,2 0 0,0 14,9A2,2 0 0,0 12,7Z" />
                  </svg>
                  <span>{{ $t('passwordModal.secureStorage') }}</span>
                </div>
                
                <div class="flex items-center space-x-3">
                  <button
                    @click="close"
                    type="button"
                    class="px-6 py-2.5 text-zinc-700 dark:text-zinc-300 bg-white dark:bg-zinc-700 border border-zinc-300 dark:border-zinc-600 rounded-lg hover:bg-zinc-50 dark:hover:bg-zinc-600 transition-colors duration-200 font-medium"
                  >
                    {{ $t('common.cancel') }}
                  </button>
                  <button
                    @click="saveEntry"
                    :disabled="!form.platform || saving"
                    class="px-6 py-2.5 bg-gradient-to-r from-primary-600 to-primary-700 hover:from-primary-700 hover:to-primary-800 text-white rounded-lg font-medium shadow-lg hover:shadow-xl transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed flex items-center space-x-2"
                  >
                    <svg v-if="saving" class="w-4 h-4 animate-spin" fill="currentColor" viewBox="0 0 24 24">
                      <path d="M12,4V2A10,10 0 0,0 2,12H4A8,8 0 0,1 12,4Z" />
                    </svg>
                    <svg v-else class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                      <path d="M15,9H5V5H15M12,19A3,3 0 0,1 9,16A3,3 0 0,1 12,13A3,3 0 0,1 15,16A3,3 0 0,1 12,19M17,3H5C3.89,3 3,3.9 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V7L17,3Z" />
                    </svg>
                    <span>{{ saving ? $t('passwordModal.saving') : $t('passwordModal.save') }}</span>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </transition>
      </div>
    </div>
  </transition>
</template>

<script>
import { SavePasswordEntry } from '../../wailsjs/go/main/App.js'

export default {
  name: 'PasswordSaveModal',
  props: {
    show: {
      type: Boolean,
      default: false
    },
    password: {
      type: String,
      required: true
    }
  },
  emits: ['close', 'saved'],
  data() {
    return {
      showPassword: false,
      copyMessage: '',
      saving: false,
      form: {
        platform: '',
        email: '',
        username: '',
        description: ''
      }
    }
  },
  watch: {
    show(newVal) {
      if (newVal) {
        // Reset form when modal opens
        this.form = {
          platform: '',
          email: '',
          username: '',
          description: ''
        }
        this.showPassword = false
        this.copyMessage = ''
        this.saving = false
      }
    }
  },
  methods: {
    close() {
      this.$emit('close')
    },
    handleBackdropClick(event) {
      if (event.target === event.currentTarget) {
        this.close()
      }
    },
    togglePasswordVisibility() {
      this.showPassword = !this.showPassword
    },
    async copyPassword() {
      try {
        await navigator.clipboard.writeText(this.password)
        this.copyMessage = this.$t('passwordModal.copied')
        setTimeout(() => {
          this.copyMessage = ''
        }, 2000)
      } catch (error) {
        console.error('Failed to copy password:', error)
        // Fallback
        const textArea = document.createElement('textarea')
        textArea.value = this.password
        document.body.appendChild(textArea)
        textArea.select()
        try {
          document.execCommand('copy')
          this.copyMessage = this.$t('passwordModal.copied')
          setTimeout(() => {
            this.copyMessage = ''
          }, 2000)
        } catch (fallbackError) {
          console.error('Fallback copy failed:', fallbackError)
        }
        document.body.removeChild(textArea)
      }
    },
    async saveEntry() {
      if (!this.form.platform || this.saving) return
      
      this.saving = true
      try {
        const entry = {
          platform: this.form.platform,
          email: this.form.email,
          username: this.form.username,
          description: this.form.description,
          password: this.password
        }
        
        const savedEntry = await SavePasswordEntry(entry)
        this.$emit('saved', savedEntry)
        this.close()
        
        // Show success message (you can implement a toast notification here)
        console.log('Password entry saved successfully:', savedEntry)
      } catch (error) {
        console.error('Failed to save password entry:', error)
        alert(this.$t('passwordModal.errorSaving'))
      } finally {
        this.saving = false
      }
    }
  }
}
</script>

<style scoped>
/* Custom scrollbar for the content area */
.max-h-\[70vh\]::-webkit-scrollbar {
  width: 6px;
}

.max-h-\[70vh\]::-webkit-scrollbar-track {
  background: transparent;
}

.max-h-\[70vh\]::-webkit-scrollbar-thumb {
  background: rgba(156, 163, 175, 0.5);
  border-radius: 3px;
}

.max-h-\[70vh\]::-webkit-scrollbar-thumb:hover {
  background: rgba(156, 163, 175, 0.7);
}
</style>
