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
		Head: []byte{'a'},
		Questions: []dns_messages.Question{dns_messages.Question{
			Name:  []byte{'b'},
			Type:  []byte{0x00},
			Class: []byte{0x00},
		}},
		Answers:    []byte{'c'},
		Authority:  []byte{'d'},
		Additional: []byte{'e'}}

	msgConverted := msg.ToByteSlice()

	if !testEq(msgConverted, []byte{'a', 'b', 0x00, 0x00, 'c', 'd', 'e'}) {
		t.Error("ToByteSlice didn't append all fields in message struct together. Converted: " + string(msgConverted) + ", Expected: abcde")
	}
}

func TestNametoQuery(t *testing.T) {
	id, msg := dns_messages.NametoQuery([]byte("www.google.com"))
	expectedQueryAfterId := []byte{
		0x01, 0x00,
		0x00, 0x01,
		0x00, 0x00,
		0x00, 0x00,
		0x00, 0x00,
		3, 'w', 'w', 'w',
		6, 'g', 'o', 'o', 'g', 'l', 'e',
		3, 'c', 'o', 'm',
		0x00,
		0x00, 0x01, // qtype
		0x00, 0x01, // qclass
	}
	expectedQuery := append(id, expectedQueryAfterId...)
	actualQuery := msg.ToByteSlice()
	if !testEq(actualQuery, expectedQuery) {
		t.Log(actualQuery)
		t.Error("NameToQuery didn't construct expected query for www.google.com")
	}
}

func TestFromByteSlice(t *testing.T) {
	byteQuery := []byte{
		0xfa, 0xcb,
		0x01, 0x00,
		0x00, 0x01,
		0x00, 0x00,
		0x00, 0x00,
		0x00, 0x00,
		3, 'w', 'w', 'w',
		6, 'g', 'o', 'o', 'g', 'l', 'e',
		3, 'c', 'o', 'm',
		0x00,
		0x00, 0x01, // qtype
		0x00, 0x01, // qclass
	}
	msg, err := dns_messages.FromByteSlice(byteQuery)
	if err != nil {
		t.Error("Error parsing msg from bytes: " + err.Error())
	}
	t.Log(msg)
}

func TestExtractName(t *testing.T) {
	_, msg := dns_messages.NametoQuery([]byte("www.google.com"))
	actual, _, _ := msg.ExtractNameFromQuery()
	if !testEq(actual, []byte("www.google.com")) {
		t.Log(string(actual))
		t.Error("should have extracted google.com")
	}
}

func TestExtractNameFromByteSlice(t *testing.T) {
	byteQuery := []byte{
		0xfa, 0xcb,
		0x01, 0x00,
		0x00, 0x01,
		0x00, 0x00,
		0x00, 0x00,
		0x00, 0x00,
		3, 'w', 'w', 'w',
		6, 'g', 'o', 'o', 'g', 'l', 'e',
		3, 'c', 'o', 'm',
		0x00,
		0x00, 0x01, // qtype
		0x00, 0x01, // qclass
	}
	msg, err := dns_messages.FromByteSlice(byteQuery)
	if err != nil {
		t.Error("Error parsing msg from bytes: " + err.Error())
	}

	actual, _, _ := msg.ExtractNameFromQuery()
	if !testEq(actual, []byte("www.google.com")) {
		t.Log(string(actual))
		t.Error("should have extracted google.com")
	}
}
