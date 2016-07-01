package main

import (
	"./dns_messages"
	"./networking"
	"fmt"
)

func main() {
	_, query := dns_messages.NametoQuery([]byte("www.google.com"))
	fmt.Println("QUERY: ")
	fmt.Println(query)

	fmt.Println("CONVERTED QUERY: ")
	fmt.Println(query.ToByteSlice())

	response, _, _ := networking.SendQuery(query, "8.8.8.8")
	fmt.Println("RESPONSE: ")
	fmt.Println(response)
}
