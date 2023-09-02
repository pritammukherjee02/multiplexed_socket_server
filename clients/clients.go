package clients

import (
	"net"

	eventlogger "github.com/pritammukherjee02/multiplexed_socket_server/event_logger"
)

var Clients map[string]net.Conn
var loggers *eventlogger.Loggers

func ClientsMapInit(logger *eventlogger.Loggers){
	Clients = make(map[string]net.Conn)
	loggers = logger
}

func AppendClient(id string, conn net.Conn){
	Clients[id] = conn
	loggers.DEBUG("Appended client id = " + id)
}

func RemoveClient(id string){
	delete(Clients, id)
	loggers.DEBUG("Removed client id = " + id)
}

func GetConnectionForClient(id string) net.Conn {
	return Clients[id]
}