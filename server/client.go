package server

import (
	"bufio"
	"fmt"
	"net"
)

type client struct {
	conn     net.Conn
	addr     net.Addr
	nick     string
	messages chan message
}

func (c *client) readInput() {
	for {
		in := bufio.NewReader(c.conn)
		text, _, err := in.ReadLine()
		if err != nil {
			c.messages <- message{
				QUIT,
				"",
				c.addr,
				fmt.Sprintf("[%s] left", c.nick),
			}
			return
		}
		msg := string(text)

		c.messages <- message{
			MSG,
			c.nick,
			c.conn.RemoteAddr(),
			msg,
		}
	}
}

func (c *client) msg(msg string) {
	_, err := c.conn.Write([]byte(msg + "\n"))
	if err != nil {
		return
	}
}
