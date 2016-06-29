package networking

import (
	//"../dns_messages" // uncomment when ready
	"net"
)

func connect(dns_server string) {
	ServerAddr, err := net.ResolveUDPAddr("udp", dns_server)
	if err != nil {
		return nil, err
	}

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
