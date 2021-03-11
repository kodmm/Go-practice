package main

import (
	"time"
	"fmt"
)

func main() {
	done := false
	go func() {
		time.Sleep(3 * time.Second)
		done = true
	}()
	for !done {
		time.Sleep(time.Millisecond)
	}
	fmt.Println("done!")
}