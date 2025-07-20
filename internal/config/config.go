package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// AppConfig holds all application configuration
type AppConfig struct {
	App      AppSettings      `json:"app"`
	Database DatabaseSettings `json:"database"`
	Security SecuritySettings `json:"security"`
	UI       UISettings       `json:"ui"`
}

// AppSettings contains basic application settings
type AppSettings struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Author      Author `json:"author"`
	Debug       bool   `json:"debug"`
}

// Author represents the application author
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	URL   string `json:"url"`
}

// DatabaseSettings contains database configuration
type DatabaseSettings struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	LogLevel string `json:"log_level"`
}

// SecuritySettings contains security-related configuration
type SecuritySettings struct {
	PasswordMinLength int  `json:"password_min_length"`
	PasswordMaxLength int  `json:"password_max_length"`
	UseFingerprint    bool `json:"use_fingerprint"`
	EncryptDatabase   bool `json:"encrypt_database"`
}

// UISettings contains user interface settings
type UISettings struct {
	Theme         string   `json:"theme"`
	Language      string   `json:"language"`
	WindowWidth   int      `json:"window_width"`
	WindowHeight  int      `json:"window_height"`
	Resizable     bool     `json:"resizable"`
	SupportedLang []string `json:"supported_languages"`
}

// Config holds the global configuration instance
var Config *AppConfig

// Load loads configuration from file or creates default configuration
func Load() (*AppConfig, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get config path: %v", err)
	}

	// Create config directory if it doesn't exist
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %v", err)
	}

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create default configuration
		config := getDefaultConfig()
		if err := Save(config, configPath); err != nil {
			return nil, fmt.Errorf("failed to save default config: %v", err)
		}
		Config = config
		return config, nil
	}

	// Load existing configuration
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config AppConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	Config = &config
	return &config, nil
}

// Save saves the configuration to file
func Save(config *AppConfig, configPath string) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}

	return nil
}

// getDefaultConfig returns the default application configuration
func getDefaultConfig() *AppConfig {
	return &AppConfig{
		App: AppSettings{
			Name:        "OwlLock",
			Version:     "1.0.0",
			Description: "A secure, cross-platform password manager with owl-like vigilance built with Go and Vue.js",
			Author: Author{
				Name:  "Akif Kadıoğlu",
				Email: "akif.kadioglu.28@gmail.com",
				URL:   "https://github.com/akifkadioglu",
			},
			Debug: false,
		},
		Database: DatabaseSettings{
			Type:     "sqlite",
			Name:     "passwords.db",
			Path:     "", // Will be set dynamically
			LogLevel: "silent",
		},
		Security: SecuritySettings{
			PasswordMinLength: 6,
			PasswordMaxLength: 128,
			UseFingerprint:    true,
			EncryptDatabase:   false,
		},
		UI: UISettings{
			Theme:         "system",
			Language:      "en",
			WindowWidth:   1200,
			WindowHeight:  800,
			Resizable:     true,
			SupportedLang: []string{"en", "tr"},
		},
	}
}

// getConfigPath returns the appropriate config file path for the current OS
func getConfigPath() (string, error) {
	var configDir string

	switch runtime.GOOS {
	case "windows":
		appData := os.Getenv("APPDATA")
		if appData == "" {
			return "", fmt.Errorf("APPDATA environment variable not set")
		}
		configDir = filepath.Join(appData, "OwlLock")
	case "darwin":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configDir = filepath.Join(homeDir, "Library", "Application Support", "OwlLock")
	default: // Linux and other Unix-like systems
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		// Check for XDG_CONFIG_HOME first
		if xdgConfig := os.Getenv("XDG_CONFIG_HOME"); xdgConfig != "" {
			configDir = filepath.Join(xdgConfig, "owllock")
		} else {
			configDir = filepath.Join(homeDir, ".config", "owllock")
		}
	}

	return filepath.Join(configDir, "config.json"), nil
}

// GetDataDir returns the data directory path based on configuration
func GetDataDir() (string, error) {
	switch runtime.GOOS {
	case "windows":
		appData := os.Getenv("APPDATA")
		if appData == "" {
			return "", fmt.Errorf("APPDATA environment variable not set")
		}
		return filepath.Join(appData, "OwlLock"), nil
	case "darwin":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, "Library", "Application Support", "OwlLock"), nil
	default: // Linux and other Unix-like systems
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		// Check for XDG_DATA_HOME first
		if xdgData := os.Getenv("XDG_DATA_HOME"); xdgData != "" {
			return filepath.Join(xdgData, "owllock"), nil
		}
		return filepath.Join(homeDir, ".local", "share", "owllock"), nil
	}
}

// UpdateConfig updates a specific configuration section
func UpdateConfig(updater func(*AppConfig) error) error {
	if Config == nil {
		return fmt.Errorf("configuration not loaded")
	}

	if err := updater(Config); err != nil {
		return err
	}

	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	return Save(Config, configPath)
}
