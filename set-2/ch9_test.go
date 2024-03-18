package set2

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestPKCS7Pad(t *testing.T){
	blockSize := 20
	text := "YELLOW SUBMARINE"
	result := PKCS7Pad([]byte(text), blockSize)
	expected := "YELLOW SUBMARINE\x04\x04\x04\x04"
	fmt.Printf("challenge 8 test: %q\n", result)
	assert.Equal(t, result, []byte(expected))
}

