package main
/*
#include <errno.h>
static int div(int a, int b) {
	if (b == 0) {
		errno = EINVAL;
		return 0;
	}
	return a/b;
}
*/
import "C"
import "fmt"

func main() {
	v0, errno0 := C.div(2, 1)
	fmt.Println(v0, errno0)
	v1, errno1 := C.div(1, 0)
	fmt.Println(v1, errno1)
}