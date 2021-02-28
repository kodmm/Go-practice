package main

import(
	"errors"
	"fmt"
)

type InvalidChar struct {
	err error
}

func (ic *InvalidChar) Error() string {
	ic.err = errors.New("INVALID CHARACTER")
	return fmt.Errorf("%w", ic.err).Error()
}

func Parse() error {
	return &InvalidChar{}
}

func main() {
	err := Parse()
	fmt.Println(err)
	if ierr, ok := err.(*InvalidChar); ok {
		fmt.Println(ierr)
	}
}