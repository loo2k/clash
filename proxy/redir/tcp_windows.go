package redir

import (
	"errors"
	"net"

	"github.com/loo2k/go-shadowsocks2/socks"
)

func parserPacket(conn net.Conn) (socks.Addr, error) {
	return nil, errors.New("Windows not support yet")
}
