package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetSHA256(secret string) string {
	sha256Bytes := sha256.Sum256([]byte(secret))
	return hex.EncodeToString(sha256Bytes[:])
}
