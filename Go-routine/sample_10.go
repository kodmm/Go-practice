package main
import (
	"fmt"
	"time"
)

// func main() {
// 	go func () {
// 		fmt.Println("別のゴールーチン")
// 	}()
// 	fmt.Println("mainゴールーチン")
// 	time.Sleep(50*time.Millisecond)
// }

func main() {
	defer fmt.Println("main done")
	go func() {
		defer fmt.Println("goroutine1 done")
		time.Sleep(3 * time.Second)
	}()

	go func() {
		defer fmt.Println("goroutine2 done")
		time.Sleep(1 * time.Second)
	}()
	time.Sleep(5 * time.Second)
}