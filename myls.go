package main
import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func main(){
	aFlg := flag.Bool("a", false, "Include directory entries whose names begin with a dot (.).")
	tFlg := flag.Bool("t", false, "Sort by time modified (most recently modified first) before sorting the operands by lexicographical order.")
	flag.Parse()

	//カレントディレクトリ取得
	dir, err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}

	//指定ディクトリ配下のファイルからソート済みのファイルリストを取得
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println(fileInfos)
	}

	if *aFlag{
		for _, fileInfo := range fileInfos {
			fmt.Println(fileInfo.Name())
		}
		return
	}

	//更新時間順にソート
	if *tFlg{
		sort.Slice(fileInfos, func(i, j int) bool {return fileInfos[i].ModTime().After(fileInfos[j].ModTime()) })

		for _, fileInfo := range fileInfos {
			//隠しファイルは非表示  strings.HasPrefix　文字列の始まり（先頭)を真偽値で確認する != HasSuffix 文字列の終わり（末尾)を確認する
			if strings.Hasprefix(fileInfo.Name(), "."){
				continue
			}
			fmt.Println(fileInfo.Name())
		}
		return
	}

	//オプションなし
	for_, fileInfo := range fileInfos {
		if strings.Hasprefix(fileInfo.Name(), "."){
			continue
		}
		fmt.Println(fileInfo.Name())
	}


}