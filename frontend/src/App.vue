<script setup>
import Navbar from './components/Navbar.vue'
</script>

<template>
  <div class="select-none min-h-screen bg-gradient-to-br from-zinc-50 via-white to-zinc-100 dark:from-zinc-900 dark:via-zinc-800 dark:to-zinc-900">
    <!-- Navigation -->
    <Navbar />
    
    <!-- Main Content -->
    <main class="relative">
      <router-view />
    </main>
  </div>
</template>

<style>
/* Global styles for smooth transitions */
* {
  transition-property: color, background-color, border-color;
  transition-duration: 200ms;
}

/* Custom scrollbar */
::-webkit-scrollbar {
  width: 8px;
}

::-webkit-scrollbar-track {
  background-color: #f4f4f5;
}

.dark ::-webkit-scrollbar-track {
  background-color: #27272a;
}

::-webkit-scrollbar-thumb {
  background-color: #d4d4d8;
  border-radius: 9999px;
}

.dark ::-webkit-scrollbar-thumb {
  background-color: #52525b;
}

::-webkit-scrollbar-thumb:hover {
  background-color: #a1a1aa;
}

.dark ::-webkit-scrollbar-thumb:hover {
  background-color: #71717a;
}

/* Focus styles for accessibility */
*:focus {
  outline: none;
  box-shadow: 0 0 0 2px rgba(245, 158, 11, 0.5), 0 0 0 4px #ffffff;
}

.dark *:focus {
  box-shadow: 0 0 0 2px rgba(245, 158, 11, 0.5), 0 0 0 4px #18181b;
}

/* Selection styles */
::selection {
  background-color: #fef3c7;
  color: #78350f;
}

.dark ::selection {
  background-color: #451a03;
  color: #fef3c7;
}
</style>
