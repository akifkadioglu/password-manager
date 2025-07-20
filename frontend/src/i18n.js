import { createI18n } from 'vue-i18n'
import en from './locales/en.json'
import tr from './locales/tr.json'

const messages = {
  en,
  tr
}

// Get the user's preferred language from localStorage or browser
const getDefaultLocale = () => {
  const savedLocale = localStorage.getItem('locale')
  if (savedLocale && messages[savedLocale]) {
    return savedLocale
  }
  
  // Get browser language
  const browserLang = navigator.language.split('-')[0]
  if (messages[browserLang]) {
    return browserLang
  }
  
  return 'en' // fallback
}

const i18n = createI18n({
  legacy: false,
  locale: getDefaultLocale(),
  fallbackLocale: 'en',
  messages,
})

export default i18n
