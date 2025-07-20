<template>
  <div class="min-h-screen">
    <div class="container mx-auto px-4 py-16">
      <div class="text-center">
        <!-- Hero Section -->
        <div class="mb-12">
          <div class="flex justify-center mb-6">
            <div class="relative">
              <img 
                src="../assets/images/icon.png" 
                alt="OwlLock Logo" 
                class="w-24 h-24 rounded-2xl shadow-2xl hover:shadow-3xl transition-shadow duration-300"
              />
              <div class="absolute -top-2 -right-2 w-8 h-8 bg-green-500 rounded-full border-4 border-white dark:border-zinc-900 flex items-center justify-center">
                <svg class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z" />
                </svg>
              </div>
            </div>
          </div>
          <h1 class="text-5xl md:text-6xl font-bold bg-gradient-to-r from-primary-600 via-primary-700 to-primary-800 bg-clip-text text-transparent mb-6 leading-tight">
            {{ $t('landing.title') }}
          </h1>
          <p class="text-xl md:text-2xl text-zinc-600 dark:text-zinc-300 mb-8 max-w-3xl mx-auto leading-relaxed">
            {{ $t('landing.subtitle') }}
          </p>
        </div>

        <!-- CTA Button -->
        <div class="mb-16">
          <router-link to="/password-generator">
            <button class="group relative inline-flex items-center space-x-3 bg-gradient-to-r from-primary-600 to-primary-700 hover:from-primary-700 hover:to-primary-800 text-white px-8 py-4 rounded-xl shadow-lg hover:shadow-2xl transition-all duration-300 text-lg font-semibold transform hover:scale-105">
              <svg class="w-6 h-6 group-hover:rotate-12 transition-transform duration-300" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12,1L21,5V9C21,16.55 15.84,23.74 9,23.74C2.16,23.74 -3,16.55 -3,9V5L12,1M12,7A2,2 0 0,0 10,9A2,2 0 0,0 12,11A2,2 0 0,0 14,9A2,2 0 0,0 12,7Z" />
              </svg>
              <span>{{ $t('landing.getStarted') }}</span>
              <svg class="w-5 h-5 group-hover:translate-x-1 transition-transform duration-300" fill="currentColor" viewBox="0 0 24 24">
                <path d="M4,11V13H16L10.5,18.5L11.92,19.92L19.84,12L11.92,4.08L10.5,5.5L16,11H4Z" />
              </svg>
            </button>
          </router-link>
        </div>
      </div>

      <!-- Features Section -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div class="group bg-white dark:bg-zinc-800 p-8 rounded-2xl shadow-lg hover:shadow-2xl border border-zinc-200 dark:border-zinc-700 transition-all duration-300 hover:transform hover:scale-105">
          <div class="flex items-center justify-center w-16 h-16 bg-gradient-to-br from-blue-500 to-blue-600 rounded-2xl mb-6 mx-auto group-hover:scale-110 transition-transform duration-300">
            <svg class="w-8 h-8 text-white" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12,1L21,5V9C21,16.55 15.84,23.74 9,23.74C2.16,23.74 -3,16.55 -3,9V5L12,1M12,7A2,2 0 0,0 10,9A2,2 0 0,0 12,11A2,2 0 0,0 14,9A2,2 0 0,0 12,7Z" />
            </svg>
          </div>
          <h3 class="text-xl font-semibold text-zinc-900 dark:text-zinc-100 mb-4 text-center">
            {{ $t('landing.features.secureStorage.title') }}
          </h3>
          <p class="text-zinc-600 dark:text-zinc-300 text-center leading-relaxed">
            {{ $t('landing.features.secureStorage.description') }}
          </p>
        </div>

        <div class="group bg-white dark:bg-zinc-800 p-8 rounded-2xl shadow-lg hover:shadow-2xl border border-zinc-200 dark:border-zinc-700 transition-all duration-300 hover:transform hover:scale-105">
          <div class="flex items-center justify-center w-16 h-16 bg-gradient-to-br from-green-500 to-green-600 rounded-2xl mb-6 mx-auto group-hover:scale-110 transition-transform duration-300">
            <svg class="w-8 h-8 text-white" fill="currentColor" viewBox="0 0 24 24">
              <path d="M10,20V14H14V20H19V12H22L12,3L2,12H5V20H10Z" />
            </svg>
          </div>
          <h3 class="text-xl font-semibold text-zinc-900 dark:text-zinc-100 mb-4 text-center">
            {{ $t('landing.features.easyAccess.title') }}
          </h3>
          <p class="text-zinc-600 dark:text-zinc-300 text-center leading-relaxed">
            {{ $t('landing.features.easyAccess.description') }}
          </p>
        </div>

        <div class="group bg-white dark:bg-zinc-800 p-8 rounded-2xl shadow-lg hover:shadow-2xl border border-zinc-200 dark:border-zinc-700 transition-all duration-300 hover:transform hover:scale-105">
          <div class="flex items-center justify-center w-16 h-16 bg-gradient-to-br from-purple-500 to-purple-600 rounded-2xl mb-6 mx-auto group-hover:scale-110 transition-transform duration-300">
            <svg class="w-8 h-8 text-white" fill="currentColor" viewBox="0 0 24 24">
              <path d="M9,22A1,1 0 0,1 8,21V18H4A2,2 0 0,1 2,16V4C2,2.89 2.9,2 4,2H20A2,2 0 0,1 22,4V16A2,2 0 0,1 20,18H13.9L10.2,21.71C10,21.9 9.75,22 9.5,22V22H9M10,16V19.08L13.08,16H20V4H4V16H10Z" />
            </svg>
          </div>
          <h3 class="text-xl font-semibold text-zinc-900 dark:text-zinc-100 mb-4 text-center">
            {{ $t('landing.features.passwordGenerator.title') }}
          </h3>
          <p class="text-zinc-600 dark:text-zinc-300 text-center leading-relaxed">
            {{ $t('landing.features.passwordGenerator.description') }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
