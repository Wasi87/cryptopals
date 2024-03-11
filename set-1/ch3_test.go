package set1

import (
	"fmt"
	"testing"
)

func TestGetTopFiveScore(t *testing.T) {
	result := GetTopFiveScore("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Println("challenge 3 test:", result[0].XordString)
}
