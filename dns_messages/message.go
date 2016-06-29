package dns_messages

type Message struct {
	head       []byte
	question   []byte
	answers    []byte
	authority  []byte
	additional []byte
}
