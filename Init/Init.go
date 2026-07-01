package Init

import (
	"flag"
	"log"
	"net"
	"strconv"
)

// Init Main initializing function to call other functions in this package in sequence
func Init() net.Listener {
	port := parse()
	return createServer(strconv.Itoa(port))
}

// createServer Initializes a tcp server and return the listener(server)
func createServer(port string) net.Listener {
	listener, err := net.Listen("tcp", ":"+port)
	// verifies that net.Listen executed without errors
	if err != nil {
		log.Fatal("Couldn't use the selected port(" + port + "), Error: " + err.Error())
	}

	log.Print("Listening on: " + port)
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
