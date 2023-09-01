package socket

import (
	"net"
)

var Clients map[string]net.Conn

func ClientsMapInit(){
	Clients = make(map[string]net.Conn)
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