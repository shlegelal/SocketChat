package main

import (
	"flag"
	"log"
	"socket_chat/client"
	"socket_chat/server"
)

func main() {
	service := flag.String("service", "client", "")
	port := flag.Int("port", 8080, "")
	flag.Parse()

	switch *service {
	case "client":
		client.Start(*port)
	case "server":
		server.Start(*port)
	default:
		log.Fatalf("service expected client or server")
	}
}
