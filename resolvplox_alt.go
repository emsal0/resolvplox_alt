package main

import (
	"./dns_messages"
	"fmt"
)

func main() {
	fmt.Println(dns_messages.NametoQuery([]byte("www.google.com")))
}
