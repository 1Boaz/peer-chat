package Init

import (
	"log"
	"net"
)

func CreateServer(port string) net.Listener {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Couldn't use the selected port(" + port + "), Error: " + err.Error())
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal("Couldn't close the selected port(" + port + "), Error: " + err.Error())
		}
	}(listener)

	return listener
}
