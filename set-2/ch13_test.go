package set2

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestChal13(t *testing.T) {
	kvResult := KvParser("foo=bar&baz=qux&zap=zazzle")
	kvExpected := map[string]string{
		"foo": "bar",
		"baz": "qux",
		"zap": "zazzle",
	}
	assert.Equal(t, kvResult, kvExpected)


	proResult := ProfileFor("foo@bar.com")
	proExpected := map[string]interface{}{
		"email": "foo@bar.com",
		"uid": 11,
		"role": "user",
	  }
	assert.Equal(t, proResult, proExpected)

	encodedResult := ProfileEncode(proResult)
	encodedExpected := "email=foo@bar.com&uid=11&role=user"
	assert.Equal(t, encodedResult, encodedExpected)
}

func TestBreakChal13(t *testing.T) {
	// email=foooo@bar., com&uid=11&role=, user
	// email=foooo@bar., admin\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b, com&uid=11&role=, user
	original := "foooo@bar.com"
	admin := "foooo@bar.admin\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0bcom"
	encOri, encAdmin := Chal13(original), Chal13(admin)
	ad := append(encOri[:32], encAdmin[16:32]...)
	decAdmin := BreakChal13(ad)

	fmt.Println("challenge 13 test:", decAdmin)
}