package main 

import "fmt"

func main() {
	ch1 := make(chan int)
	var ch2 chan string //ゼロ値はnil
	go func() { ch1 <- 100 }()
	go func() { ch2 <- "hi"}() 

	select{
	case v1 := <- ch1:
		fmt.Println(v1)
	case v2 := <- ch2: //nilの場合は無視
		fmt.Println(v2)
	}
}