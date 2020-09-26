package main
import (
	"strconv"
	"fmt"
)
/* interface: インターフェースで宣言されているメソッドが全て実装されている構造体ならばどんな型でも対象のインターフェースとして使用できる　*/

type Car interface {
	run(int) string
	stop()
}
/* run, stopメソッドを実装。 => MyCarはCarインターフェースとして扱うことができる */
type MyCar struct {
	name string
	speed int
}

func (u *MyCar) run(speed int) string {
	u.speed = speed
	return strconv.Itoa(speed) + "kmで走ります"
}


func (u *MyCar) stop() {
	fmt.Println("停止します")
	u.speed = 0
}

func main(){
	myCar := &MyCar{name: "マイカー", speed: 0}  //メソッドのレシーバがポインタ型なので、ポインタ型を使用する

	var i Car = myCar
	fmt.Println(i.run(50))
	i.stop()
}
