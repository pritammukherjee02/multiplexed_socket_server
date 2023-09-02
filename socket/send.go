package socket

import (
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func SendData(data []byte, conn net.Conn) error {
	err := wsutil.WriteServerMessage(conn, ws.OpText, data)
	if err != nil {
		loggers.ERR("Encountered error while writing to the client connection: " + err.Error())
		return err
	}
	loggers.DEBUG("Sent data: " + string(data))
	return err
}