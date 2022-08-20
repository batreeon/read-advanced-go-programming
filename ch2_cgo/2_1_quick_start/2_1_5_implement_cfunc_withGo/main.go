package main
// #include <hello.h>
import "C"
func main() {
	sayHello(C.CString("hello world\n"))
}