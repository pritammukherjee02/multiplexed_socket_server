package handlers

import (
	"fmt"
)

// Handles what happens when the socket server receives some data from a specific
// Client with a unique ID
func (h *Handlers) SocketReadHandler(recv []byte, clientId string) error {
	h.loggers.INFO(fmt.Sprintf("Received (id: %s): %s", clientId, string(recv)))
	return nil
}