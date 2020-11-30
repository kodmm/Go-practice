package main

import (
	"fmt"
	"os"
	"log"
	"path/filepath"
)

func main() {
	//pathに実行fファイルの絶対パスが入る。
	path, error := os.Executable()
	if error != nil {
		fmt.Println(os.Stderr, "読み込みに失敗しました。", error)
	}
	fmt.Println(path)
	path = filepath.Dir(path)
	fmt.Println(path)

	r, error := os.Open(os.Args[1]) //os.Open: ファイルを開く。読み込み専用。ファイルがなければ作成されずエラー。
	if error != nil {
		log.Fatal(error)
	}
	//os.OpenFile: ファイルを開く。読み書き作成などのオプション指定可。
	w, error := os.OpenFile(os.Args[2], os.O_WRONLY | os.O_CREATE, 0644) 
	if error != nil {
		log.Fatal(error)
	}
	buf := make([]byte, 5)
	for {
		//nはバイト数を示す。
		n, _ := r.Read(buf)
		fmt.Println(string(buf[:n]), n)
		if n == 0 {
			break
		}
		w.Write(buf[:n])
	}

}

// go run mycopy.go read.txt write.go