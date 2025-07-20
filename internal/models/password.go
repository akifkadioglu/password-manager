package models

import "time"

// PasswordEntry represents a saved password entry
type PasswordEntry struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Platform    string    `json:"platform" gorm:"not null"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Description string    `json:"description"`
	Password    string    `json:"password" gorm:"not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// PasswordGenerationOptions represents password generation options
type PasswordGenerationOptions struct {
	Length           int  `json:"length"`
	IncludeUppercase bool `json:"includeUppercase"`
	IncludeLowercase bool `json:"includeLowercase"`
	IncludeNumbers   bool `json:"includeNumbers"`
	IncludeSymbols   bool `json:"includeSymbols"`
	UseFingerprint   bool `json:"useFingerprint"`
}
