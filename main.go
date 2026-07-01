package main

import (
	"log"
	"net"

	"github.com/1Boaz/peer-chat/Init"
)

// Main function should be minimal
func main() {
	server := Init.Init()

	// closes the connection when program is done
	defer func(server net.Listener) {
		err := server.Close()
		if err != nil {
			log.Fatal("Error closing server: " + err.Error())
		}
	}(server)
}
