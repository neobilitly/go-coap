package udp

import "github.com/neobilitly/go-coap/v3/udp/server"

func NewServer(opt ...server.Option) *server.Server {
	return server.New(opt...)
}
