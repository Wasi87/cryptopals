package set2

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestVerifyPadding(t *testing.T) {
	t.Run("Valid padding", func(t *testing.T) {
		result, err := VerifyPadding([]byte("ICE ICE BABY\x04\x04\x04\x04"), 16)
		if err != nil {
			t.Errorf("Expected valid padding, but got error: %v", err)
		}
		expected := "ICE ICE BABY"
		assert.Equal(t, string(result), expected)
	})

	t.Run("Invalid padding", func(t *testing.T) {
		_, err := VerifyPadding([]byte("ICE ICE BABY\x05\x05\x05\x05"), 16)
		if err == nil {
			t.Error("Expected invalid padding")
		} 
	
		_, err = VerifyPadding([]byte("ICE ICE BABY\x01\x02\x03\x04"), 16)
		if err == nil {
			t.Error("Expected invalid padding")
		}

		_, err = VerifyPadding([]byte("ICE ICE BABY\x01\x02\x03\x04"), 16)
		if err == nil {
			t.Error("Expected invalid padding")
		}
	})
}