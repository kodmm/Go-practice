package main

import "fmt"

func main() {
	g(0)
	fmt.Println("Program finished successfully!")

}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic("Panic in g() (major)")
	}
	defer fmt.Println("Defer in g()", i)
	fmt.Println("Printing in g()", i)
	g(i + 1)
}
