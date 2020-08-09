package main
import (
	"os"
	"fmt"
	
)



func main() {
	//読み込み用にファイルを開く
	sf, err := os.Open("src")
	if err != nil {
		fmt.Printf("%v", err)
	}
	
	//関数終了時に閉じる
	defer sf.Close()

	//書き込み用にファイルを開く
	df, err := os.Create("sample_8.go")
	if err != nil {
		fmt.Printf("%v", err)
	}

	//関数終了後に閉じる　
	defer func(){
		if err := df.Close(); err != nil {
			
		}
	}()

}