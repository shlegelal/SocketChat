package client

import (
	"fmt"
	"github.com/reiver/go-telnet"
	"log"
)

func Start(port int) {
	err := telnet.DialToAndCall(fmt.Sprintf(":%d", port), telnet.StandardCaller)
	if err != nil {
		log.Printf("failed to connect port %d, err=%v", port, err)
		return
	}
	log.Printf("server failed")
}
