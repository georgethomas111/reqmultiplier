package main

import (
	"fmt"
	"log"
	"net"

	"github.com/georgethomas111/reqmultiplier/multiconn"
)

func ListenMultiplexer(listPort string, proto string, multi *multiconn.MultiConn) error {
	l, err := net.Listen(proto, listPort)
	if err != nil {
		return err
	}
	// Listener should be closed eventually.
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error while accepting connection.", err.Error())
		}

		log.Println("Got the request connection, now should multiply it.")
		err = multi.SendConn(conn)
		if err != nil {
			log.Println("Error while multiplying req.", err.Error())
		}
	}
}

func main() {
	listPort := ":8081"
	proto := "tcp"
	destAddrs := []string{":8080", ":8080"}
	multServ := multiconn.New(proto, destAddrs)

	fmt.Printf("Port %s with protocol and address %s\n", listPort, proto)

	ListenMultiplexer(listPort, proto, multServ)
	exitChan := make(chan int)
	go func() {
		<-exitChan
	}()
}
