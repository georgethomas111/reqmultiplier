package multiconn

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

type MultiConn struct {
	destAddres []string
	destConns  []net.Conn
}

// New
// dial to the address and save the address
func New(proto string, destAddres []string) *MultiConn {
	var conns []net.Conn
	for _, addr := range destAddres {
		c, err := net.Dial(proto, addr)
		if err != nil {
			log.Println("Connecton error to dest", err.Error())
			continue
		}

		conns = append(conns, c)
	}

	return &MultiConn{
		destAddres: destAddres,
		destConns:  conns,
	}
}

func (m *MultiConn) SendConn(conn net.Conn) error {
	data, err := ioutil.ReadAll(conn)
	if err != nil {
		return err
	}

	for _, destConn := range m.destConns {
		fmt.Fprintf(destConn, "%s", data)
	}

	return nil
}
