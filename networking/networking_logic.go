package networking

import (
	"../dns_messages" // uncomment when ready
	//"fmt"
	"net"
)

func SendQuery(msg dns_messages.Message, dns_server string) (response []byte, numBytes int, err error) {
	serverAddr, err := net.ResolveUDPAddr("udp", dns_server+":53")
	if err != nil {
		return nil, -1, err
	}

	localAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:0")
	if err != nil {
		return nil, -1, err
	}

	conn, err := net.DialUDP("udp", localAddr, serverAddr)
	if err != nil {
		return nil, -1, err
	}

	response = make([]byte, 1024)
	conn.Write(msg.ToByteSlice())
	numBytes, _, err = conn.ReadFromUDP(response)

	err = nil
	conn.Close()

	return response, numBytes, err
}
