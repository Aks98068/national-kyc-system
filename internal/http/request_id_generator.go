package httpserver

import (
	"crypto/rand"
	"encoding/hex"
)

func generateRequestID() string {
	buffer := make([]byte, 16)

	if _, err := rand.Read(buffer); err != nil {
		return "unknown"
	}

	return hex.EncodeToString(buffer)
}
