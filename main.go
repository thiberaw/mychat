package main

import (
	"flag"
	"mychat/lib"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "listen on the specified ip address")
	flag.Parse()

	if isHost {
		connIp := os.Args[2]
		lib.RunHost(connIp)
	} else {
		connIp := os.Args[1]
		lib.RunGuest(connIp)
	}
}
