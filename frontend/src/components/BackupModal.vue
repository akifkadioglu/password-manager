<template>
  <div v-if="isOpen" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white dark:bg-zinc-800 rounded-2xl p-8 w-full max-w-md mx-4 shadow-2xl">
      <!-- Header -->
      <div class="flex items-center justify-between mb-6">
        <h2 class="text-2xl font-bold text-zinc-900 dark:text-zinc-100">
          {{ currentStep === 'options' ? $t('backup.title') : $t('backup.exportTitle') }}
        </h2>
        <button
          @click="closeModal"
          class="text-zinc-400 hover:text-zinc-600 dark:hover:text-zinc-300 transition-colors"
        >
          <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
            <path d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z"/>
          </svg>
        </button>
      </div>

      <!-- Options Step -->
      <div v-if="currentStep === 'options'">
        <p class="text-zinc-600 dark:text-zinc-400 mb-6">
          {{ $t('backup.description') }}
        </p>

        <div class="space-y-4">
          <!-- Export Option -->
          <button
            @click="showExportStep"
            class="w-full p-6 text-left bg-gradient-to-r from-primary-50 to-primary-100 dark:from-primary-900/20 dark:to-primary-800/20 border-2 border-primary-200 dark:border-primary-700 rounded-xl hover:from-primary-100 hover:to-primary-150 dark:hover:from-primary-900/30 dark:hover:to-primary-800/30 transition-all duration-200 group"
          >
            <div class="flex items-center space-x-4">
              <div class="w-12 h-12 bg-primary-500 rounded-xl flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z"/>
                </svg>
              </div>
              <div class="flex-1">
                <h3 class="font-semibold text-zinc-900 dark:text-zinc-100 group-hover:text-primary-600 dark:group-hover:text-primary-400">
                  {{ $t('backup.exportOption') }}
                </h3>
                <p class="text-sm text-zinc-600 dark:text-zinc-400 mt-1">
                  {{ $t('backup.exportDescription') }}
                </p>
              </div>
              <svg class="w-5 h-5 text-zinc-400 group-hover:text-primary-500" fill="currentColor" viewBox="0 0 24 24">
                <path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z"/>
              </svg>
            </div>
          </button>

          <!-- Import Option -->
          <button
            @click="triggerFileImport"
            class="w-full p-6 text-left bg-gradient-to-r from-green-50 to-green-100 dark:from-green-900/20 dark:to-green-800/20 border-2 border-green-200 dark:border-green-700 rounded-xl hover:from-green-100 hover:to-green-150 dark:hover:from-green-900/30 dark:hover:to-green-800/30 transition-all duration-200 group"
          >
            <div class="flex items-center space-x-4">
              <div class="w-12 h-12 bg-green-500 rounded-xl flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z"/>
                </svg>
              </div>
              <div class="flex-1">
                <h3 class="font-semibold text-zinc-900 dark:text-zinc-100 group-hover:text-green-600 dark:group-hover:text-green-400">
                  {{ $t('backup.importOption') }}
                </h3>
                <p class="text-sm text-zinc-600 dark:text-zinc-400 mt-1">
                  {{ $t('backup.importDescription') }}
                </p>
              </div>
              <svg class="w-5 h-5 text-zinc-400 group-hover:text-green-500" fill="currentColor" viewBox="0 0 24 24">
                <path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z"/>
              </svg>
            </div>
          </button>
        </div>
      </div>

      <!-- Export Step -->
      <div v-else-if="currentStep === 'export'">
        <div class="mb-6">
          <label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2">
            {{ $t('backup.backupPassword') }}
          </label>
          <div class="relative">
            <input
              v-model="backupPassword"
              :type="showBackupPassword ? 'text' : 'password'"
              :placeholder="$t('backup.enterBackupPassword')"
              class="w-full px-4 py-3 border border-zinc-300 dark:border-zinc-600 rounded-lg bg-zinc-50 dark:bg-zinc-700 text-zinc-900 dark:text-zinc-100 placeholder-zinc-500 dark:placeholder-zinc-400 focus:ring-2 focus:ring-primary-500 focus:border-primary-500 transition-colors pr-12"
            />
            <button
              @click="showBackupPassword = !showBackupPassword"
              class="absolute inset-y-0 right-0 pr-3 flex items-center"
              type="button"
            >
              <svg class="w-5 h-5 text-zinc-400" fill="currentColor" viewBox="0 0 24 24">
                <path :d="showBackupPassword ? 'M12,9A3,3 0 0,0 9,12A3,3 0 0,0 12,15A3,3 0 0,0 15,12A3,3 0 0,0 12,9M12,17A5,5 0 0,1 7,12A5,5 0 0,1 12,7A5,5 0 0,1 17,12A5,5 0 0,1 12,17M12,4.5C7,4.5 2.73,7.61 1,12C2.73,16.39 7,19.5 12,19.5C17,19.5 21.27,16.39 23,12C21.27,7.61 17,4.5 12,4.5Z' : 'M11.83,9L15,12.16C15,12.11 15,12.05 15,12A3,3 0 0,0 12,9C11.94,9 11.89,9 11.83,9M7.53,9.8L9.08,11.35C9.03,11.56 9,11.77 9,12A3,3 0 0,0 12,15C12.22,15 12.44,14.97 12.65,14.92L14.2,16.47C13.53,16.8 12.79,17 12,17A5,5 0 0,1 7,12C7,11.21 7.2,10.47 7.53,9.8M2,4.27L4.28,6.55L4.73,7C3.08,8.3 1.78,10 1,12C2.73,16.39 7,19.5 12,19.5C13.55,19.5 15.03,19.2 16.38,18.66L16.81,19.09L19.73,22L21,20.73L3.27,3M12,7A5,5 0 0,1 17,12C17,12.64 16.87,13.26 16.64,13.82L19.57,16.75C21.07,15.5 22.27,13.86 23,12C21.27,7.61 17,4.5 12,4.5C10.6,4.5 9.26,4.75 8,5.2L10.17,7.35C10.76,7.13 11.37,7 12,7Z'"/>
              </svg>
            </button>
          </div>
          <p class="text-xs text-zinc-500 dark:text-zinc-400 mt-2">
            {{ $t('backup.passwordHint') }}
          </p>
        </div>

        <!-- Action Buttons -->
        <div class="flex space-x-3">
          <button
            @click="goBack"
            class="flex-1 px-4 py-3 border border-zinc-300 dark:border-zinc-600 text-zinc-700 dark:text-zinc-300 rounded-lg hover:bg-zinc-50 dark:hover:bg-zinc-700 transition-colors"
          >
            {{ $t('common.back') }}
          </button>
          <button
            @click="handleExport"
            :disabled="!backupPassword.trim() || exporting"
            class="flex-1 px-4 py-3 bg-primary-600 hover:bg-primary-700 disabled:bg-zinc-300 disabled:cursor-not-allowed text-white rounded-lg transition-colors flex items-center justify-center"
          >
            <div v-if="exporting" class="flex items-center space-x-2">
              <svg class="w-4 h-4 animate-spin" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12,4V2A10,10 0 0,0 2,12H4A8,8 0 0,1 12,4Z"/>
              </svg>
              <span>{{ $t('backup.exporting') }}</span>
            </div>
            <span v-else>{{ $t('backup.export') }}</span>
          </button>
        </div>
      </div>

      <!-- Import Step -->
      <div v-else-if="currentStep === 'import'">
        <div class="mb-6">
          <div class="border-2 border-dashed border-zinc-300 dark:border-zinc-600 rounded-lg p-8 text-center">
            <div class="mb-4">
              <svg class="w-16 h-16 text-zinc-400 mx-auto" fill="currentColor" viewBox="0 0 24 24">
                <path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z"/>
              </svg>
            </div>
            <p class="text-zinc-600 dark:text-zinc-400 mb-4">
              {{ selectedFile ? selectedFile.name : $t('backup.selectFile') }}
            </p>
            <p v-if="selectedFile && selectedFile.name.endsWith('.owllock')" class="text-xs text-primary-600 dark:text-primary-400 mb-2">
              {{ $t('backup.owllockFileDetected') }}
            </p>
            <button
              @click="triggerFileImport"
              class="px-6 py-3 bg-primary-600 hover:bg-primary-700 text-white rounded-lg transition-colors"
            >
              {{ $t('backup.chooseFile') }}
            </button>
          </div>
        </div>

        <!-- Import Password -->
        <div v-if="selectedFile" class="mb-6">
          <label class="block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2">
            {{ $t('backup.importPassword') }}
          </label>
          <div class="relative">
            <input
              v-model="importPassword"
              :type="showImportPassword ? 'text' : 'password'"
              :placeholder="$t('backup.enterImportPassword')"
              class="w-full px-4 py-3 border border-zinc-300 dark:border-zinc-600 rounded-lg bg-zinc-50 dark:bg-zinc-700 text-zinc-900 dark:text-zinc-100 placeholder-zinc-500 dark:placeholder-zinc-400 focus:ring-2 focus:ring-primary-500 focus:border-primary-500 transition-colors pr-12"
            />
            <button
              @click="showImportPassword = !showImportPassword"
              class="absolute inset-y-0 right-0 pr-3 flex items-center"
              type="button"
            >
              <svg class="w-5 h-5 text-zinc-400" fill="currentColor" viewBox="0 0 24 24">
                <path :d="showImportPassword ? 'M12,9A3,3 0 0,0 9,12A3,3 0 0,0 12,15A3,3 0 0,0 15,12A3,3 0 0,0 12,9M12,17A5,5 0 0,1 7,12A5,5 0 0,1 12,7A5,5 0 0,1 17,12A5,5 0 0,1 12,17M12,4.5C7,4.5 2.73,7.61 1,12C2.73,16.39 7,19.5 12,19.5C17,19.5 21.27,16.39 23,12C21.27,7.61 17,4.5 12,4.5Z' : 'M11.83,9L15,12.16C15,12.11 15,12.05 15,12A3,3 0 0,0 12,9C11.94,9 11.89,9 11.83,9M7.53,9.8L9.08,11.35C9.03,11.56 9,11.77 9,12A3,3 0 0,0 12,15C12.22,15 12.44,14.97 12.65,14.92L14.2,16.47C13.53,16.8 12.79,17 12,17A5,5 0 0,1 7,12C7,11.21 7.2,10.47 7.53,9.8M2,4.27L4.28,6.55L4.73,7C3.08,8.3 1.78,10 1,12C2.73,16.39 7,19.5 12,19.5C13.55,19.5 15.03,19.2 16.38,18.66L16.81,19.09L19.73,22L21,20.73L3.27,3M12,7A5,5 0 0,1 17,12C17,12.64 16.87,13.26 16.64,13.82L19.57,16.75C21.07,15.5 22.27,13.86 23,12C21.27,7.61 17,4.5 12,4.5C10.6,4.5 9.26,4.75 8,5.2L10.17,7.35C10.76,7.13 11.37,7 12,7Z'"/>
              </svg>
            </button>
          </div>
        </div>

        <!-- Action Buttons -->
        <div class="flex space-x-3">
          <button
            @click="goBack"
            class="flex-1 px-4 py-3 border border-zinc-300 dark:border-zinc-600 text-zinc-700 dark:text-zinc-300 rounded-lg hover:bg-zinc-50 dark:hover:bg-zinc-700 transition-colors"
          >
            {{ $t('common.back') }}
          </button>
          <button
            @click="handleImport"
            :disabled="!selectedFile || !importPassword.trim() || importing"
            class="flex-1 px-4 py-3 bg-green-600 hover:bg-green-700 disabled:bg-zinc-300 disabled:cursor-not-allowed text-white rounded-lg transition-colors flex items-center justify-center"
          >
            <div v-if="importing" class="flex items-center space-x-2">
              <svg class="w-4 h-4 animate-spin" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12,4V2A10,10 0 0,0 2,12H4A8,8 0 0,1 12,4Z"/>
              </svg>
              <span>{{ $t('backup.importing') }}</span>
            </div>
            <span v-else>{{ $t('backup.import') }}</span>
          </button>
        </div>
      </div>

      <!-- Hidden file input -->
      <input
        ref="fileInput"
        type="file"
        accept=".json,.owllock"
        @change="handleFileSelect"
        class="hidden"
      />
    </div>
  </div>
