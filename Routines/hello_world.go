package main

import (
	"time"
)

func asyncOp(ch chan string, amt int, identifier string) {
	time.Sleep(time.Duration(amt) * time.Millisecond)
	ch <- identifier
}

// func main() {
// 	ch := make(chan string)
// 	go asyncOp(ch, 4, "a")
// 	go asyncOp(ch, 10000, "b")
// 	go asyncOp(ch, 1, "e")
// 	go asyncOp(ch, 1, "f")
// 	msg, ok := <-ch
// 	fmt.Printf("msg='%s', ok='%v'\n", msg, ok)
// 	msg, ok = <-ch
// 	fmt.Printf("msg='%s', ok='%v'\n", msg, ok)
// 	msg, ok = <-ch
// 	fmt.Printf("msg='%s', ok='%v'\n", msg, ok)
// 	msg, ok = <-ch
// 	fmt.Printf("msg='%s', ok='%v'\n", msg, ok)
// }
