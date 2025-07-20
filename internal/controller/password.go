package controller

import (
	"crypto/md5"
	"fmt"
	"os"
	"runtime"
	"strings"

	"owllock/internal/config"
	"owllock/internal/database"
	"owllock/internal/models"
	"owllock/pkg/utils"
)

// PasswordController handles password-related operations
type PasswordController struct {
	dbService *database.Service
}

// NewPasswordController creates a new password controller
func NewPasswordController(dbService *database.Service) *PasswordController {
	return &PasswordController{
		dbService: dbService,
	}
}

// GetSystemFingerprint returns a unique fingerprint for the current system
func (pc *PasswordController) GetSystemFingerprint() string {
	// Collect system information
	hostname, _ := os.Hostname()
	username := os.Getenv("USER")
	if username == "" {
		username = os.Getenv("USERNAME") // Windows fallback
	}

	// Build fingerprint string
	fingerprintData := fmt.Sprintf("%s-%s-%s-%s",
		runtime.GOOS,
		runtime.GOARCH,
		hostname,
		username,
	)

	// Hash the fingerprint data
	hash := md5.Sum([]byte(fingerprintData))
	return fmt.Sprintf("%x", hash)
}

// GeneratePassword generates a secure password with given options
func (pc *PasswordController) GeneratePassword(options models.PasswordGenerationOptions) (string, error) {
	if options.Length <= 0 {
		return "", fmt.Errorf("password length must be greater than 0")
	}

	// Validate length against config limits
	if config.Config != nil {
		minLength := config.Config.Security.PasswordMinLength
		maxLength := config.Config.Security.PasswordMaxLength

		if options.Length < minLength {
			return "", fmt.Errorf("password length must be at least %d characters", minLength)
		}
		if options.Length > maxLength {
			return "", fmt.Errorf("password length cannot exceed %d characters", maxLength)
		}
	}

	// Character sets
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	symbols := "!@#$%^&*()_+-=[]{}|;:,.<>?"

	// Build character pool
	var charPool strings.Builder
	var guaranteedChars []string

	if options.IncludeUppercase {
		charPool.WriteString(uppercase)
		guaranteedChars = append(guaranteedChars, string(uppercase[utils.RandomInt(len(uppercase))]))
	}
	if options.IncludeLowercase {
		charPool.WriteString(lowercase)
		guaranteedChars = append(guaranteedChars, string(lowercase[utils.RandomInt(len(lowercase))]))
	}
	if options.IncludeNumbers {
		charPool.WriteString(numbers)
		guaranteedChars = append(guaranteedChars, string(numbers[utils.RandomInt(len(numbers))]))
	}
	if options.IncludeSymbols {
		charPool.WriteString(symbols)
		guaranteedChars = append(guaranteedChars, string(symbols[utils.RandomInt(len(symbols))]))
	}

	if charPool.Len() == 0 {
		return "", fmt.Errorf("at least one character type must be selected")
	}

	charset := charPool.String()

	// Add system fingerprint entropy if requested
	if options.UseFingerprint {
		fingerprint := pc.GetSystemFingerprint()
		// Mix fingerprint into charset for additional entropy
		charset = charset + fingerprint
	}

	// Generate password
	password := make([]byte, options.Length)

	// Ensure at least one character from each selected type
	for i, char := range guaranteedChars {
		if i < options.Length {
			password[i] = char[0]
		}
	}

	// Fill remaining positions
	for i := len(guaranteedChars); i < options.Length; i++ {
		password[i] = charset[utils.RandomInt(len(charset))]
	}

	// Shuffle the password
	for i := range password {
		j := utils.RandomInt(len(password))
		password[i], password[j] = password[j], password[i]
	}

	return string(password), nil
}

// SavePasswordEntry saves a password entry
func (pc *PasswordController) SavePasswordEntry(entry models.PasswordEntry) (models.PasswordEntry, error) {
	// Generate ID if not provided
	if entry.ID == "" {
		entry.ID = utils.GenerateID()
	}

	return pc.dbService.SavePasswordEntry(entry)
}

// GetPasswordEntries returns all saved password entries
func (pc *PasswordController) GetPasswordEntries() ([]models.PasswordEntry, error) {
	return pc.dbService.GetPasswordEntries()
}

// DeletePasswordEntry deletes a password entry by ID
func (pc *PasswordController) DeletePasswordEntry(id string) error {
	return pc.dbService.DeletePasswordEntry(id)
}
