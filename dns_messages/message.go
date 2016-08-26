package dns_messages

import (
	"errors"
	"strings"
)

// The data structure for DNS messages.
// Structured after standard DNS messages as per RFC 1035.
type Message struct {
	Head       []byte
	Questions  []Question
	Answers    []byte
	Authority  []byte
	Additional []byte
}

type Question struct {
	Name  []byte
	Type  []byte
	Class []byte
}

func (q *Question) ToByteSlice() []byte {
	return append(append(q.Name, q.Type...), q.Class...)
}

func NameToQuestion(name []byte) Question {
	segments := strings.Split(string(name), ".")
	questionBody := []byte{}
	for _, segment := range segments {
		segmentString := append([]byte{byte(len(segment))}, []byte(segment)...)
		questionBody = append(questionBody, segmentString...)
	}
	questionBody = append(questionBody, 0x00)

	return Question{
		Name:  questionBody,
		Type:  []byte{0x00, 0x01},
		Class: []byte{0x00, 0x01},
	}
}

// Convert a dns Message type to a byte slice for sending over UDP.
func (msg *Message) ToByteSlice() []byte {
	questionsSlice := []byte{}
	for _, elt := range msg.Questions {
		questionsSlice = append(questionsSlice, elt.ToByteSlice()...)
	}
	all_components := [][]byte{msg.Head, questionsSlice, msg.Answers, msg.Authority, msg.Additional}
	rlt := []byte{}
	for _, elt := range all_components {
		rlt = append(rlt, elt...)
	}
	return rlt
}

func FromByteSlice(byteQuery []byte) (msg Message, err error) {
	headLength := 12
	if len(byteQuery) < headLength {
		return Message{}, errors.New("Not a possible query: length too short")
	}
	head := byteQuery[:headLength]
	QDCOUNT := int(head[4])<<8 + int(head[5])
	// TODO -- support this
	//ANCOUNT := int(head[6])<<8 + int(head[7])
	//NSCOUNT := int(head[8])<<8 + int(head[9])
	//ARCOUNT := int(head[10])<<8 + int(head[11])

	curPos := headLength
	questions := []Question{}
	for i := 0; i < QDCOUNT; i++ {
		curQ := Question{}
		for byteQuery[curPos] != 0x00 {
			segmentLength := byteQuery[curPos]
			curQ.Name = byteQuery[curPos+1 : segmentLength]

			curPos += (int(segmentLength) + 1)
		}
		curPos++
		curQ.Type = byteQuery[curPos : curPos+2]
		curQ.Class = byteQuery[curPos+2 : curPos+4]

		curPos += 4
		questions = append(questions, curQ)
	}

	return Message{
		Head:      head,
		Questions: questions,
	}, nil
}
