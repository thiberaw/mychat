package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

func RunHost(ip string) {
	ipAndPort := ip + ":" + port

	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error: ", listenErr)
	}
	fmt.Println("Listening through tcp on ip ", ip)

	conn, connErr := listener.Accept()
	if connErr != nil {
		log.Fatal("Error:", connErr)
	}
	fmt.Println("connection accepted")

	for {
		handleRead(conn)
		handleReply(conn)
	}
}

func RunGuest(ip string) {
	ipAndPort := ip + ":" + port

	conn, dialError := net.Dial("tcp", ipAndPort)
	if dialError != nil {
		log.Fatal("Error: ", dialError)
	}

	for {
		handleReply(conn)
		handleRead(conn)
	}
}

func handleReply(conn net.Conn) {
	fmt.Print("Send message: ")

	replyReader := bufio.NewReader(os.Stdin)
	replyMessage, replyReaderError := replyReader.ReadString('\n')
	if replyReaderError != nil {
		log.Fatal("Error;", replyReaderError)
	}

	fmt.Fprint(conn, replyMessage)
}

func handleRead(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error;", readErr)
	}

	fmt.Print("message received:", message)

}
