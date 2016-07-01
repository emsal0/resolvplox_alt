package main

import (
	"./dns_messages"
	//"./networking"
	"fmt"
)

func main() {
	fmt.Println(dns_messages.NametoQuery([]byte("www.google.com")))
}
