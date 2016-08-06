package main

import (
	"fmt"
	"github.com/emsal1863/resolvplox_alt/config_reader"
	"github.com/emsal1863/resolvplox_alt/networking"
	"net"
)

func readUDP(conn *net.UDPConn, msg_chan chan []byte, addr_chan chan *net.UDPAddr) {
	for {
		bs := make([]byte, 1024)
		_, _, _, addr, _ := conn.ReadMsgUDP(bs, []byte{})
		msg_chan <- bs
		addr_chan <- addr
	}
}

func getResponse(conn *net.UDPConn, msg_chan chan []byte, addr_chan chan *net.UDPAddr) {
	for {
		msg := <-msg_chan
		addr := <-addr_chan
		response, _, _ := networking.SendQueryByteSlice(msg, "8.8.8.8")
		conn.WriteTo(response, addr)
	}
}

func main() {
	x := config_reader.Config{}
	fmt.Println(x)
	local_addr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:20841")
	conn, _ := net.ListenUDP("udp", local_addr) //TODO -- add error handling
	msg_chan := make(chan []byte)
	addr_chan := make(chan *net.UDPAddr)

	go readUDP(conn, msg_chan, addr_chan)
	go getResponse(conn, msg_chan, addr_chan)

	fmt.Scanln()
}
