package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	doClientWork(client)
}

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string
		err := client.Call("KVStoreService.Watch", 10, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keyChanged)
	}()
	// 为什么要加这个等待呢？是为了上面的watch方法能先往Filter里面添加函数，不然后调用Set时，没有Filter函数，看不到输出。
	time.Sleep(time.Second * 1)
	err := client.Call(
		"KVStoreService.Set", [2]string{"abc", "abc-value"}, new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 3)
}
