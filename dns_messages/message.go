package dns_messages

// This is a struct that represents a DNS message
// head, question, answers, authority, and additional sections are byte strings that
// are specified according to RFC 1035.
type Message struct {
	head       []byte
	question   []byte
	answers    []byte
	authority  []byte
	additional []byte
}

// Turns the DNS message into a byte slice for sending
func (msg *Message) ToByteSlice() []byte {
	all_components := [][]byte{msg.head, msg.question, msg.answers, msg.authority, msg.additional}
	rlt := []byte{}
	for _, elt := range all_components {
		rlt = append(rlt, elt...)
	}
	return rlt
}
