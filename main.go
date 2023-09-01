package main

import (
	"flag"
	"io"
	"os"

	eventlogger "github.com/pritammukherjee02/multiplexed_socket_server/event_logger"
	"github.com/pritammukherjee02/multiplexed_socket_server/socket"
)

var (
	verbosity          = flag.String("verbosity", "debug", "Log Verbosity Level")
)

func main() {
	// Getting logger verbosity level
	flag.Usage = func() {
		io.WriteString(os.Stderr, `Multiplexed Socket Server
Example usage: ./multiplexed_socket -verbosity=debug
`)
		flag.PrintDefaults()
	}
	flag.Parse()

	loggers := eventlogger.NewLoggers(*verbosity)

	socket.RunSocketServer(loggers)
}