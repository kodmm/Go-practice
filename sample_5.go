package main
import (
	"fmt"
	"os"
	"flag"
	"strings"
)

//設定される変数のポインタを取得
var msg = flag.String("msg", "デフォルト値", "説明")
var n int

func init() {
	//ポインタを指定して設定を予約
	flag.IntVar(&n, "n", 1, "回数")
	

}

func main() {
	fmt.Println(os.Args)
	
	flag.Parse()
	//プログラム処理できる書式に変更
	fmt.Println(strings.Repeat(*msg, n))

	//os.Argsだとフラグも含まれるがflag.Argsはフラグの部分は除外される
	fmt.Println(flag.Args())
	
}

//go run sample_5.go -msg Yamada -n 3