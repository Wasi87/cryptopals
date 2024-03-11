package set1

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRepeatingKeyXor(t *testing.T){
	key := "ICE"
	input := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	result := RepeatingKeyXor(key, input)
	assert.Equal(t, result, expected)
	fmt.Println("challenge 5 test:", result)
}