package set1

import (
	"fmt"
	"testing"
)

func TestDetectSingleCharXor(t *testing.T) {
	result := DetectSingleCharXor("texts/ch4.txt")
	fmt.Println("challenge 4 test:", result[0].XordString)
}