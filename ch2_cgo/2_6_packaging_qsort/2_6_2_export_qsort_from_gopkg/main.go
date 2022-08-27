package main

// extern int less(void*, void*);
import "C"
import (
	"advanced-go-programming/ch2_cgo/2_6_packaging_qsort/2_6_2_export_qsort_from_gopkg/qsort"
	"fmt"
	"unsafe"
)

// export less
func less(pa, pb unsafe.Pointer) C.int {
	a, b := *((*C.int)(pa)), *((*C.int)(pb))
	return a - b
}

func main() {
	nums := []int{7, 1, 6, 10, 9, 4}
	qsort.Sort(unsafe.Pointer(&nums[0]), len(nums), int(unsafe.Sizeof(nums[0])), qsort.CmpFunc(C.less))
	fmt.Println(nums)
}
