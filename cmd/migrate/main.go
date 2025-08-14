package main

import (
	"crypto/rand"
	"encoding/hex"
)

func generateRandomAddress() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func main() {

}
