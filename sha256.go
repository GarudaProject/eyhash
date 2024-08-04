package eyhash

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func SHA256File(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(file)
	return hex.EncodeToString(sum[:]), nil
}
