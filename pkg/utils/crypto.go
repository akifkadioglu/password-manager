package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// GenerateID creates a unique ID for password entries
func GenerateID() string {
	timestamp := GetCurrentTimestamp()
	randomBytes := make([]byte, 4)
	rand.Read(randomBytes)
	return fmt.Sprintf("%s-%x", timestamp, randomBytes)
}

// GetCurrentTimestamp returns current timestamp as string
func GetCurrentTimestamp() string {
	return time.Now().Format("2006-01-02T15:04:05Z")
}

// RandomInt generates a cryptographically secure random integer
func RandomInt(max int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(n.Int64())
}
