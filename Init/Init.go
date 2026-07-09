package Init

import (
	"errors"
	"flag"
	"net"
)

// Init Main initializing function to call other functions in this package in sequence
func Init() (*net.UDPConn, error) {
	port, err := parse()
	if err != nil {
		return nil, err
	}
	return createServer(port)
}

// createServer Initializes an udp server and returns the listener(server)
func createServer(port string) (*net.UDPConn, error) {
	// Resolve the server listening address
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		return nil, errors.New("invalid port: " + port)
	}

	// Creates a listener on the resolved address
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	return listener, nil
}

// parse parses the port cmd argument
func parse() (string, error) {
	// Crates a flag named port, with a default value and a help message
	var Port = flag.String("port", "31415", "port to listen and send on")
	flag.Parse()

	return *Port, nil
}
