package networking

import (
	"errors"
	"github.com/emsal1863/resolvplox_alt/dns_messages"
	"net"
	"strconv"
	"strings"
)

func parseIP(ipStr string) (ip net.IP, err error) {
	splitIpStr := strings.Split(ipStr, ".")
	standardErrorMessage := "not a valid IP (v4 only supported at this point) string: "
	if len(splitIpStr) < 4 {
		return net.IPv4(0, 0, 0, 0), errors.New(standardErrorMessage + ipStr)
	}
	splitIpStrBytes := []byte{}

	for _, v := range splitIpStr {
		i, err := strconv.Atoi(v)
		if err != nil {
			return net.IPv4(0, 0, 0, 0), errors.New(standardErrorMessage + ipStr + " -- can't convert string")
		}

		if !(0 <= i && i <= 255) {
			return net.IPv4(0, 0, 0, 0), errors.New(standardErrorMessage + ipStr + " -- int not in range")
		}

		splitIpStrBytes = append(splitIpStrBytes, byte(i))
	}

	ip = net.IPv4(splitIpStrBytes[0], splitIpStrBytes[1], splitIpStrBytes[2], splitIpStrBytes[3])
	err = nil
	return
}

// Sends a DNS Message to the desired server (only port 53 supported right now)
func SendQuery(msg dns_messages.Message, dns_server string) (response []byte, numBytes int, err error) {
	ip, err := parseIP(dns_server)
	if err != nil {
		return nil, -1, err
	}

	serverAddr := &net.UDPAddr{
		IP:   ip,
		Port: 53,
		Zone: "",
	}

	localAddr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 0,
		Zone: "",
	}

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

	return
}

// Sends a DNS Message in byte slice form to the desired server (only port 53 supported right now)
func SendQueryByteSlice(byteMsg []byte, dns_server string) (response []byte, numBytes int, err error) {
	msg, err := dns_messages.FromByteSlice(byteMsg)
	if err != nil {
		return nil, -1, err
	}

	response, numBytes, err = SendQuery(msg, dns_server)
	return
}
