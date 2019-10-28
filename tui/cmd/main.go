package main

import (
	"chat/client"
	"chat/tui"
	"flag"
	"log"

)

func main() {
	address := flag.String("server", "localhost:3333", "Which server to connect to")

	flag.Parse()

	cl := client.NewClient()
	err := cl.Dial(*address)

	if err != nil {
		log.Fatal(err)
	}

	defer cl.Close()

	// start the client to listen for incoming message
	go cl.Start()

	tui.StartUi(cl)
}