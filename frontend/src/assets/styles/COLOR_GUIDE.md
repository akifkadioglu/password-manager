# Renk Paleti Kullanım Rehberi

Bu dosya, password manager projesinde kullanılan renk paletinin nasıl kullanılacağını açıklar.

## Renk Sistemi

### Ana Renkler (Primary)
- **primary-500**: Ana mavi renk - Butonlar, linkler, vurgulanacak öğeler
- **primary-600**: Hover durumları için koyu mavi
- **primary-50**: Açık mavi - Arka plan vurguları için

### İkincil Renkler (Secondary) 
- **secondary-500**: Yeşil - Başarı durumları, ikincil butonlar
- **secondary-600**: Koyu yeşil - Hover durumları

### Vurgu Renkleri (Accent)
- **accent-400**: Altın sarı - Önemli bildirimler, premium özellikler

### Durum Renkleri
- **success-500**: Başarı mesajları
- **warning-500**: Uyarı mesajları  
- **error-500**: Hata mesajları
- **info-500**: Bilgi mesajları

### Nötr Renkler
- **neutral-50 to 950**: Metin, kenarlar, arka planlar için gri tonları

## CSS Değişkenleri ile Kullanım

### Arka Plan Renkleri
```css
background-color: var(--bg-primary);     /* Ana arka plan */
background-color: var(--bg-secondary);   /* İkincil arka plan */
background-color: var(--bg-card);        /* Kart arka planları */
```

### Metin Renkleri
```css
color: var(--text-primary);    /* Ana metin */
color: var(--text-secondary);  /* İkincil metin */
color: var(--text-tertiary);   /* Üçüncül metin (placeholders) */
```

### Buton Stilleri
```css
.btn-primary     /* Ana buton stili */
.btn-secondary   /* İkincil buton stili */
```

## Tailwind CSS ile Kullanım

### Arka Plan
```html
<div class="bg-app-primary">     <!-- Ana arka plan -->
<div class="bg-app-card">        <!-- Kart arka planı -->
<div class="bg-primary-500">     <!-- Mavi arka plan -->
```

### Metin
```html
<p class="text-app-primary">     <!-- Ana metin rengi -->
<p class="text-app-secondary">   <!-- İkincil metin rengi -->
<p class="text-primary-500">     <!-- Mavi metin -->
```

### Kenarlar
```html
<div class="border-app-light">   <!-- Açık kenar -->
<div class="border-primary-500"> <!-- Mavi kenar -->
```

### Gölgeler
```html
<div class="shadow-app-md">      <!-- Orta gölge -->
<div class="shadow-app-lg">      <!-- Büyük gölge -->
```

## Tema Desteği

Sistem otomatik olarak light/dark theme destekler:
- Light tema: `data-theme="light"` veya varsayılan
- Dark tema: `data-theme="dark"`

## Yeni Renkler Ekleme

Yeni renk eklemek için:

1. `colors.css` dosyasına CSS değişkeni ekle
2. `tailwind.config.js` dosyasında ilgili bölüme ekle
3. Bu rehberi güncelle

## Önerilen Kullanım

### Butonlar
- Ana buton: `btn-primary` class'ı veya `bg-primary-500 hover:bg-primary-600`
- İkincil buton: `btn-secondary` class'ı
- Tehlike butonu: `bg-error-500 hover:bg-error-600`

### Kartlar
- Kart container: `card` class'ı veya `bg-app-card border-app-light`
- Kart gölgesi: `shadow-app-md`

### Form Elemanları
- Input arka planı: `bg-app-card`
- Input kenarı: `border-app-medium focus:border-primary-500`
- Placeholder: `text-app-tertiary`

### Navigasyon
- Nav arka planı: `bg-app-card`
- Nav kenarı: `border-app-light`
- Aktif link: `text-primary-500`

Bu rehberi takip ederek tutarlı ve profesyonel bir görünüm elde edebilirsiniz.
