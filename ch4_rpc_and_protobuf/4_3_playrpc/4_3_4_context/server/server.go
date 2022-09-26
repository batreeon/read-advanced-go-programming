package main

import (
	"4_3_4_context/helloservice"
	"log"
	"net"
	"net/rpc"
)

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen err: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept err: ", err)
		}
		go func() {
			defer conn.Close()
			
			s := rpc.NewServer()
			s.Register(&helloservice.HelloService{
				Conn: conn,
			})
			s.ServeConn(conn)
		}()
	}
}