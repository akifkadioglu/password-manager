package controller

import (
	"encoding/json"
	"fmt"
	"time"

	"owllock/internal/database"
	"owllock/internal/models"
	"owllock/pkg/utils"
)

// BackupOption represents available backup options
type BackupOption struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Available   bool   `json:"available"`
}

// BackupController handles backup and restore operations
type BackupController struct {
	dbService         *database.Service
	encryptionService *utils.EncryptionService
}

// NewBackupController creates a new BackupController
func NewBackupController(dbService *database.Service) *BackupController {
	return &BackupController{
		dbService:         dbService,
		encryptionService: utils.NewEncryptionService(),
	}
}

// BackupData represents the backup file structure
type BackupData struct {
	Version    string                 `json:"version"`
	ExportDate string                 `json:"exportDate"`
	AppName    string                 `json:"appName"`
	Passwords  []models.PasswordEntry `json:"passwords"`
}

// EncryptedBackupData represents encrypted backup file structure
type EncryptedBackupData struct {
	Version          string `json:"version"`
	ExportDate       string `json:"exportDate"`
	AppName          string `json:"appName"`
	Encrypted        bool   `json:"encrypted"`
	EncryptedPayload string `json:"encryptedPayload"`
	EncryptionMethod string `json:"encryptionMethod"`
}

// BackupOptions represents backup creation options
type BackupOptions struct {
	Encrypt  bool   `json:"encrypt"`
	Password string `json:"password,omitempty"`
}

// BackupMetadata represents backup metadata
type BackupMetadata struct {
	PasswordCount int    `json:"passwordCount"`
	ExportDate    string `json:"exportDate"`
	Version       string `json:"version"`
	Encrypted     bool   `json:"encrypted"`
}

// GetBackupOptions returns available backup options
func (bc *BackupController) GetBackupOptions() []BackupOption {
	return []BackupOption{
		{
			ID:          "local",
			Name:        "Local Computer",
			Description: "Save backup file to your computer",
			Icon:        "computer",
			Available:   true,
		},
		{
			ID:          "google_drive",
			Name:        "Google Drive",
			Description: "Save backup to your Google Drive",
			Icon:        "google-drive",
			Available:   false, // Will be implemented later
		},
		{
			ID:          "dropbox",
			Name:        "Dropbox",
			Description: "Save backup to your Dropbox",
			Icon:        "dropbox",
			Available:   false, // Will be implemented later
		},
		{
			ID:          "onedrive",
			Name:        "OneDrive",
			Description: "Save backup to your OneDrive",
			Icon:        "onedrive",
			Available:   false, // Will be implemented later
		},
	}
}

// CreateBackup creates a backup of all passwords
func (bc *BackupController) CreateBackup() (BackupData, error) {
	passwords, err := bc.dbService.GetPasswordEntries()
	if err != nil {
		return BackupData{}, fmt.Errorf("failed to get password entries: %w", err)
	}

	// Clean passwords by removing IDs to avoid conflicts during import
	cleanedPasswords := make([]models.PasswordEntry, len(passwords))
	for i, password := range passwords {
		cleanedPasswords[i] = models.PasswordEntry{
			ID:          "", // Remove ID to avoid UNIQUE constraint issues
			Platform:    password.Platform,
			Email:       password.Email,
			Username:    password.Username,
			Password:    password.Password,
			Description: password.Description,
			CreatedAt:   password.CreatedAt,
			UpdatedAt:   password.UpdatedAt,
		}
	}

	backup := BackupData{
		Version:    "1.0",
		ExportDate: time.Now().Format(time.RFC3339),
		AppName:    "OwlLock Password Manager",
		Passwords:  cleanedPasswords,
	}

	return backup, nil
}

// CreateEncryptedBackup creates an encrypted backup with password
func (bc *BackupController) CreateEncryptedBackup(backupPassword string) (EncryptedBackupData, error) {
	// Create regular backup first
	backup, err := bc.CreateBackup()
	if err != nil {
		return EncryptedBackupData{}, err
	}

	// Convert to JSON
	backupJSON, err := json.Marshal(backup)
	if err != nil {
		return EncryptedBackupData{}, fmt.Errorf("failed to marshal backup data: %w", err)
	}

	// Encrypt the JSON data
	encryptedPayload, err := bc.encryptionService.EncryptBackupToBase64(backupJSON, backupPassword)
	if err != nil {
		return EncryptedBackupData{}, fmt.Errorf("failed to encrypt backup: %w", err)
	}

	// Create encrypted backup structure
	encryptedBackup := EncryptedBackupData{
		Version:          "1.0",
		ExportDate:       backup.ExportDate,
		AppName:          backup.AppName,
		Encrypted:        true,
		EncryptedPayload: encryptedPayload,
		EncryptionMethod: "AES-256-CFB-PBKDF2",
	}

	return encryptedBackup, nil
}

// ExportToLocal exports backup data for local download (unencrypted)
func (bc *BackupController) ExportToLocal() (BackupData, error) {
	return bc.CreateBackup()
}

