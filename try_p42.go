package main
import "fmt"

func main() {
	for i:= 0; i <= 100; i++ {
		if i%2 == 1 {
			fmt.Println(i ,"- Odd number")
		}else {
			fmt.Println(i ,"- Even number")
		}
	}
	for n:= 0; n <= 10; n++ {
		switch {
			case n%2 == 1: 
				fmt.Printf("%d - Odd number\n", n)
			case n%2 == 0: 
				fmt.Println(n, "- Even number")
			default: 
				fmt.Println("default")
			}
		}
	

}


