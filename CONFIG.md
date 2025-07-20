# Configuration Guide

## Overview

OwlLock uses a JSON-based configuration system that allows you to customize various aspects of the application.

## Configuration File Location

The configuration file is automatically created in the appropriate location for your operating system:

### Windows
```
%APPDATA%\OwlLock\config.json
```

### macOS
```
~/Library/Application Support/OwlLock/config.json
```

### Linux
```
~/.config/owllock/config.json
```
or if `XDG_CONFIG_HOME` is set:
```
$XDG_CONFIG_HOME/owllock/config.json
```

## Configuration Options

### App Settings
- `name`: Application name
- `version`: Current version
- `description`: Application description
- `author`: Author information
- `debug`: Enable debug mode (default: false)

### Database Settings
- `type`: Database type (currently only "sqlite" is supported)
- `name`: Database filename (default: "passwords.db")
- `log_level`: Database logging level ("silent" or "info")

### Security Settings
- `password_min_length`: Minimum password length (default: 6)
- `password_max_length`: Maximum password length (default: 128)
- `use_fingerprint`: Use system fingerprint for password generation (default: true)
- `encrypt_database`: Enable database encryption (default: false)

### UI Settings
- `theme`: UI theme ("light", "dark", or "system")
- `language`: Default language ("en" or "tr")
- `window_width`: Default window width (default: 1200)
- `window_height`: Default window height (default: 800)
- `resizable`: Allow window resizing (default: true)
- `supported_languages`: List of supported languages

## Example Configuration

See `config.example.json` for a complete example configuration file.

## Modifying Configuration

The configuration file is created automatically with default values on first run. You can:

1. Edit the file directly using any text editor
2. Use the application's settings interface (if available)
3. Use the API methods to update settings programmatically

## Environment Variables

Some settings can be overridden using environment variables:

- `XDG_CONFIG_HOME`: Custom config directory (Linux only)
- `XDG_DATA_HOME`: Custom data directory (Linux only)

## Data Directory

Application data (including the database) is stored in:

### Windows
```
%APPDATA%\OwlLock\
```

### macOS
```
~/Library/Application Support/OwlLock/
```

### Linux
```
~/.local/share/owllock/
```
or if `XDG_DATA_HOME` is set:
```
$XDG_DATA_HOME/owllock/
```
