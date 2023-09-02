package socket

import (
	"net"
)

func SendData(data []byte, conn net.Conn) error {
	loggers.DEBUG("About to send data")
	_, err := conn.Write(data)
	loggers.DEBUG("Sent")
	return err
}