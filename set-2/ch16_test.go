package set2

import (
	"testing"
)


func TestCheckAdmin(t *testing.T) {
	forgedContent := AddAdmin()
	check := CheckAdmin(forgedContent)
	if !check{
		t.Error("Expected CheckAdmin to return true")
	}
}