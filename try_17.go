package main
import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
	"log"
)

type RuneScanner struct {
	r io.Reader
	buf [16]byte //16バイトを格納する
}

func NewRuneScanner(r io.Reader) *RuneScanner {
	fmt.Println(r)
	return &RuneScanner{r: r}
}


func (s *RuneScanner) Scan() (rune, error) {
	//string型はbyte単位なので絵文字などの大きなbyte文字は文字化けしてしまうがcode pointを単位として扱うrune型ならば文字化けしない
	n, err := s.r.Read(s.buf[:])
	fmt.Println(s.buf)
	if err != nil {
		return 0, err
	}
 
	r, size := utf8.DecodeRune(s.buf[:n]) //1文字抽出
	//fmt.Println(size)
	if r == utf8.RuneError {
		return 0, errors.New("RuneError")
	}
	fmt.Println(s.r)
	s.r = io.MultiReader(bytes.NewReader(s.buf[size:n]), s.r)
	return r, nil
}
	
func  main() {
	
	s := NewRuneScanner(strings.NewReader("Hello, world")) //strings.NewReader: strings.NewReaderのポインタを生成する。
	for {
		r, err := s.Scan()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%c\n", r)
	}
}