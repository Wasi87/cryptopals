package set1

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHexToBase64(t *testing.T){
	result := HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	fmt.Println("challenge 1 test:", result)
	assert.Equal(t, result, expected)
}