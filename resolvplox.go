package main

import (
	"fmt"
	"github.com/emsal1863/resolvplox_alt/dns_messages"
	"github.com/emsal1863/resolvplox_alt/networking"
)

func main() {
	id, query := dns_messages.NametoQuery([]byte("google.com"))
	fmt.Println("QUERY: ")
	fmt.Println(id, query)

	fmt.Println("CONVERTED QUERY: ")
	fmt.Printf("%x\n", query.ToByteSlice())

	response, _, err := networking.SendQuery(query, "8.8.8.8")
	fmt.Println("RESPONSE: ")
	fmt.Println(response)
	fmt.Println(err)
}
