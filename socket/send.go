package socket

import "net"

func SendData(data []byte, conn net.Conn) error {
	_, err := conn.Write(data)
	return err
}