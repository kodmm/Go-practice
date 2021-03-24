package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}
}

// 7行目のタスクをスタックに格納しLIFOで出力する