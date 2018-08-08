package multiconn

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
	"os"
)

type MultiConn struct {
	destURL []*url.URL
}

// New
func New(proto string, destAddres []string) (*MultiConn, error) {
	var urls []*url.URL
	for _, addr := range destAddres {
		u, err := url.ParseRequestURI(addr)
		if err != nil {
			return nil, errors.New("parsing error - " + err.Error())
		}
		urls = append(urls, u)
	}

	m := &MultiConn{
		destURL: urls,
	}

	return m, nil
}

// HandleConnection
func (m *MultiConn) HandleConnection(conn net.Conn) error {
	// Close the connection what ever be the case.
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	//	req, err := http.ReadRequest(reader)
	//	if err != nil {
	//		return errors.New("HandleConnection Http - " + err.Error())
	//	}
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, string(d))

	//	for _, u := range m.destURL {
	// update the url to be the destination URL.
	//	req.URL = u
	//	fmt.Println("Request URI", req.RequestURI)
	//	req.RequestURI = ""
	//	resp, err := http.DefaultClient.Do(req)
	//	if err != nil {
	//		return errors.New("Client do http - " + err.Error())
	//	}

	fmt.Fprintf(writer, "abcd")
	//	if err != nil {
	//		return errors.New("Writer error - " + err.Error())
	//	}
	//	}

	return nil
}
