package socket

import (
	"net"
)

type family string

const (
	//AF_INET  family = ""
	AF_INET4 family = "4"
	//AF_INET6 family = "6"
)

type TCPsocket struct {
	family   family
	addr     *net.TCPAddr
	listener net.Listener
	conn     net.Conn
}

func (s *TCPsocket) Bind(IP []byte, port uint16) {
	s.addr = new(net.TCPAddr)
	s.addr.IP = IP
	s.addr.Port = int(port)
}

func (s *TCPsocket) Listen() error {
	network := "tcp" + string(s.family)

	var err error
	s.listener, err = net.ListenTCP(network, s.addr)

	return err
}

func (s *TCPsocket) Accept() *TCPsocket {
	connection, _ := s.listener.Accept()

	TCPsock := new(TCPsocket)
	TCPsock.conn = connection

	return TCPsock
}

/*
func (s *TCPsocket) connectLocalRemote(localIP []byte, localPort uint16, remoteIP []byte, remotePort uint16) (*TCPconnection, error) {
	local_addr := new(net.TCPAddr)
	local_addr.IP = localIP
	local_addr.Port = int(localPort)

	remote_addr := new(net.TCPAddr)
	remote_addr.IP = remoteIP
	remote_addr.Port = int(remotePort)

	network := "tcp" + string(s.family)
	conn, err := net.DialTCP(network, local_addr, remote_addr)

	TCPconn := new(TCPconnection)
	TCPconn.conn = conn
	return TCPconn, err
}

*/

func (s *TCPsocket) Connect(remoteIP []byte, remotePort uint16) error {
	remote_addr := new(net.TCPAddr)
	remote_addr.IP = remoteIP
	remote_addr.Port = int(remotePort)

	network := "tcp" + string(s.family)
	conn, err := net.DialTCP(network, nil, remote_addr)

	s.conn = conn
	return err
}

func (s *TCPsocket) Send(data []byte) (n int, err error) {
	return s.conn.Write(data)
}

func (s *TCPsocket) Recv(buffer_size int) (buffer []byte, err error, n int) {
	buffer = make([]byte, buffer_size)
	n, err = s.conn.Read(buffer)
	buffer = buffer[:n]

	return buffer, err, n
}

func (s *TCPsocket) Close() error {
	return s.conn.Close()
}

func NewTCPsocket(family family) *TCPsocket {
	socket := new(TCPsocket)
	socket.family = family

	return socket
}
