package main
import (
	"os"
	"fmt"
	"log"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

func main() {
	/*f, err := os.Open("sample.go")
	p := make([]byte, m) //1. スライス確保
	n, err := f.Read(p) //2. Read実行
	if n < m {
		log.Fatalf("%dバイト読み込もうとしましたが、%dバイトしか読めませんでした¥n", m, n)
	}
	if err != nil {
		log.Fatalf("読み込み中にエラーが発生しました：%v¥n", err)
	}
	*/
	fmt.Fprintln(os.Stderr, "error") //接頭語にFがついているものは、書き込み先を明示的に指定できる。 os.Stderr(標準エラー出力)
	os.Exit(1) //0 : 成功(default)
}