// ExportToLocalEncrypted exports encrypted backup data for local download
func (bc *BackupController) ExportToLocalEncrypted(backupPassword string) (EncryptedBackupData, error) {
	return bc.CreateEncryptedBackup(backupPassword)
}

// GenerateBackupPassword generates a secure backup password
func (bc *BackupController) GenerateBackupPassword() (string, error) {
	return bc.encryptionService.GenerateSecurePassword(16)
}

// ImportFromBackup imports passwords from backup data (unencrypted)
func (bc *BackupController) ImportFromBackup(backupData BackupData) error {
	// Validate backup data
	if backupData.Version == "" {
		return fmt.Errorf("invalid backup format: missing version")
	}

	if len(backupData.Passwords) == 0 {
		return fmt.Errorf("no passwords found in backup file")
	}

	return bc.importPasswords(backupData.Passwords)
}

// ImportFromEncryptedBackup imports passwords from encrypted backup data
func (bc *BackupController) ImportFromEncryptedBackup(encryptedBackup EncryptedBackupData, backupPassword string) error {
	// Validate encrypted backup
	if !encryptedBackup.Encrypted {
		return fmt.Errorf("backup is not encrypted")
	}

	if encryptedBackup.EncryptedPayload == "" {
		return fmt.Errorf("no encrypted data found")
	}

	// Decrypt the payload
	decryptedJSON, err := bc.encryptionService.DecryptBackupFromBase64(encryptedBackup.EncryptedPayload, backupPassword)
	if err != nil {
		// Return more specific error for wrong password
		return fmt.Errorf("wrong password: failed to decrypt backup data")
	}

	// Parse the decrypted JSON
	var backup BackupData
	if err := json.Unmarshal(decryptedJSON, &backup); err != nil {
		return fmt.Errorf("corrupted backup data: failed to parse decrypted content")
	}

	// Validate decrypted backup structure
	if backup.Version == "" {
		return fmt.Errorf("invalid backup format: missing version information")
	}

	if len(backup.Passwords) == 0 {
		return fmt.Errorf("no passwords found in backup file")
	}

	// Try to import passwords
	err = bc.importPasswords(backup.Passwords)
	if err != nil {
		return fmt.Errorf("import failed: %w", err)
	}

	return nil
}

// importPasswords is a helper method to import password entries
func (bc *BackupController) importPasswords(passwords []models.PasswordEntry) error {
	successCount := 0
	var lastError error

	for _, password := range passwords {
		// Create a new password entry (don't use existing ID to avoid conflicts)
		newEntry := models.PasswordEntry{
			ID:          "", // Empty ID so a new one will be generated
			Platform:    password.Platform,
			Email:       password.Email,
			Username:    password.Username,
			Password:    password.Password,
			Description: password.Description,
			// CreatedAt will be set automatically by the database
		}

		_, err := bc.dbService.SavePasswordEntry(newEntry)
		if err != nil {
			lastError = err
			fmt.Printf("Failed to import password for %s: %v\n", password.Platform, err)
			continue
		}
		successCount++
	}

	if successCount == 0 {
		return fmt.Errorf("failed to import any passwords: %v", lastError)
	}

	if lastError != nil {
		return fmt.Errorf("imported %d out of %d passwords, last error: %v",
			successCount, len(passwords), lastError)
	}

	return nil
}

// ValidateBackup validates backup file format and returns metadata (unencrypted)
func (bc *BackupController) ValidateBackup(backupJSON string) (BackupMetadata, error) {
	var backup BackupData
	if err := json.Unmarshal([]byte(backupJSON), &backup); err != nil {
		return BackupMetadata{}, fmt.Errorf("invalid JSON format: %w", err)
	}

	if backup.Version == "" {
		return BackupMetadata{}, fmt.Errorf("invalid backup format: missing version")
	}

	if backup.Passwords == nil {
		return BackupMetadata{}, fmt.Errorf("invalid backup format: missing passwords array")
	}

	metadata := BackupMetadata{
		PasswordCount: len(backup.Passwords),
		ExportDate:    backup.ExportDate,
		Version:       backup.Version,
		Encrypted:     false,
	}

	return metadata, nil
}

// ValidateEncryptedBackup validates encrypted backup file format and returns metadata
func (bc *BackupController) ValidateEncryptedBackup(backupJSON string) (BackupMetadata, error) {
	var encryptedBackup EncryptedBackupData
	if err := json.Unmarshal([]byte(backupJSON), &encryptedBackup); err != nil {
		return BackupMetadata{}, fmt.Errorf("invalid JSON format: %w", err)
	}

	if encryptedBackup.Version == "" {
		return BackupMetadata{}, fmt.Errorf("invalid backup format: missing version")
	}

	if !encryptedBackup.Encrypted || encryptedBackup.EncryptedPayload == "" {
		return BackupMetadata{}, fmt.Errorf("invalid encrypted backup format")
	}

	metadata := BackupMetadata{
		PasswordCount: -1, // Cannot determine without decryption
		ExportDate:    encryptedBackup.ExportDate,
		Version:       encryptedBackup.Version,
		Encrypted:     true,
	}

	return metadata, nil
}
