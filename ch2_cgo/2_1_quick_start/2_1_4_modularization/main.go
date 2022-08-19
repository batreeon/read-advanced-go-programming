package main
//void sayHello(const char*);
import "C"

func main() {
	C.sayHello(C.CString("hello world\n"))
}