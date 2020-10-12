package main
import (
	"fmt"
)


type Hoge struct {
	n int
}

func (this *Hoge) Do() {
	fmt.Println(this.n)
}

type Fuga struct {
	Hoge
	n int
}

func main() {
	f := &Fuga{Hoge: Hoge{n: 100}, n: 200}
	f.Do()
	f.Hoge.Do()
}