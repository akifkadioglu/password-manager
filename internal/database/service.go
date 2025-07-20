package database

import (
	"fmt"
	"os"
	"path/filepath"

	"owllock/internal/config"
	"owllock/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Service handles all database operations
type Service struct {
	db *gorm.DB
}

// NewService creates a new database service
func NewService() *Service {
	return &Service{}
}

// Initialize initializes the SQLite database
func (s *Service) Initialize() error {
	// Get user's data directory from config
	userDataDir, err := config.GetDataDir()
	if err != nil {
		return fmt.Errorf("failed to get user data directory: %v", err)
	}

	// Create data directory if it doesn't exist
	if err := os.MkdirAll(userDataDir, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %v", err)
	}

	// Get database file name from config
	dbFileName := "passwords.db"
	if config.Config != nil && config.Config.Database.Name != "" {
		dbFileName = config.Config.Database.Name
	}

	// Database file path
	dbPath := filepath.Join(userDataDir, dbFileName)

	// Determine log level from config
	logLevel := logger.Silent
	if config.Config != nil && config.Config.Database.LogLevel == "info" {
		logLevel = logger.Info
	}

	// Open database
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	s.db = db

	// Auto-migrate the schema
	if err := s.db.AutoMigrate(&models.PasswordEntry{}); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	return nil
}

// SavePasswordEntry saves a password entry to database
func (s *Service) SavePasswordEntry(entry models.PasswordEntry) (models.PasswordEntry, error) {
	if s.db == nil {
		return entry, fmt.Errorf("database not initialized")
	}

	if err := s.db.Create(&entry).Error; err != nil {
		return entry, fmt.Errorf("failed to save password entry: %v", err)
	}

	return entry, nil
}

// GetPasswordEntries returns all saved password entries
func (s *Service) GetPasswordEntries() ([]models.PasswordEntry, error) {
	if s.db == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	var entries []models.PasswordEntry
	if err := s.db.Find(&entries).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve password entries: %v", err)
	}

	return entries, nil
}

// DeletePasswordEntry deletes a password entry by ID
func (s *Service) DeletePasswordEntry(id string) error {
	if s.db == nil {
		return fmt.Errorf("database not initialized")
	}

	if err := s.db.Delete(&models.PasswordEntry{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete password entry: %v", err)
	}

	return nil
}
