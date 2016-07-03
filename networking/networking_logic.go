package networking

import (
	"../dns_messages" // uncomment when ready
	"net"
)

func SendQuery(msg dns_messages.Message, dns_server string) (response []byte, numBytes int, err interface{}) {
	serverAddr, err := net.ResolveUDPAddr("udp", dns_server)
	localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	conn, err := net.DialUDP("udp", localAddr, serverAddr)

	conn.Write(msg.ToByteSlice())
	response = []byte{}
	numBytes, _, err = conn.ReadFromUDP(response)

	err = nil
	conn.Close()

	return response, numBytes, err
}
