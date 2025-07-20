package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

// EncryptionService handles backup encryption/decryption
type EncryptionService struct{}

// NewEncryptionService creates a new encryption service
func NewEncryptionService() *EncryptionService {
	return &EncryptionService{}
}

// deriveKey derives a key from password using PBKDF2
func (es *EncryptionService) deriveKey(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)
}

// EncryptBackup encrypts backup data with a password
func (es *EncryptionService) EncryptBackup(data []byte, password string) ([]byte, error) {
	// Generate a random salt
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}

	// Derive key from password
	key := es.deriveKey(password, salt)

	// Create cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	// Generate IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("failed to generate IV: %w", err)
	}

	// Encrypt data
	stream := cipher.NewCFBEncrypter(block, iv)
	encrypted := make([]byte, len(data))
	stream.XORKeyStream(encrypted, data)

	// Combine salt + iv + encrypted data
	result := make([]byte, 0, len(salt)+len(iv)+len(encrypted))
	result = append(result, salt...)
	result = append(result, iv...)
	result = append(result, encrypted...)

	return result, nil
}

// DecryptBackup decrypts backup data with a password
func (es *EncryptionService) DecryptBackup(encryptedData []byte, password string) ([]byte, error) {
	if len(encryptedData) < 32 { // 16 salt + 16 IV minimum
		return nil, fmt.Errorf("encrypted data too short")
	}

	// Extract salt and IV
	salt := encryptedData[:16]
	iv := encryptedData[16:32]
	encrypted := encryptedData[32:]

	// Derive key from password
	key := es.deriveKey(password, salt)

	// Create cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	// Decrypt data
	stream := cipher.NewCFBDecrypter(block, iv)
	decrypted := make([]byte, len(encrypted))
	stream.XORKeyStream(decrypted, encrypted)

	return decrypted, nil
}

// EncryptBackupToBase64 encrypts and encodes to base64
func (es *EncryptionService) EncryptBackupToBase64(data []byte, password string) (string, error) {
	encrypted, err := es.EncryptBackup(data, password)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// DecryptBackupFromBase64 decodes from base64 and decrypts
func (es *EncryptionService) DecryptBackupFromBase64(encryptedBase64, password string) ([]byte, error) {
	encrypted, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %w", err)
	}
	return es.DecryptBackup(encrypted, password)
}

// GenerateSecurePassword generates a secure backup password
func (es *EncryptionService) GenerateSecurePassword(length int) (string, error) {
	if length < 8 {
		length = 16
	}

	// Character sets
	const (
		lowercase = "abcdefghijklmnopqrstuvwxyz"
		uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numbers   = "0123456789"
		symbols   = "!@#$%^&*()_+-=[]{}|;:,.<>?"
	)

	charset := lowercase + uppercase + numbers + symbols

	password := make([]byte, length)
	for i := range password {
		randomByte := make([]byte, 1)
		if _, err := rand.Read(randomByte); err != nil {
			return "", fmt.Errorf("failed to generate random byte: %w", err)
		}
		password[i] = charset[int(randomByte[0])%len(charset)]
	}

	return string(password), nil
}
