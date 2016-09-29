package main_test

import (
	"net"
	"testing"
)

// THIS TEST WILL ONLY PASS IF THE APP IS CURRENTLY RUNNING
func TestResolvploxServer(t *testing.T) {
	testQuery := []byte{
		0x04, 0x9a,
		0x01, 0x00,
		0x00, 0x01,
		0x00, 0x00,
		0x00, 0x00,
		0x00, 0x00,
		3, 'w', 'w', 'w',
		6, 'g', 'o', 'o', 'g', 'l', 'e',
		3, 'c', 'o', 'm',
		0x00,
		0x00, 0x01, // qtype
		0x00, 0x01, // qclass
	}

	remote_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:20841")
	local_addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:0")
	if err != nil {
		t.Error(err)
	}

	conn, err := net.DialUDP("udp", local_addr, remote_addr)
	if err != nil {
		t.Error(err)
	}
	response := make([]byte, 1024)
	conn.Write(testQuery)

	numBytes, _, err := conn.ReadFromUDP(response)
	if err != nil {
		t.Error(err)
	}
	response = response[:numBytes]

	t.Log(response)
}
