package eyhash

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

func MD5File(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := md5.Sum(file)
	return hex.EncodeToString(sum[:]), nil
}
