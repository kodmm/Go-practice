package main 
import "fmt"

func makeCh() chan int {
	return make(chan int)
}

//受信専用channel
func recvCh(recv <-chan int) int {
	return <- recv
}

func main() {
	ch := makeCh()
	go func(ch chan <- int ){ ch <- 100}(ch) //送信専用channel
	fmt.Println(recvCh(ch))
}