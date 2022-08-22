package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := Item{"one", "shirt one"}
	b := Item{"two", "shirt two"}
	// c := Item{"three", "shirt three"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	// client.Call("API.AddItem", c, &reply)

	client.Call("API.GetDB", "", &db)
	fmt.Println("DB from ", db)
}
