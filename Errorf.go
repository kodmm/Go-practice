package main

import (
	"errors"
	"fmt"
	"os"
	
)

func main() {
	err := fmt.Errorf("bar: %w", errors.New("fuga"))
	fmt.Println(err)
	fmt.Println(errors.Unwrap(err))
	
	if errors.Is(err, os.ErrExist) {
		fmt.Println(err)
	}else {
	}

	var pathError *os.PathError
	if errors.As(err, &pathError) {
		fmt.Println("Failid at Path:", pathError.Path)
		fmt.Printf("%#v \n", pathError)
	}else {
		fmt.Println(err)
	}

	ErrFoo := errors.New("foo error")

	wrapped := fmt.Errorf("wrapped woo: %w", ErrFoo)
	if errors.Is(wrapped, ErrFoo) {
		fmt.Printf("this error is caused by %#v", ErrFoo)
	}
}