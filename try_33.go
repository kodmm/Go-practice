package main

func main() {
	n1 := []int{19, 86, 1, 12}

	var sum int
	for i := 0; i < len(n1); i++ {
		sum += n1[i]
	}
	println(sum)
}	
