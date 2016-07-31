package dns_messages_test

import (
	"github.com/emsal1863/resolvplox_alt/dns_messages"
	"testing"
)

func testEq(a, b []byte) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestToByteSlice(t *testing.T) {
	msg := dns_messages.Message{
		Head:       []byte{'a'},
		Question:   []byte{'b'},
		Answers:    []byte{'c'},
		Authority:  []byte{'d'},
		Additional: []byte{'e'}}

	msgConverted := msg.ToByteSlice()

	if !testEq(msgConverted, []byte{'a', 'b', 'c', 'd', 'e'}) {
		t.Error("ToByteSlice didn't append all fields in message struct together. Converted: " + string(msgConverted) + ", Expected: abcde")
	}
}
