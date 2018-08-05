// httplistener listens for http traffic on a port.
// It then writes the request data info a buffer.
package httplistener

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/textproto"
	"strings"
)

type Listener struct {
	ReqData chan []byte
	Proto   string
	Port    string
	list    net.Listener
}

func New(proto string, port string, size int) *Listener {
	return &Listener{
		Proto:   proto,
		Port:    port,
		ReqData: make(chan []byte, size),
	}
}

// HandleConnection handles the connection and dumps the request data into the channel.
func (l *Listener) HandleConnection(conn net.Conn) error {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	protoReader := textproto.NewReader(reader)
	reqData := []byte{}
	for {
		line, err := protoReader.ReadLineBytes()
		if err != nil {
			return err
		}
		reqData = append(reqData, line...)
		reqData = append(reqData, []byte("\n")...)
		// end of an http request.
		if len(line) == 0 {
			break
		}
	}

	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader("OK")),
	}

	err := resp.Write(writer)
	if err != nil {
		fmt.Println("Error writing", err)
	}

	l.ReqData <- reqData
	return nil
}

// Receive receives the data on a channel.
func (l *Listener) Receive() chan []byte {
	return l.ReqData
}

// Listen listens to a paticular protocol and port.
func (l *Listener) Listen() error {
	nl, err := net.Listen(l.Proto, l.Port)
	if err != nil {
		return err
	}

	// Save on the object so that it can be closed later.
	l.list = nl

	go func() {
		defer nl.Close()
		for {
			conn, err := nl.Accept()
			if err != nil {
				log.Println("Error while accepting connection.", err.Error())
			}

			err = l.HandleConnection(conn)
			if err != nil {
				log.Println("Error while multiplying req.", err.Error())
			}
		}
	}()
	return nil

}

// Close the listener and cleanup the things used.
func (l *Listener) Close() {
	l.list.Close()
}
