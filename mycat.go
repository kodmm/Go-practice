package main
import (
	"bufio"
	"path/filepath"
	"fmt"
	"os"
	"flag"
)

func main () {
	var n = flag.Bool("n", false, "通し番号を付与する")
	flag.Parse()
	var (
		//Parseした値をString型で取得する
		files = flag.Args()
		//path に　実行ファイルの絶対パスが入る
		path, err = os.Executable()
	)

	if err != nil {
		fmt.Println(os.Stderr, "読み込みに失敗しました", err)
	}

	//実行ファイルのディレクトリ名を取得
	path=filepath.Dir(path)
	
	//通し番号の用意
	i := 1

	for x := 0; x <len(files); x++ {
		sf, err := os.Open(filepath.Join(path,files[x]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "読み込みに失敗しました。", err)
		}else{
			scanner := bufio.NewScanner(sf)
			for ;scanner.Scan();i++{
				if *n{
					//オプションがある場合
					fmt.Printf("%v",i)
				}
				fmt.Println(scanner.Text())
			}
		}
	}

}
