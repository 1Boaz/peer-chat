package Init

import (
	"errors"
	"flag"
	"net"
	"strconv"
)

// Init Main initializing function to call other functions in this package in sequence
func Init() (*net.UDPConn, error) {
	port, err := parse()
	if err != nil {
		return nil, err
	}
	return createServer(strconv.Itoa(port))
}

// createServer Initializes an udp server and returns the listener(server)
func createServer(port string) (*net.UDPConn, error) {
	// Resolve the server listening address
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		return nil, err
	}

	// Creates a listener on the resolved address
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	return listener, nil
}

// parse parses the port cmd argument
func parse() (int, error) {
	// Crates a flag named port, with a default value and a help message
	var Port = flag.String("port", "31415", "port to listen and send on")
	flag.Parse()

	// Converts the string flag into integer helping with error handling
	port, err := strconv.Atoi(*Port)
	if err != nil {
		return 0, errors.New("invalid port, port must be a number")
	}

	// verifies flag(port) number is valid
	if port > 65_535 || port < 1 {
		return 0, errors.New("port must be between 1 and 65_535")
	}
	return port, nil
}
