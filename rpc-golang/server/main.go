package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type API int

type Item struct {
	Title string
	Body  string
}

var database []Item

func (a *API) GetByName(Title string, reply *Item) error {
	for i := 0; i < len(database); i++ {
		if database[i].Title == Title {
			*reply = database[i]
			break
		}
	}
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	for i, v := range database {
		if v.Title == edit.Title {
			database[i] = Item{edit.Title, edit.Body}
			*reply = edit
			break
		}
	}
	return nil
}

func (a *API) DeleteItem(item Item, repl *Item) error {
	var del Item
	for i, v := range database {
		if item.Title == v.Title && item.Body == v.Body {
			database = append(database[:i], database[i+1:]...)
			del = item
			break
		}
	}
	*repl = del
	return nil
}

func (a *API) GetDB(title string, reply *[]Item) error {
	*reply = database
	return nil
}

func main() {
	port := flag.String("port", ":4040", "set port for server")
	flag.Parse()
	var api = new(API)

	err := rpc.Register(api)
	if err != nil {
		log.Fatalf("error registering API %v", err)
	}

	rpc.HandleHTTP()

	list, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("API listening error %v", err)
	}
	log.Printf("serving rpc on port %v", *port)
	{
		err := http.Serve(list, nil)
		if err != nil {
			log.Fatalf("error serving: %v", err)
		}
	}

	// fmt.Println("init db: ", database)

	// a := Item{"one", "shit one"}
	// b := Item{"two", "shit two"}
	// c := Item{"three", "shit three"}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("second database: ", database)

	// DeleteItem(b)
	// fmt.Println("deleting", database)

	// EditItem("three", Item{"shit", "bullshit"})
	// fmt.Println("editing", database)

	// x := GetByName("shit")
	// y := GetByName("one")
	// fmt.Println(x, y)
}
