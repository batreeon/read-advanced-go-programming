package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	goroutineId := getGoId()
	fmt.Println(goroutineId)
}

func getGoId() int {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.Trim(string(buf[:n]), " ")
	)
	ss := strings.TrimPrefix(stk, "goroutine ")
	idFiled := strings.Fields(ss)[0]
	id, err := strconv.Atoi(idFiled)
	if err != nil {
		log.Fatalf("can not get goroutine id, err: %v", err)
	}
	return id
}
