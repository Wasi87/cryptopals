package set1

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHammingDistance(t *testing.T){
	str1 := []byte("this is a test")
	str2 := []byte("wokka wokka!!!")
	expected := 37
	result, _ := HammingDistance(str1, str2)
	assert.Equal(t, expected, result)
}

func TestBreakRepeatingXor(t *testing.T){
	key, decrypted := BreakRepeatingXor("texts/ch6.txt")
	fmt.Println("challenge 6 key:\n", string(key))
	fmt.Println("challenge 6 Content:\n", string(decrypted))
}