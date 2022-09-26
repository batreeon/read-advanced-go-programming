package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dial err: ", err)
	}

	var reply string
	err = conn.Call("HelloService.Hello", "mike", &reply)
	if err != nil {
		log.Println("Call Hello err: ", err)
	}

	// 虽然reply在Login方法里没有用到，但是也还是传进去吧。不然有时候会报错，
	// 好像即使没用到也会往reply里去写，你传个""，会报错往空地址写
	// 在这里，第一个调用没报，下面第二个调用报了
	err = conn.Call("HelloService.Login", "", &reply)
	if err != nil {
		log.Println("Call Login err: ", err)
	}
	
	err = conn.Call("HelloService.Login", "user:password", &reply)
	if err != nil {
		log.Println("Call Login err: ", err)
	}

	err = conn.Call("HelloService.Hello", "kevin", &reply)
	if err != nil {
		log.Println("Call Hello err: ", err)
	} else {
		fmt.Println(reply)
	}
}
