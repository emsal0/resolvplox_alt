package dns_messages

import (
	"math/rand"
	"time"
)

type Message struct {
	head       []byte
	question   []byte
	answers    []byte
	authority  []byte
	additional []byte
}

func NametoQuery(name []byte) ([]byte, Message) {
	src := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(src)

	var id []byte
	id = make([]byte, 4)
	gen.Read(id)

	qr := []byte{1}
	opcode := []byte{0, 0, 0, 0}
	aa := []byte{0}
	tc := []byte{0}
	rd := []byte{1}
	ra := []byte{0}
	z := []byte{0, 0, 0}
	rcode := []byte{0, 0, 0, 0}
	qdcount := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	ancount := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	nscount := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	arcount := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	header := []byte{}
	headerComponents := [][]byte{id, qr, opcode, aa, tc, rd, ra, z, rcode, qdcount, ancount, nscount, arcount}
	for _, elt := range headerComponents {
		header = append(header, elt...)
	}

	qlength := uint8(len(name))

	message := Message{
		head:       header,
		question:   append(append([]byte{}, qlength), name...),
		answers:    []byte{},
		authority:  []byte{},
		additional: []byte{}}

	return id, message
}
