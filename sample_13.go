package main
import "fmt"

type Hex int

func (h Hex) String() string{
	return fmt.Sprintf("%x", int(h))
}

func main() {
	/*type Stringer interface {
		String() string
	}

	var s Stringer = Hex(100)
	fmt.Println(s.String())
	*/
	//interface{}型　 どんな型でも格納できる特殊な型・型チェックや型変換などで使える。
	var x, y interface{}

	print("%#v", x) // -> nil
	x = 1
	x = 2.1
	y = []int{1, 2, 3}
	fmt.Printf("%#v", y)
	y = "hello"
	y = 2
	fmt.Printf("%#v\n", y)
	z := [...]string{"go", "java", "python", "ruby", "rust"}
	fmt.Printf("%#v\n", z)  // fmt.Printf: フォーマット指定可。 %#v: Go言語の構文で表現する
	// [5]string{"go", "java", "python", "ruby", "rust"}%

	// z := x * y 演算不可

	//型チェック
	receive := 3
	var a interface{}

	//型チェック
	a = receive 
	if xi, ok := a.(int); ok{
		fmt.Println( xi, ok)
		fmt.Println( xi * xi)
	}

	//型チェックによる分岐
	switch b := a.(type){
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("other")
		fmt.Println(b)
	}

}
