package handlers

import (
	eventlogger "github.com/pritammukherjee02/multiplexed_socket_server/event_logger"
)

type Handlers struct {
	loggers *eventlogger.Loggers
}

func NewHandlers(logger *eventlogger.Loggers) *Handlers {
	return &Handlers{
		loggers: logger,
	}
}