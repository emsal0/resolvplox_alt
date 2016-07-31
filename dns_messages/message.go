package dns_messages

// The data structure for DNS messages.
// Structured after standard DNS messages as per RFC 1035.
type Message struct {
	Head       []byte
	Question   []byte
	Answers    []byte
	Authority  []byte
	Additional []byte
}

// Convert a dns Message type to a byte slice for sending over UDP.
func (msg *Message) ToByteSlice() []byte {
	all_components := [][]byte{msg.Head, msg.Question, msg.Answers, msg.Authority, msg.Additional}
	rlt := []byte{}
	for _, elt := range all_components {
		rlt = append(rlt, elt...)
	}
	return rlt
}
