package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		cn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleconnection(cn)
	}
}

func handleconnection(cn net.Conn) {
	defer cn.Close()
	for {
		_, err := io.WriteString(cn, time.Now().Format(time.RFC850+"\n"))
		if err != nil {
			return
		}
		time.Sleep(3 * time.Second)
	}
}
