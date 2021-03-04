package main

import(
	"fmt"
	"errors"

)

func main() {
	//recover関数をdeferで呼び出すことでリカバーを実現できる。 発生したパニックを取得する。
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%v", r)
		}
	}()

	panic(errors.New("Error"))
}