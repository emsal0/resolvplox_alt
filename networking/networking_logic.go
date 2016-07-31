package networking

import (
	"github.com/emsal1863/resolvplox_alt/dns_messages"
	"net"
)

// Sends a DNS Message to the desired server (only port 53 supported right now)
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
