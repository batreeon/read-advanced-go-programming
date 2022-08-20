package main
// static void noreturn() {}
import "C"
import "fmt"
func main() {
	_, err := C.noreturn()
	fmt.Println(err)
	v := C.noreturn()
	fmt.Printf("%#v\n", v)
	fmt.Println(C.noreturn())
}