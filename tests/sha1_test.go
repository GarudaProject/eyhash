package tests

import (
	"testing"

	"github.com/GarudaProject/eyhash"
	"github.com/stretchr/testify/assert"
)

func TestSHA1FileEqual(t *testing.T) {
	hash1, err := eyhash.SHA1File("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	hash2, err := eyhash.SHA1File("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, hash1, hash2)
}

func TestSHA1FileNotEqual(t *testing.T) {
	hash1, err := eyhash.SHA1File("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	hash2, err := eyhash.SHA1File("test2.txt")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, hash1, hash2)
}
