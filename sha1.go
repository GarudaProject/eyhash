package eyhash

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
)

func SHA1File(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha1.Sum(file)
	return hex.EncodeToString(sum[:]), nil
}
