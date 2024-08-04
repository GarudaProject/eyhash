package eyhash

import (
	"crypto/sha512"
	"encoding/hex"
	"os"
)

func SHA512File(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha512.Sum512(file)
	return hex.EncodeToString(sum[:]), nil
}
