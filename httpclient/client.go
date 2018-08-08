package tcphttp

import (
	"errors"
	"net"
)

type Client struct {
	conn net.Conn
}

// Connect is used to make a connection to the server.
func (c *Client) Connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	c.conn = conn
	return nil
}

// Req makes the request to an http sever using th
func (c *Client) Req(httpCommand []byte) ([]byte, error) {
	n, err := c.conn.Write(httpCommand)
	if err != nil {
		return nil, err
	}

	if n != len(httpCommand) {
		return nil, errors.New("bytes written to tcp conn != len(httpCommand)")
	}

	var resp []byte
	_, err = c.conn.Read(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
