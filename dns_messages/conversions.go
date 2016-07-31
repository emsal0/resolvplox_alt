package dns_messages

import (
	"bytes"
	"math/rand"
	"time"
)

// Generate a query based on a domain name (e.g. "google.com")
func NametoQuery(name []byte) (id []byte, msg Message) {
	id, header := generateQueryHeader()

	message := Message{
		Head:       header,
		Question:   generateQuery(name),
		Answers:    []byte{},
		Authority:  []byte{},
		Additional: []byte{},
	}

	return id, message
}

func generateQueryHeader() (id []byte, header []byte) {
	src := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(src)

	id = make([]byte, 2)
	gen.Read(id)

	row2 := []byte{0x01, 0x00}
	qdcount := []byte{0x00, 0x01}
	ancount := []byte{0x00, 0x00}
	nscount := []byte{0x00, 0x00}
	arcount := []byte{0x00, 0x00}

	header = []byte{}
	headerComponents := [][]byte{id, row2, qdcount, ancount, nscount, arcount}
	for _, elt := range headerComponents {
		header = append(header, elt...)
	}

	return id, header
}

func generateQuery(name []byte) []byte {
	qname := generateQueryName(name)
	qtype := []byte{0x00, 0x01}
	qclass := []byte{0x00, 0x01}

	all := [][]byte{qname, []byte{0x00}, qtype, qclass}
	ret := []byte{}
	for _, elt := range all {
		ret = append(ret, elt...)
	}

	return ret
}

func generateQueryName(name []byte) []byte {
	splat := bytes.Split(name, []byte{'.'})
	ret := []byte{}

	for _, field := range splat {
		ret = append(ret, []byte{uint8(len(field))}...)
		ret = append(ret, field...)
	}

	return ret
}
