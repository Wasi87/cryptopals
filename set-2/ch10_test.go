package set2

import (
	"fmt"
	"testing"
)

func TestChal10(t *testing.T){
	key := "YELLOW SUBMARINE"
	result := Chal10(key, "texts/ch10.txt", byte(0))
	fmt.Printf("challenge 10 test: %s\n", result)
}

