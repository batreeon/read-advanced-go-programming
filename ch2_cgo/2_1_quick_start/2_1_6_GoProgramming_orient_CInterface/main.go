package main

// 注意，下面两句不可少
// void sayHello1(char*);
// void sayHello2(_GoString_);

import "C"
import "fmt"

func main() {
	sayHello1(C.CString("1:hello world\n"))
	sayHello2("2:hello world\n")
}

func sayHello1(s *C.char) {
	fmt.Print(C.GoString(s))
}

func sayHello2(s string){
	fmt.Print(s)
}