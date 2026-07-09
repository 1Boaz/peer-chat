package Init

import (
	"errors"
	"net"
	"testing"
)

func TestCreateServer(t *testing.T) {
	// Creates a function to verify each case more easily
	// exp stands for expected
	verify := func(expOut string, expErr error, out *net.UDPConn, err error) {
		if expErr != nil {
			// Makes sure the error is the same as expected
			if err.Error() != expErr.Error() {
				t.Errorf("CreateServer() returned %v but expected %v", err, expErr)
			}
		} else {
			if err != nil {
				t.Errorf("CreateServer() returned error: %v; when it shouldn`t", err)
			} else {
				// This checks the expected output(expOut) against the host + port
				if expOut != out.LocalAddr().String() {
					t.Errorf("CreateServer() returned %v but expected %v", out.LocalAddr().String(), expOut)
				}
			}
		}
	}

	// Case for a valid input
	server, err := createServer("8080")
	if server != nil {
		defer server.Close()
	}
	verify("[::]:8080", nil, server, err)

	// Case for an invalid port number input
	server, err = createServer("1000000")
	if server != nil {
		defer server.Close()
	}
	verify("", errors.New("invalid port: 1000000"), server, err)

	// Case for an invalid port number input
	server, err = createServer("10000.5")
	if server != nil {
		defer server.Close()
	}
	verify("", errors.New("invalid port: 10000.5"), server, err)

	// Case for a noninteger port
	server, err = createServer("-21")
	if server != nil {
		defer server.Close()
	}
	verify("", errors.New("invalid port: -21"), server, err)

	// Case for a reserved port
	server, err = createServer("22")
	if server != nil {
		defer server.Close()
	}
	verify("", errors.New("listen udp :22: bind: permission denied"), server, err)
}
