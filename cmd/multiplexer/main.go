package main

import (
	"fmt"

	"github.com/georgethomas111/reqmultiplier/httplistener"
)

func main() {
	listPort := ":8081"
	proto := "tcp"
	l := httplistener.New(proto, listPort, 10)
	err := l.Listen()
	if err != nil {
		fmt.Println("Error while iniitializing. ", err.Error())
		return
	}

	fmt.Printf("Port %s with protocol and address %s\n", listPort, proto)

	reqCh := l.Receive()

	reqData := <-reqCh
	fmt.Println("Got message ", string(reqData))
}
