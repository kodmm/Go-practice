package main

import(
	"errors"
	"fmt"
)

func main() {
	err := errors.New("cause")
	err = fmt.Errorf("first: %w", err)
	err = fmt.Errorf("second: %w", err)
	err = fmt.Errorf("third: %w", err)
	fmt.Println(Cause(err))
}

func Cause(err error) error {
	for err != nil {
		u, ok := err.(interface {
			Unwrap() error 
		})
		if !ok {
			break
		}
	
		err = u.Unwrap()
	}
	return err
}