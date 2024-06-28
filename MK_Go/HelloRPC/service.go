package main

import (
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil

}

func main() {
	_ = rpc.RegisterName("HelloService", &HelloService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic("ListenTCP error:" + err.Error())
	}
	conn, err := listener.Accept()
	if err != nil {
		panic("Accept error:" + err.Error())
	}
	rpc.ServeConn(conn)
}
