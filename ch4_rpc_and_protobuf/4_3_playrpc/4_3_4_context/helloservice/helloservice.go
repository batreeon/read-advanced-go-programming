package helloservice

import (
	"fmt"
	"log"
	"net"
)

// 注意这里Conn，IsLogin首字母要大写，不然server里面看不到这两个字段
type HelloService struct {
	Conn    net.Conn
	IsLogin bool
}

func (hs *HelloService) Login(request string, reply *string) error {
	if request != "user:password" {
		return fmt.Errorf("auth failed!")
	}
	hs.IsLogin = true
	log.Println("login successed!")
	return nil
}

func (hs *HelloService) Hello(request string, reply *string) error {
	if !hs.IsLogin {
		return fmt.Errorf("please login before do other operations!")
	}
	*reply = "hello:" + request + ", from " + hs.Conn.RemoteAddr().String()
	return nil
}
