package main

func swap2(n, m *int){  //*intなのはn, mの値(10, 20)がintだから(10.5, 20.0)なら*float64になる
	print(*n)
	print(*m)
	*n = *m
	*m = 10
}

func main() {
	n, m := 10, 20
	swap2(&n, &m)
	print(&n)
	println(n, m)
}