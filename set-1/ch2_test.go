package set1

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestXor(t *testing.T){
	result, err := Xor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	if err != nil {
        t.Errorf("Error: %v", err)
        return
    }
	expected := "746865206b696420646f6e277420706c6179"
	fmt.Println("challenge 2 test:", result)
	assert.Equal(t, result, expected)
}