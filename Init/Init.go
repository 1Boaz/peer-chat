package Init

import (
	"flag"
	"log"
	"net"
	"strconv"
)

// Init Main initializing function to call other functions in this package in sequence
func Init() *net.UDPConn {
	port := parse()
	return createServer(strconv.Itoa(port))
}

// createServer Initializes an udp server and returns the listener(server)
func createServer(port string) *net.UDPConn {
	// Resolve the server listening address
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		log.Fatal("Error resolving address for UDP server: " + err.Error())
	}

	// Creates a listener on the resolved address
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("Error starting UDP server: " + err.Error())
	}

	return listener
}

// parse parses the port cmd argument
func parse() int {
	// crates a flag named port, with a default value and a help message
	var Port = flag.Int("port", 31415, "port to listen and send on")
	flag.Parse()

	// verifies flag(port) number is valid
	if *Port > 65_535 || *Port < 1 {
		log.Fatal("Port must be lower than 65_535(max port number) and higher than 1")
	}

	return *Port
}
