package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go func() {
		time.Sleep(time.Second * 3)
		done <- true
	}()
	<- done 
	fmt.Println("done")
}