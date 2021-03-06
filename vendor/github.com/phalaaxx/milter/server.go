// Package milter A Go library for milter support
package milter

import (
	"net"
)

// Init initializes milter options
// multiple options can be set using a bitmask
type Init func() (Milter, OptAction, OptProtocol)

// RunServer provides a convenient way to start a milter server
func RunServer(server net.Listener, init Init) error {
	for {
		// accept connection from client
		client, err := server.Accept()
		if err != nil {
			return err
		}
		// create milter object
		milter, actions, protocol := init()
		session := milterSession{
			actions:  actions,
			protocol: protocol,
			sock:     client,
			milter:   milter,
		}
		// handle connection commands
		go session.HandleMilterCommands()
	}
}
