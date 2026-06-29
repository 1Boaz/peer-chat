package Init

import (
	"flag"
	"log"
)

func Parse() int {
	var Port = flag.Int("port", 31415, "port to listen and send on")
	flag.Parse()

	if *Port > 65_535 {
		log.Fatal("Port must be less than 65_535(max port number)")
	}

	return *Port
}
