package tcphttp

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got something")
		fmt.Fprintln(w, "OK")
	}))
	defer ts.Close()

	client := &Client{}
	tcpURL := strings.Split(ts.URL, "//")[1]
	t.Log("Url :=", tcpURL)
	err := client.Connect(tcpURL)
	if err != nil {
		t.Errorf("Connect error: %s", err.Error())
	}
	defer client.Close()

	getReq := []byte(`GET / HTTP/1.1
			Host: ` + tcpURL + `
			User-Agent: Go-http-client/1.1
			Accept-Encoding: gzip


				`)

	t.Log("req:=", string(getReq))
	resp, err := client.Req(getReq)
	if err != nil {
		t.Errorf("Req error: %s", err.Error())
	}

	if string(resp) != "OK" {
		t.Errorf("Expected response to be OK not %s ", resp)
	}
}
