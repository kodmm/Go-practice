package main
import "fmt"
//main packageはmain関数を定義することでエントリポイント(プログラム実行時の処理開始位置)
func main() {
	msg := "Hello, world"
	println(msg)

	const n = 1000000000000 / 1000000000000
	var m = n
	println(m)

	const message = "Hello world"
	println(message)

	const (
		a = iota
		b
	)
	const (
		c = 1 << iota
		d
		e
	)

	hoge := 100 + 200
	
	msg = "hoge" + "fuga"
	if hoge == 1 {
		println("xは1")
	} else if hoge == 2 {
		println("xは2")
	} else {
		println("xは1でも2でもない")
	}
	fmt.Println(a, b, c, d, e)

	switch a {
	case 1, 2:
		fmt.Println("a is 1 or 2")
	default:
		fmt.Println("default")
	}

	switch {
	case a == 1:
		fmt.Println("a is 1")
	default:
		fmt.Println("default")
	}

	for i := 0; i <= 100; i++ {

	}

	var i int
	for {
		if i%2 == 1 {
			break
		}
		i++
	}

	/* var ko int
	LOOP:
	for {
		switch {
		case ko%2 == 1:
			break LOOP
		}
		i++
	}
	*/

	// 繰り返しはforしかない
	for i := 0; i <= 3; i++ {
		println(i)
	}

	for i <= 5 {
		print(i)
		i++
	}
}