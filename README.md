# socket
Python-like socket libary for GO

#How to use
TCP server
``` go
socket := NewTCPsocket(AF_INET4)
socket.Bind([]byte{127, 0 , 0, 1}, 6900)
socket.Listen()

for{
	conn := socket.Accept()

	conn.Send([]byte("Hello"))
	data, _, _ :=  conn.Recv(1024)
	fmt.Println(string(data))
}
```
TCP client
``` go
socket := NewTCPsocket(AF_INET4)
socket.Connect([]byte{127, 0 , 0, 1}, 6900)

data, _, _ :=  conn.Recv(1024)
fmt.Println(string(data))
conn.Send([]byte("Hello"))
```
