package socket

import (
	"fmt"
	"net"
	"reflect"
	"sync"
	"syscall"

	"github.com/pritammukherjee02/multiplexed_socket_server/clients"
	"golang.org/x/sys/unix"
)

type ClientConectionsEpoll struct {
	id string
	conn net.Conn
}

type epoll struct {
	fd          int
	connections map[int]ClientConectionsEpoll
	lock        *sync.RWMutex
}

func CreateEpoll() (*epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}

	// Makes an empty map of client ids mapped to their connections
	clients.ClientsMapInit(loggers)

	return &epoll{
		fd:          fd,
		lock:        &sync.RWMutex{},
		connections: make(map[int]ClientConectionsEpoll),
	}, nil
}

func (e *epoll) Add(id string, conn net.Conn) error {
	// Extract file descriptor associated with the connection
	fd := websocketFD(conn)
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, &unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, Fd: int32(fd)})
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	e.connections[fd] = ClientConectionsEpoll{
		id: id,
		conn: conn,
	}
	clients.AppendClient(id, conn)
	
	loggers.DEBUG(fmt.Sprintf("Append: Total number of connections: %v", len(e.connections)))
	return nil
}

func (e *epoll) Remove(conn net.Conn) error {
	fd := websocketFD(conn)
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	id := e.connections[fd].id
	delete(e.connections, fd)
	clients.RemoveClient(id)

	loggers.DEBUG(fmt.Sprintf("Remove: Total number of connections: %v", len(e.connections)))
	return nil
}

func (e *epoll) Wait() ([]ClientConectionsEpoll, error) {
	events := make([]unix.EpollEvent, 100)
	n, err := unix.EpollWait(e.fd, events, 100)
	if err != nil {
		return nil, err
	}
	e.lock.RLock()
	defer e.lock.RUnlock()
	var connections []ClientConectionsEpoll
	for i := 0; i < n; i++ {
		clientConn := e.connections[int(events[i].Fd)]
		connections = append(connections, clientConn)
	}
	return connections, nil
}

func websocketFD(conn net.Conn) int {
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")

	return int(pfdVal.FieldByName("Sysfd").Int())
}