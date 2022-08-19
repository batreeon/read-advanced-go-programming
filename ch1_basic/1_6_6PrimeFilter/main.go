package main

import (
	"fmt"
	"advanced-go-programming/ch1_basic/1_6_6PrimeFilter/primefilter"
)

func main() {
	ch := primefilter.GenerateNatural() // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = primefilter.PrimeFilter(ch, prime) // 基于新素数构造的过滤器
	}
}