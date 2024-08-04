package tests

import (
	"testing"

	"github.com/GarudaProject/eyhash"
	"github.com/stretchr/testify/assert"
)

func TestMD5FileEqual(t *testing.T) {
	hash1, err := eyhash.MD5File("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	hash2, err := eyhash.MD5File("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, hash1, hash2)
}

func TestMD5FileNotEqual(t *testing.T) {
	hash1, err := eyhash.MD5File("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	hash2, err := eyhash.MD5File("test2.txt")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, hash1, hash2)
}
