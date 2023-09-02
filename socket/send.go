package socket

import (
	"net"

	"github.com/gobwas/ws/wsutil"
)

func SendData(data []byte, conn net.Conn) error {
	loggers.DEBUG("About to send data")
	// _, err := conn.Write(data)
	// err := wsutil.WriteClientText(conn, data)
	err := wsutil.WriteClientBinary(conn, data)
	if err != nil {
		loggers.ERR("Encountered error while writing to the client connection: " + err.Error())
	}
	loggers.DEBUG("Sent")
	return err
}