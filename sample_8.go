package main
import(
	"fmt"
	"os"
)



func main() {
	msg := "111"
	defer fmt.Println(msg)
	msg = "world"
	defer fmt.Println(msg)
	fmt.Println("hello")
	// defer: 関数終了時に実行される
	//        スタック形式で実行される(LIFO)

}