package main

import (
	"fmt"
	"regexp"
)

//パッケージの初期化時に不適切な初期化ならpanicを発生させる
var validId = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

func main() {
	fmt.Println(validId.MatchString("adam[23]"))
	validID2, err := regexp.Compile(`^[a-z]+\[[0-9]+\]$`)
	// 関数内で行う場合はエラー処理をする
	if err != nil {
	}
	fmt.Println(validID2.MatchString("adam[23]"))
}
