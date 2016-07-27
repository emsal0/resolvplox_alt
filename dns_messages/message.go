package dns_messages

// The data structure for DNS messages.
// Structured after standard DNS messages as per RFC 1035.
type Message struct {
	head       []byte
	question   []byte
	answers    []byte
	authority  []byte
	additional []byte
}

// Convert a dns Message type to a byte slice for sending over UDP.
func (msg *Message) ToByteSlice() []byte {
	all_components := [][]byte{msg.head, msg.question, msg.answers, msg.authority, msg.additional}
	rlt := []byte{}
	for _, elt := range all_components {
		rlt = append(rlt, elt...)
	}
	return rlt
}
