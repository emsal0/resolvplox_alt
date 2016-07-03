package dns_messages

import (
	"math/rand"
	"time"
)

func NametoQuery(name []byte) (id []byte, msg Message) {
	id, header := generateQueryHeader()

	qlength := uint8(len(name))

	message := Message{
		head:       header,
		question:   append(append([]byte{}, qlength), name...),
		answers:    []byte{},
		authority:  []byte{},
		additional: []byte{},
	}

	return id, message
}

func generateQueryHeader() (id []byte, header []byte) {
	src := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(src)

	id = make([]byte, 2)
	gen.Read(id)

	row2 := []byte{0x81, 0x00}
	qdcount := []byte{0x01}
	ancount := []byte{0x00}
	nscount := []byte{0x00}
	arcount := []byte{0x00}

	header = []byte{}
	headerComponents := [][]byte{id, row2, qdcount, ancount, nscount, arcount}
	for _, elt := range headerComponents {
		header = append(header, elt...)
	}

	return id, header
}
