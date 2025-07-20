<template>
  <div class="bg-zinc-100 dark:bg-zinc-900 min-h-screen">
    <div class="container mx-auto px-4 py-16">
      <div class="text-center">
        <h1 class="text-4xl font-bold text-zinc-900 dark:text-zinc-100 mb-8">
          {{ $t('landing.title') }}
        </h1>
        <p class="text-xl text-zinc-700 dark:text-zinc-300 mb-8">
          {{ $t('landing.subtitle') }}
        </p>
        <div class="space-y-4">
          <button class="btn-primary px-6 py-3 rounded-lg hover:bg-primary-600 transition-colors shadow-app-md">
            {{ $t('landing.getStarted') }}
          </button>
        </div>
      </div>

      <div class="mt-16 grid grid-cols-1 md:grid-cols-3 gap-8">
        <div class="bg-white dark:bg-zinc-800 p-6 rounded-lg shadow-app-md border border-zinc-200 dark:border-zinc-700">
          <h3 class="text-xl font-semibold text-zinc-900 dark:text-zinc-100 mb-4">
            {{ $t('landing.features.secureStorage.title') }}
          </h3>
          <p class="text-zinc-700 dark:text-zinc-300">
            {{ $t('landing.features.secureStorage.description') }}
          </p>
        </div>

        <div class="bg-white dark:bg-zinc-800 p-6 rounded-lg shadow-app-md border border-zinc-200 dark:border-zinc-700">
          <h3 class="text-xl font-semibold text-zinc-900 dark:text-zinc-100 mb-4">
            {{ $t('landing.features.easyAccess.title') }}
          </h3>
          <p class="text-zinc-700 dark:text-zinc-300">
            {{ $t('landing.features.easyAccess.description') }}
          </p>
        </div>

        <div class="bg-white dark:bg-zinc-800 p-6 rounded-lg shadow-app-md border border-zinc-200 dark:border-zinc-700">
          <h3 class="text-xl font-semibold text-zinc-900 dark:text-zinc-100 mb-4">
            {{ $t('landing.features.passwordGenerator.title') }}
          </h3>
          <p class="text-zinc-700 dark:text-zinc-300">
            {{ $t('landing.features.passwordGenerator.description') }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
