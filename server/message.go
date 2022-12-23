package server

import (
	"fmt"
	"net"
)

type messageType int

const (
	MSG messageType = iota
	JOIN
	QUIT
)

type message struct {
	Type   messageType
	Author string
	From   net.Addr
	Text   string
}

func (m *message) string() string {
	switch m.Type {
	case MSG:
		return fmt.Sprintf("[%s] %s", m.Author, m.Text)
	default:
		return fmt.Sprintf("> %s", m.Text)
	}
}
