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

//func newUDPsocket() {}

/*
package main

import "net"

type family string
type protocol string

const (
	//families
	AF_INET  family = ""
	AF_INET4 family = "4"
	AF_INET6 family = "6"

	//protocols
	SOCK_STREAM protocol = "tcp"
	SOCK_DGRAM  protocol = "udp"
)

type sock interface{
	socket
}

type socket struct {
	p        protocol
	f        family
	TCPAddr  net.TCPAddr
	conn     net.Conn
	listener net.Listener
}

func (s *socket) bind(IP net.IP, port int16) error {
	switch s.p {
	case "tcp":
		addr := new(net.TCPAddr)
		addr.IP = IP[:]
		addr.Port = int(port)

		s.TCPAddr = *addr

	case "udp":
		addr := new(net.UDPAddr)
		addr.IP = IP[:]
		addr.Port = int(port)

		network := string(s.p) + string(s.f)
		var err error
		s.conn, err = net.ListenUDP(network, addr)

		return err

	default:
		return &net.AddrError{}
	}

	return nil
}

//Listen is only for use with the TCP protocol
func (s *socket) listen() error {
	network := string(s.p) + string(s.f)

	var err error
	s.listener, err = net.ListenTCP(network, &s.TCPAddr)

	return err
}

//Accept is only for use with the TCP protocol
func (s *socket) accept() (*socket, error) {
	conn, err := s.listener.Accept()

	conn_sock := new(socket)
	conn_sock.conn =

	return conn, err
}

func (s *socket) send() {

}

func createSocket(f family, p protocol) socket {
	sock := new(socket)
	sock.p = p
	sock.f = f

	return *sock
}

package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

//families
const (
	AF_INET  = ""
	AF_INET4 = "4"
	AF_INET6 = "6"
)

type listener struct {
	l net.Listener
}

type TCPconnection struct {
	conn net.Conn
}

type UDPconnection struct {
	conn net.Conn
}

//converts ip string to byte slice
func convIPv4(IP string) []byte {
	splitIP := strings.Split(IP, ".")
	var byteIP []byte

	for _, i := range splitIP {
		section, _ := strconv.Atoi(i)
		byteIP = append(byteIP, byte(section))
	}

	return byteIP
}

func (l *listener) Accept() *TCPconnection {
	c, err := l.l.Accept()
	fmt.Println(err)

	conn := new(TCPconnection)
	conn.conn = c

	return conn
}

func (c *TCPconnection) send(data []byte) {
	c.conn.Write(data)
}

//data from connection is written into the data pointer parameter
func (c *TCPconnection) recv(buffer_size int, data *[]byte) {
	buffer := make([]byte, buffer_size)
	bytes_written, _ := c.conn.Read(buffer)
	buffer = buffer[:bytes_written]

	*data = buffer
}

func (c *UDPconnection) sendto(data []byte, address string){

}

func newTCPsocket(address string) *listener {
	spilt_address := strings.Split(address, ":")
	port, _ := strconv.Atoi(spilt_address[1])
	IP := spilt_address[0]

	TCPaddr := new(net.TCPAddr)
	TCPaddr.Port = port
	TCPaddr.IP = convIPv4(IP)

	l, _ := net.ListenTCP("tcp", TCPaddr)

	new_listener := new(listener)
	new_listener.l = l

	return new_listener
}

func newUDPsocket(address string) *UDPconnection {
	spilt_address := strings.Split(address, ":")
	port, _ := strconv.Atoi(spilt_address[1])
	IP := spilt_address[0]

	UDPaddr := new(net.UDPAddr)
	UDPaddr.Port = port
	UDPaddr.IP = convIPv4(IP)

	c, _ := net.ListenUDP("udp4", UDPaddr)

	new_conn := new(UDPconnection)
	new_conn.conn = c

	return new_conn
}

func main() {
	socket := newTCPsocket("127.0.0.1:3341")

	for {
		conn := socket.Accept()
		conn.send([]byte("hello there"))
	}
}

H E R M I T C R A F T
H E R M A T R I X
H _ _ _ _ _ _ _ _

*/
