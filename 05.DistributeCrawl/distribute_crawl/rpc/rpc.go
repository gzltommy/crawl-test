package rpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, serveice interface{}) error {
	err := rpc.Register(serveice)
	if err != nil {
		return err
	}

	listen, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error : %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)

	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
