package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"owllock/internal/config"
	"owllock/internal/controller"
	"owllock/internal/database"
	"owllock/internal/models"
)

// Type aliases for Wails compatibility
type PasswordEntry = models.PasswordEntry
type PasswordGenerationOptions = models.PasswordGenerationOptions
type BackupOption = controller.BackupOption

// BackupData represents the backup file structure (alias for controller type)
type BackupData = controller.BackupData

// EncryptedBackupData represents encrypted backup file structure (alias for controller type)
type EncryptedBackupData = controller.EncryptedBackupData

// BackupMetadata represents backup metadata (alias for controller type)
type BackupMetadata = controller.BackupMetadata

// App struct
type App struct {
	ctx                context.Context
	dbService          *database.Service
	passwordController *controller.PasswordController
	backupController   *controller.BackupController
}

// NewApp creates a new App application struct
func NewApp() *App {
	dbService := database.NewService()
	passwordController := controller.NewPasswordController(dbService)
	backupController := controller.NewBackupController(dbService)

	return &App{
		dbService:          dbService,
		passwordController: passwordController,
		backupController:   backupController,
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

// GetBackupOptions returns available backup options
func (a *App) GetBackupOptions() []BackupOption {
	return a.backupController.GetBackupOptions()
}

// ExportPasswordsEncrypted exports all passwords as an encrypted backup
func (a *App) ExportPasswordsEncrypted(backupPassword string) (EncryptedBackupData, error) {
	return a.backupController.ExportToLocalEncrypted(backupPassword)
}

// ImportPasswordsEncrypted imports passwords from encrypted backup data
func (a *App) ImportPasswordsEncrypted(encryptedBackup EncryptedBackupData, backupPassword string) error {
	return a.backupController.ImportFromEncryptedBackup(encryptedBackup, backupPassword)
}

// ValidateEncryptedBackup validates encrypted backup file format and returns metadata
func (a *App) ValidateEncryptedBackup(backupJSON string) (BackupMetadata, error) {
	return a.backupController.ValidateEncryptedBackup(backupJSON)
}

// GetOwlLockFolderPath returns the path to the owllock folder in user's home directory
func (a *App) GetOwlLockFolderPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user home directory: %w", err)
	}

	owlLockDir := filepath.Join(homeDir, "owllock")

	// Create the directory if it doesn't exist
	if _, err := os.Stat(owlLockDir); os.IsNotExist(err) {
		err = os.MkdirAll(owlLockDir, 0755)
		if err != nil {
			return "", fmt.Errorf("could not create owllock directory: %w", err)
		}
	}

	return owlLockDir, nil
}

// SaveBackupToOwlLockFolder saves the encrypted backup to owllock folder
func (a *App) SaveBackupToOwlLockFolder(encryptedBackup EncryptedBackupData, filename string) error {
	owlLockDir, err := a.GetOwlLockFolderPath()
	if err != nil {
		return err
	}

	filePath := filepath.Join(owlLockDir, filename)

	// Convert backup to JSON
	backupJSON, err := json.Marshal(encryptedBackup)
	if err != nil {
		return fmt.Errorf("could not marshal backup data: %w", err)
	}

	// Write to file
	err = os.WriteFile(filePath, backupJSON, 0644)
	if err != nil {
		return fmt.Errorf("could not write backup file: %w", err)
	}

	return nil
}
