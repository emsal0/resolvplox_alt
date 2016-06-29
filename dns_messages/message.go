package dns_messages

type Message struct {
	head       []byte
	question   []byte
	answers    []byte
	authority  []byte
	additional []byte
}

func (msg *Message) ToByteSlice() []byte {
	all_components := [][]byte{msg.head, msg.question, msg.answers, msg.authority, msg.additional}
	rlt := []byte{}
	for _, elt := range all_components {
		rlt = append(rlt, elt...)
	}
	return rlt
}
