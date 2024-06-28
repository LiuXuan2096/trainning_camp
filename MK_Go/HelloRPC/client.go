package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dialing:", err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "葛诗颖", &reply)
	if err != nil {
		log.Fatal("Call error:", err)
	}

	fmt.Println(reply)
}
