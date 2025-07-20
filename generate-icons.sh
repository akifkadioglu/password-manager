#!/bin/bash

set -e

INPUT="logo.png"
BASENAME="icon"

PNG="${BASENAME}.png"

if command -v convert &> /dev/null; then
    CONVERT="convert"
elif command -v magick &> /dev/null; then
    CONVERT="magick convert"
else
    echo "ImageMagick (convert veya magick) bulunamadı."
    exit 1
fi

mkdir -p assets/icon/windows
mkdir -p assets/icon/linux
mkdir -p assets/icon/macos
mkdir -p assets/icon/web

# PNG 512x512
$CONVERT "$INPUT" -resize 512x512 "$PNG"
echo "[✓] PNG üretildi: $PNG"

# Linux için PNG kopyala
cp "$PNG" "assets/icon/linux/${BASENAME}.png"
echo "[✓] Linux için PNG hazır: assets/icon/linux/${BASENAME}.png"

# Windows için ICO üretimi
$CONVERT "$PNG" -define icon:auto-resize=256,128,64,48,32,16 "assets/icon/windows/${BASENAME}.ico"
echo "[✓] Windows için ICO üretildi: assets/icon/windows/${BASENAME}.ico"

# macOS için iconset klasörünü assets altında oluştur
ICONSET_DIR="assets/icon/macos/${BASENAME}.iconset"
mkdir -p "$ICONSET_DIR"

$CONVERT "$PNG" -resize 16 "$ICONSET_DIR/icon_16x16.png"
$CONVERT "$PNG" -resize 32 "$ICONSET_DIR/icon_16x16@2x.png"
$CONVERT "$PNG" -resize 32 "$ICONSET_DIR/icon_32x32.png"
$CONVERT "$PNG" -resize 64 "$ICONSET_DIR/icon_32x32@2x.png"
$CONVERT "$PNG" -resize 128 "$ICONSET_DIR/icon_128x128.png"
$CONVERT "$PNG" -resize 256 "$ICONSET_DIR/icon_128x128@2x.png"
$CONVERT "$PNG" -resize 256 "$ICONSET_DIR/icon_256x256.png"
$CONVERT "$PNG" -resize 512 "$ICONSET_DIR/icon_256x256@2x.png"
$CONVERT "$PNG" -resize 512 "$ICONSET_DIR/icon_512x512.png"
$CONVERT "$PNG" -resize 1024 "$ICONSET_DIR/icon_512x512@2x.png"

png2icns "assets/icon/macos/${BASENAME}.icns" "$ICONSET_DIR" > /dev/null
echo "[✓] macOS için ICNS üretildi: assets/icon/macos/${BASENAME}.icns"

# WebP formatında 512x512 oluştur
$CONVERT "$INPUT" -resize 512x512 "assets/icon/web/${BASENAME}.webp"
echo "[✓] WebP formatında ikon oluşturuldu: assets/icon/web/${BASENAME}.webp"

# Temizlik
rm -r "$ICONSET_DIR"
rm "$PNG"

echo -e "\n🎉 Tüm ikonlar assets/icon altında hazır:\n"\
"- windows/${BASENAME}.ico\n"\
"- linux/${BASENAME}.png\n"\
"- macos/${BASENAME}.icns\n"\
"- web/${BASENAME}.webp"
