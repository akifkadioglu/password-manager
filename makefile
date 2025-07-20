# Tailwind build
tailwind:
	cd frontend && npx tailwindcss -i ./src/style.css -o ./src/output.css --watch

# Wails dev
wails:
	wails dev
