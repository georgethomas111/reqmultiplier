package httplistener

import (
	"bytes"
	"net/http"
	"sync"
	"testing"
)

func TestListenHTTPGet(t *testing.T) {
	listener := New("tcp", ":33333", 10)
	err := listener.Listen()
	if err != nil {
		t.Errorf("Error listening %s", err.Error())
		return
	}

	reqCh := listener.Receive()

	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		// Make a mock request to the server with the http client.
		http.Get("http://0.0.0.0:33333")

	}()

	msg := <-reqCh
	t.Log("Got message", string(msg))
	wg.Wait()
}

func TestListenHTTPPost(t *testing.T) {
	listener := New("tcp", ":33335", 10)
	err := listener.Listen()
	if err != nil {
		t.Errorf("Error listening %s", err.Error())
		return
	}

	reqCh := listener.Receive()
	go func() {
		// Make a mock request to the server with the http client.
		buf := bytes.NewBuffer([]byte("abcdefgh"))
		http.Post("http://0.0.0.0:33335", "text", buf)

	}()

	msg := <-reqCh
	t.Log(string(msg))
}
