package main
import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"io"
	"path/filepath"
	"time"
)

var t int 

func init() {
	flag.IntVar(&t, "t", 1, "制限時間(分)")
	flag.Parse()
}

func main() {
	var (
		files = flag.Args()
		path, err = os.Executable()
		ch_rcv = myinput(os.Stdin)
		tm = time.After(time.Duration(t) * time.Minute)
		words []string
		score int = 0
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
		os.Exit(0)
	}else if len(files) != 1 {
		fmt.Fprintln(os.Stderr, "指定できるファイル数は1つだけです。")
		os.Exit(0)
	}
	fmt.Println(files)
	path = filepath.Dir(path)
	fmt.Println(files)
	words, err = getwords_from(filepath.Join(path, files[0]))
	if err != nil {
		fmt.Fprintln(os.Stderr, "読み込み失敗しました", err)
		os.Exit(0)
	}

	fmt.Println("タイピングゲームを始めます。制限時間は", t, "分1語1点", len(words), "点満点")
	//送信用チャネル
	for i := true; i && score < len(words); {
		question := words[score]
		fmt.Print(question)
		fmt.Print(">")
		select {
		case x := <-ch_rcv:
			if question == x {
				score++
			}
		case <-tm:
			fmt.Println("制限時間が過ぎました")
			i = false
		}
		fmt.Println("点数: ", score, "点 / ", len(words), "点")
	}
	fmt.Println(path)
}

func myinput(r io.Reader) <-chan string {
	ch1 := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch1 <- s.Text()
		}

	}()
	return ch1
}


func getwords_from(txt_path string) ([]string, error) {
	var words []string 
	sf, err := os.Open(txt_path)
	if err != nil {
		return nil, err
	} else {
		scanner := bufio.NewScanner(sf)
		for scanner.Scan() {
			words = append(words, scanner.Text())
		}
	}
	return words, err
}