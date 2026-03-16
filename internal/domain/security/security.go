package security

import (
	"crypto/sha256"
	"encoding/hex"
)

const Salt = "%D+u1zM0ZnD#tQ}Y5*+m+b7:cGasZmt}"

func GetSha256(input string) string {
	hash := sha256.Sum256([]byte(input + Salt))
	return hex.EncodeToString(hash[:])
}
