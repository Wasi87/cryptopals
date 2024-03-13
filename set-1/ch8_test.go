package set1

import (
	"fmt"
	"testing"
)

func TestDetectECB(t *testing.T) {
	filePath := "texts/ch8.txt"
	ansLine, ans := DetectECB(filePath)
	fmt.Printf("challenge 8 test: line %d repeated %d times.\n", ansLine, ans)
}