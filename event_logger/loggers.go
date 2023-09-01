package eventlogger

import (
	"log"
	"os"
)

type Loggers struct {
	stdout_info_log *log.Logger
	stdout_debug_log *log.Logger
	stderr_log *log.Logger
	verbosity string
}

func NewLoggers(verbosity string) *Loggers {
	return &Loggers{
		stdout_info_log: log.New(os.Stdout, "INFO: ", log.Ltime),
		stdout_debug_log: log.New(os.Stdout, "DEBUG: ", log.Ltime),
		stderr_log: log.New(os.Stderr, "ERR: ", log.Ltime),
		verbosity: verbosity,
	}
}

func (l Loggers) INFO(s string) error {
	if l.verbosity == "debug" || l.verbosity == "info" {
		return l.stdout_info_log.Output(2, s)
	}
	return nil
}

func (l Loggers) DEBUG(s string) error {
	if l.verbosity == "debug" {
		return l.stdout_info_log.Output(2, s)
	}
	return nil
}

func (l Loggers) ERR(s string) error {
	return l.stdout_info_log.Output(2, s)
}