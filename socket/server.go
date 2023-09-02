package socket

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"syscall"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/google/uuid"
	eventlogger "github.com/pritammukherjee02/multiplexed_socket_server/event_logger"
)

var epoller *epoll
var loggers *eventlogger.Loggers

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		return
	}

	// The id will be infered from either the request body or the headers in some way or another
	id := uuid.NewString()

	if err := epoller.Add(id, conn); err != nil {
		loggers.ERR("Failed to add connection: " + err.Error())
		conn.Close()
	}
}

func RunSocketServer(loggers_instance *eventlogger.Loggers) {
	loggers = loggers_instance
	loggers.INFO("Starting Gobwas socket server...")
	loggers.DEBUG("Setting ulimit RLIMIT_NOFILE to rLimit.Max...")

	// Increase resources limitations
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	loggers.DEBUG("Completed ulimit edit")

	// Enable pprof hooks
	go func() {
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			loggers.ERR("pprof failed: " + err.Error())
		}
	}()

	// Start epoll
	var err error
	epoller, err = CreateEpoll()
	if err != nil {
		panic(err)
	}

	go Start()

	http.HandleFunc("/", wsHandler)
	if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil {
		loggers.ERR(err.Error())
		os.Exit(1)
	}
}

func Start() {
	loggers.INFO("Starting UNIX Epoll system...")
	for {
		clientConnections, err := epoller.Wait()
		if err != nil {
			loggers.ERR("Failed to wait epoll: " + err.Error())
			continue
		}
		for i, clientConnection := range clientConnections {
			if clientConnection.conn == nil {
				break
			}
			if data, _, err := wsutil.ReadClientData(clientConnection.conn); err != nil {
				if err := epoller.Remove(clientConnection.conn); err != nil {
					loggers.ERR("Failed to remove connection: " + err.Error())
				}
				clientConnection.conn.Close()
			} else {
				loggers.INFO(fmt.Sprintf("%d > data: %s", i, string(data)))
				SendData([]byte("Hello buoy"), clientConnection.conn)
			}
		}
	}
}