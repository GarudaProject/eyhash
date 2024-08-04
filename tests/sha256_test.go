package tests

import (
	"testing"

	"github.com/GarudaProject/eyhash"
	"github.com/stretchr/testify/assert"
)

func TestSHA256FileEqual(t *testing.T) {
	hash1, err := eyhash.SHA256File("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	hash2, err := eyhash.SHA256File("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, hash1, hash2)
}

func TestSHA256FileNotEqual(t *testing.T) {
	hash1, err := eyhash.SHA256File("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	hash2, err := eyhash.SHA256File("test2.txt")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, hash1, hash2)
}
