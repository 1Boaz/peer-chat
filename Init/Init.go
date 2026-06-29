package Init

import (
	"net"
	"strconv"
)

func Init() net.Listener {
	port := Parse()

	listener := CreateServer(strconv.Itoa(port))

	return listener
}
