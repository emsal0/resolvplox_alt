package main

import (
	"fmt"
	"github.com/emsal1863/resolvplox_alt/config_reader"
	"github.com/emsal1863/resolvplox_alt/dns_messages"
	"github.com/emsal1863/resolvplox_alt/networking"
	"net"
	"os"
)

func readUDP(conn *net.UDPConn, msg_chan chan []byte, addr_chan chan *net.UDPAddr) {
	for {
		bs := make([]byte, 1024)
		_, _, _, addr, _ := conn.ReadMsgUDP(bs, []byte{})
		msg_chan <- bs
		addr_chan <- addr
	}
}

func getResponse(conn *net.UDPConn, config config_reader.Config, msg_chan chan []byte, addr_chan chan *net.UDPAddr) {
	for {
		byteMsg := <-msg_chan
		addr := <-addr_chan

		msg, err := dns_messages.FromByteSlice(byteMsg)
		if err != nil {
			continue
		}
		name, _, err := msg.ExtractNameFromQuery()

		if err != nil {
			continue
		}

		response, _, _ := networking.SendQuery(msg, config.GetServer(string(name)))
		conn.WriteTo(response, addr)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: resolvplox {CONFIG_FILE}\n")
		os.Exit(1)
	}
	local_addr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:20841")
	conn, _ := net.ListenUDP("udp", local_addr) //TODO -- add error handling
	msg_chan := make(chan []byte)
	addr_chan := make(chan *net.UDPAddr)
	config, err := config_reader.FromFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading config file: "+err.Error())
		os.Exit(1)
	}

	go readUDP(conn, msg_chan, addr_chan)
	go getResponse(conn, config, msg_chan, addr_chan)

	fmt.Scanln()
}
