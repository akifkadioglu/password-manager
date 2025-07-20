package main

import (
	"context"
	"fmt"

	"owllock/internal/config"
	"owllock/internal/controller"
	"owllock/internal/database"
	"owllock/internal/models"
)

// Type aliases for Wails compatibility
type PasswordEntry = models.PasswordEntry
type PasswordGenerationOptions = models.PasswordGenerationOptions

// App struct
type App struct {
	ctx                context.Context
	dbService          *database.Service
	passwordController *controller.PasswordController
}

// NewApp creates a new App application struct
func NewApp() *App {
	dbService := database.NewService()
	passwordController := controller.NewPasswordController(dbService)

	return &App{
		dbService:          dbService,
		passwordController: passwordController,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Initialize configuration
	if _, err := config.Load(); err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
	}

	// Initialize database
	if err := a.dbService.Initialize(); err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetSystemFingerprint delegates to password controller
func (a *App) GetSystemFingerprint() string {
	return a.passwordController.GetSystemFingerprint()
}

// GeneratePassword delegates to password controller
func (a *App) GeneratePassword(options PasswordGenerationOptions) (string, error) {
	return a.passwordController.GeneratePassword(options)
}

// SavePasswordEntry delegates to password controller
func (a *App) SavePasswordEntry(entry PasswordEntry) (PasswordEntry, error) {
	return a.passwordController.SavePasswordEntry(entry)
}

// GetPasswordEntries delegates to password controller
func (a *App) GetPasswordEntries() ([]PasswordEntry, error) {
	return a.passwordController.GetPasswordEntries()
}

// DeletePasswordEntry delegates to password controller
func (a *App) DeletePasswordEntry(id string) error {
	return a.passwordController.DeletePasswordEntry(id)
}

// GetAppInfo returns application information
func (a *App) GetAppInfo() map[string]interface{} {
	if config.Config == nil {
		return map[string]interface{}{
			"name":    "OwlLock",
			"version": "1.0.0",
			"author":  "Akif Kadıoğlu",
		}
	}

	return map[string]interface{}{
		"name":        config.Config.App.Name,
		"version":     config.Config.App.Version,
		"description": config.Config.App.Description,
		"author":      config.Config.App.Author.Name,
		"email":       config.Config.App.Author.Email,
	}
}

// GetAppSettings returns current application settings
func (a *App) GetAppSettings() map[string]interface{} {
	if config.Config == nil {
		return map[string]interface{}{}
	}

	return map[string]interface{}{
		"theme":             config.Config.UI.Theme,
		"language":          config.Config.UI.Language,
		"supportedLang":     config.Config.UI.SupportedLang,
		"passwordMinLength": config.Config.Security.PasswordMinLength,
		"passwordMaxLength": config.Config.Security.PasswordMaxLength,
		"useFingerprint":    config.Config.Security.UseFingerprint,
	}
}

// UpdateAppSettings updates application settings
func (a *App) UpdateAppSettings(settings map[string]interface{}) error {
	return config.UpdateConfig(func(cfg *config.AppConfig) error {
		if theme, ok := settings["theme"].(string); ok {
			cfg.UI.Theme = theme
		}
		if language, ok := settings["language"].(string); ok {
			cfg.UI.Language = language
		}
		if useFingerprint, ok := settings["useFingerprint"].(bool); ok {
			cfg.Security.UseFingerprint = useFingerprint
		}
		return nil
	})
}
