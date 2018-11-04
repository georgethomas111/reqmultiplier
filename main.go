package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	url      = "http://google.com"
	proto    = "GET"
	reqBody  = ""
	multiply = 10
	timeout  = 0
)

func initFlags() {
	flag.StringVar(&url, "destURL", url, "URL to hit.")
	flag.StringVar(&proto, "proto", proto, "Protocol to use.")
	flag.StringVar(&reqBody, "reqBody", reqBody, "Request body to send.")
	flag.IntVar(&multiply, "multiply", multiply, "Number of times to multiply request.")
	flag.IntVar(&timeout, "timeout", timeout, "timeout between requests.")
	flag.Parse()
}

func main() {
	initFlags()

	body := bytes.NewBufferString(reqBody)
	client := http.Client{}
	req, err := http.NewRequest(proto, url, body)
	if err != nil {
		log.Println("Error while creating request ", err.Error())
	}

	for i := 0; i < multiply; i++ {
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error making req ", err.Error())
			continue
		}

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response body", err.Error())
		}

		defer resp.Body.Close()

		log.Printf("Req No %d \n\n resp = %s\n", i, string(respBody))
		time.Sleep(time.Duration(timeout) * time.Millisecond)
	}
}
