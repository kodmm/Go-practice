package main

func swap2(n, m *int){
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