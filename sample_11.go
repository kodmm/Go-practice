package main

import (
	"fmt"
)

func main() {
//型アサーション
// > インターフェース型の値を任意の型にキャストする
var v interface{}
v = 100
n, ok := v.(int)
fmt.Println(n, ok)

s, ok := v.(string)
fmt.Println(s, ok)


var i interface{}
i = 100
switch v := i.(type) {
	case int:
		fmt.Println(v*2)
	case string:
		fmt.Println(v+ "hoge")
	default:
		fmt.Println("default")
}
}