package main
import "fmt"


func swap(x, y int) (x2, y2 int) {
	y2, x2 = x, y
	return 
}

func main() {
	

	x, y := swap(10, 20)
	_, a := swap(10, 20)
	x, y = y, x
	println(x,y)
	println(a)
	
	msg  := "Hello world"
	func() {
		println(msg)
	}()
	
	fs := make([]func() string, 2)
	fs[0] = func() string {return "hoge"}
	fs[1] = func() string {return "fuga"}
	for _, f := range fs {
		fmt.Println(f())
	}
	
}
