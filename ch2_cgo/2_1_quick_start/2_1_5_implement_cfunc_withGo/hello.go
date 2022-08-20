package main

import "C"
import "fmt"

// export sayHello
func sayHello(s *C.char) {
	// 因为参数接受的是C语言的字符串，所以需要转换成Go语言的字符串
	fmt.Print(C.GoString(s))
}