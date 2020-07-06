package main

func check(i int) bool{
	return i%2 == 0
}

func main() {
	for i:=0; i <= 100; i++ {
		print(i)
		if check(i){
			println(" -偶数")
		}else{
			println(" -奇数")
		}
	}
}