package set2

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestChal9(t *testing.T){
	blockSize := 20
	text := "YELLOW SUBMARINE"
	result := Chal9([]byte(text), blockSize)
	expected := "YELLOW SUBMARINE\x04\x04\x04\x04"
	fmt.Printf("challenge 9 test: %q\n", result)
	assert.Equal(t, result, []byte(expected))
}

