package main

import (
	"4_3_2_watch_based_on_rpc/kvstore"
	"log"
	"net"
	"net/rpc"
)

func main() {
	kvStoreServcice := kvstore.NewKVStoreService()
	rpc.RegisterName("KVStoreService", kvStoreServcice)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	for {
		go rpc.ServeConn(conn)
	}
}
