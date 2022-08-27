package qsort
/*
#include <stdlib.h>
typedef int (*qsort_cmp_cfunc_t)(const void*, const void*);
*/
import "C"
import "unsafe"
type CmpFunc C.qsort_cmp_cfunc_t

func Sort(base unsafe.Pointer, num int, size int, cmp CmpFunc) {
	C.qsort(base, C.size_t(num), C.size_t(size), C.qsort_cmp_cfunc_t(cmp))
}