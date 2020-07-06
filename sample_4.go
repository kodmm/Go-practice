package main

//*int intのポイント型
func f(xp *int) {
	*xp = 100

}


func main() {
	var x int
	//&でポインタを取得する
	f(&x)
	println(x)
	
	
	ns := []int{10, 20, 30}
	ns2 := ns
	ns[1] = 200
	
	for i:=0; i < len(ns); i++ {
		println(ns[i])
	}
	for i:=0; i < len(ns2); i++ {
		print(ns[i])
	}
	
}