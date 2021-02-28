package main

import(
	"fmt"
	"errors"
)


var MyError = myError()


func myError() error { return errors.New("myErr")}

func simpleError() error {
	return MyError 
}

func wrappedError() error {
	err := simpleError()
	return fmt.Errorf("%w", err)
}

// func main() {
// 	err := simpleError()
// 	fmt.Printf("%T \n", err)

// 	err = wrappedError()
// 	fmt.Printf("%T", err)

// }

func main () {
	err := wrappedError()
	if err != nil {
		switch err {
		case MyError:
			fmt.Println("MyError:", err)
		default:
			fmt.Println("defualt:", err)
		}
	}
}