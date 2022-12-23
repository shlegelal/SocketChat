package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type server struct {
	members  map[net.Addr]*client
	messages chan message
}

func newServer() *server {
	return &server{
		make(map[net.Addr]*client),
		make(chan message),
	}
}

func Start(port int) {
	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Printf("faild to listen port %d, err=%v", port, err)
	}
	defer listener.Close()
	log.Printf("server started on :%d", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection, err=%v", err)
			continue
		}

		go s.newClient(conn)
	}

}

func (s *server) newClient(conn net.Conn) {

	_, err := conn.Write([]byte("> Enter your nickname: "))
	if err != nil {
		log.Printf("%v", err)
	}
	in := bufio.NewReader(conn)
	text, _, _ := in.ReadLine()

	nick := string(text)

	if nick == "" {
		nick = "anon"
	}

	c := client{
		conn,
		conn.RemoteAddr(),
		nick,
		s.messages,
	}

	c.msg("> To exit enter Ctrl+C")

	s.members[conn.RemoteAddr()] = &c

	msg := fmt.Sprintf("[%s] connected", nick)
	s.messages <- message{
		JOIN,
		"",
		conn.RemoteAddr(),
		msg,
	}

	go c.readInput()
}

func (s *server) run() {
	for msg := range s.messages {
		if msg.Type == QUIT {
			delete(s.members, msg.From)
		}
		log.Printf("%s, %s", msg.string(), msg.From)
		for addr, c := range s.members {
			if msg.From != addr {
				c.msg(msg.string())
			}
		}
	}
}
