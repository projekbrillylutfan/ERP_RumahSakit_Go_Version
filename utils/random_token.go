package utils

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func GenerateToken(byteSize int) string {
	tokenBytes := make([]byte, byteSize)
	if _, err := rand.Read(tokenBytes); err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}
	return hex.EncodeToString(tokenBytes)
}