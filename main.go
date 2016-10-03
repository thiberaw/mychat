package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "listen on the specified ip address")
	flag.Parse()

	if isHost {
		connIp := os.Args[2]
		runHost(connIp)
	} else {
		connIp := os.Args[1]
		runGuest(connIp)
	}
}

const port = "8080"

func runHost(ip string) {
	ipAndPort := ip + ":" + port

	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error: ", listenErr)
	}

	conn, connErr := listener.Accept()
	if connErr != nil {
		log.Fatal("Error:", connErr)
	}

	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error;", readErr)
	}
	fmt.Println("message received:", message)
}

func runGuest(ip string) {
}
