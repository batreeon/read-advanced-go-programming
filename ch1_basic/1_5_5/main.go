package main

import (
	"fmt"
	"time"
)

var a string

func f1() {
	fmt.Println(a)
}

func hello1() {
	a = "hello world"
	go f1()
}

var b string

func f2() {
	b = "ni hao"
}

func hello2() {
	go f2()
	go f2()
	go f2()
	go f2()
	fmt.Println(b)
}

func main() {
	hello2()
	hello1()
	time.Sleep(time.Second)
}
