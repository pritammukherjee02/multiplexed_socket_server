package handlers

import (
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/pritammukherjee02/multiplexed_socket_server/clients"
)

func (h *Handlers) sendData(data []byte, conn net.Conn) error {
	err := wsutil.WriteServerMessage(conn, ws.OpText, data)
	if err != nil {
		h.loggers.ERR("Encountered error while writing to the client connection: " + err.Error())
		return err
	}
	h.loggers.DEBUG("Sent data: " + string(data))
	return err
}

func (h *Handlers) SocketWriteHandler(data []byte, id string) error {
	conn := clients.GetConnectionForClient(id)
	return h.sendData(data, conn)
}