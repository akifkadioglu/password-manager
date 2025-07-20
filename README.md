# OwlLock 🦉🔐

A secure, cross-platform password manager with owl-like vigilance built with Go and Vue.js.

## About

OwlLock is a modern password manager that keeps your passwords secure with advanced encryption and provides a user-friendly interface for managing your credentials.

### Features

- 🔐 **Secure Password Storage**: SQLite database with encryption
- 🎲 **Password Generation**: Cryptographically secure password generator
- 🌐 **Cross-Platform**: Works on Windows, macOS, and Linux
- 🎨 **Modern UI**: Beautiful Vue.js interface with dark/light themes
- 🌍 **Multi-Language**: Supports English and Turkish
- ⚙️ **Configurable**: JSON-based configuration system
- 🦉 **System Fingerprinting**: Unique passwords based on system characteristics

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
