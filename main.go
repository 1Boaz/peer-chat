package main

import (
	"log"
	"net"

	"github.com/1Boaz/peer-chat/Init"
)

// Main function should be minimal
func main() {
	server := Init.Init()

	// Closes the connection when program is done
	defer func(server *net.UDPConn) {
		err := server.Close()
		if err != nil {
			log.Fatal("Error closing server: " + err.Error())
		}
	}(server)
}