</template>

<script>
import { ExportPasswordsEncrypted, ImportPasswordsEncrypted, ValidateEncryptedBackup, SaveBackupToOwlLockFolder } from '../../wailsjs/go/main/App'

export default {
  name: 'BackupModal',
  props: {
    isOpen: {
      type: Boolean,
      default: false
    },
    backupOptions: {
      type: Array,
      default: () => []
    }
  },
  emits: ['close', 'refresh-passwords'],
  data() {
    return {
      currentStep: 'options', // 'options', 'export', 'import'
      backupPassword: '',
      showBackupPassword: false,
      importPassword: '',
      showImportPassword: false,
      selectedFile: null,
      exporting: false,
      importing: false
    }
  },
  methods: {
    closeModal() {
      this.resetModal()
      this.$emit('close')
    },
    resetModal() {
      this.currentStep = 'options'
      this.backupPassword = ''
      this.showBackupPassword = false
      this.importPassword = ''
      this.showImportPassword = false
      this.selectedFile = null
      this.exporting = false
      this.importing = false
    },
    showExportStep() {
      this.currentStep = 'export'
    },
    showImportStep() {
      this.currentStep = 'import'
    },
    goBack() {
      if (this.currentStep === 'export' || this.currentStep === 'import') {
        this.currentStep = 'options'
      }
    },
    triggerFileImport() {
      if (this.currentStep === 'options') {
        this.showImportStep()
      } else {
        this.$refs.fileInput.click()
      }
    },
    handleFileSelect(event) {
      const file = event.target.files[0]
      if (file) {
        this.selectedFile = file
      }
    },
    async handleExport() {
      if (!this.backupPassword.trim()) return

      this.exporting = true
      try {
        // Call backend to create encrypted backup
        const encryptedBackup = await ExportPasswordsEncrypted(this.backupPassword)
        
        // Create filename with timestamp
        const now = new Date()
        const timestamp = now.toISOString().replace(/[:.]/g, '-').slice(0, 19)
        const filename = `owllock-backup-${timestamp}.owllock`
        
        try {
          // Try to save to owllock folder first
          await SaveBackupToOwlLockFolder(encryptedBackup, filename)
          
          // Show success message for owllock folder
          this.$emit('show-notification', {
            type: 'success',
            message: this.$t('backup.exportSuccess')
          })
        } catch (folderError) {
          // If owllock folder save fails, fallback to browser download
          console.warn('Failed to save to owllock folder, falling back to download:', folderError)
          
          // Create blob and download
          const blob = new Blob([JSON.stringify(encryptedBackup, null, 2)], {
            type: 'application/json'
          })
          
          // Create download link
          const url = URL.createObjectURL(blob)
          const link = document.createElement('a')
          link.href = url
          link.download = filename
          document.body.appendChild(link)
          link.click()
          document.body.removeChild(link)
          URL.revokeObjectURL(url)
          
          // Show success message for download
          this.$emit('show-notification', {
            type: 'success',
            message: this.$t('backup.exportSuccess')
          })
        }
        
        this.closeModal()
      } catch (error) {
        console.error('Export error:', error)
        this.$emit('show-notification', {
          type: 'error',
          message: this.$t('backup.exportError') + ': ' + error.message
        })
      } finally {
        this.exporting = false
      }
    },
    async handleImport() {
      if (!this.selectedFile || !this.importPassword.trim()) return

      this.importing = true
      try {
        // Read file content
        const fileContent = await this.readFileAsText(this.selectedFile)
        
        // Parse JSON
        const encryptedBackup = JSON.parse(fileContent)
        
        // Validate backup format first
        try {
          await ValidateEncryptedBackup(fileContent)
        } catch (validationError) {
          throw new Error(this.$t('backup.invalidFileFormat'))
        }
        
        // Import the encrypted backup
        await ImportPasswordsEncrypted(encryptedBackup, this.importPassword)
        
        // Show success message with count if available
        let successMessage = this.$t('backup.importSuccess')
        try {
          const passwordCount = encryptedBackup.passwords ? encryptedBackup.passwords.length : 'unknown'
          successMessage = this.$t('backup.importSuccessWithCount', { count: passwordCount })
        } catch (e) {
          // Use default message if count extraction fails
        }
        
        this.$emit('show-notification', {
          type: 'success',
          message: successMessage
        })
        
        // Refresh passwords list
        this.$emit('refresh-passwords')
        
        // Close modal
        this.closeModal()
      } catch (error) {
        console.error('Import error:', error)
        let errorMessage = this.$t('backup.importError')
        
        if (error.message.includes('wrong password')) {
          errorMessage = this.$t('backup.wrongPassword')
        } else if (error.message.includes('invalid backup format') || error.message.includes('corrupted backup data')) {
          errorMessage = this.$t('backup.invalidFileFormat')
        } else if (error.message.includes('no passwords found')) {
          errorMessage = this.$t('backup.noPasswordsInFile')
        } else if (error.message.includes('import failed')) {
          errorMessage = this.$t('backup.importFailed') + ': ' + error.message.replace('import failed: ', '')
        }
        
        this.$emit('show-notification', {
          type: 'error',
          message: errorMessage
        })
      } finally {
        this.importing = false
      }
    },
    readFileAsText(file) {
      return new Promise((resolve, reject) => {
        const reader = new FileReader()
        reader.onload = (e) => resolve(e.target.result)
        reader.onerror = (e) => reject(e)
        reader.readAsText(file)
      })
    }
  }
}
</script>
