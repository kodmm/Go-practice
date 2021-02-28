package main

import(
	"errors"
	"fmt"
)

var (
	MyError = myError()
)

func myError() error { return errors.New("myErr")}

func simpleError() error {
	return MyError 
}

//! Go1.12以前の書き方
// func main() {
// 	err := simpleError()
// 	if err != nil {
// 		fmt.Println(err) //myError
// 	}
// }

// func main() {
// 	err := simpleError()
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Printf("%T \n", err)
// 		switch err {
// 		case MyError:
// 			fmt.Println("MyError:", err)
// 		default:
// 			fmt.Println("default", err)
// 		}
// 	}
// }

//! Is ver

func main() {
	err := simpleError()
	if errors.Is(err, MyError) {
		fmt.Println(err)
	}
}
