package main
import "fmt"
func main() {
	var sum int
	sum = 5 + 6 + 3
	avg := float64(sum / 3)
	if avg > 4.5 {
		println("good")
	} else {
		println("false")
	}
	

	p := struct {
		name string
		age  int
	}{
		name: "Gopher",
		age:  10,
	}
	p.age++
	println(p.name, p.age)
	

	var ns2 = []int{1, 2, 3, 4, 5}
	println(ns2[3])
	println(len(ns2))
	fmt.Println(ns2[1:2])
	//ns3 := [...]int{1, 2, 3, 4, 5}
	
	/* array := make([]int, length, capacity) */
	ns5 := make([]int, 2, 10)
	println(ns5)
	fmt.Println(ns5)
	println(len(ns5))
	ns4 := []int{5: 10, 10:100}
	fmt.Println(ns4[5])
	fmt.Println("---")
	ns2 = append(ns2, 60, 70)
	
	//capacity
	println(cap(ns2))
	a := []int{10, 20}
	b := append(a, 30)
	a[0] = 100
	fmt.Println(b, cap(b))
	c := append(b, 40)
	b[1] = 2000
	fmt.Println(c, cap(c))
	
	ns := []int{10, 20, 30, 40, 50}
	n, m := 2, 4
	
	// get slice of nth or later
	fmt.Println(ns[n:])
	//get slice of until "m-1"nth 
	fmt.Println(ns[:m])


	/*
    // zero value is nill
 	var m map[string]int
	
	//initialization by make
	m = make(map[string]int)
	
	// specify capacity
	m = make(map[string]int, 10)
	
	// initialization by literal
	m = map[string]int{"x": 10, "y": 20}
	
	// case by nil
	m := map[string]int{}
 */
	m := map[string]int{"x": 10, "y": 20}
		
	println(m["x"])
	
	m["z"] = 30
	
	// n = value, ok = True or false
	n, ok := m["z"]
	
	println(n, ok)
	
	
	delete(m, "z")
	
	n, ok = m["z"]
	println(n, ok)
	

	// type "Type name" "Base type"
	
	// based on embedded type
	type MyInt int
	
	// based on type of other package
	/* type MyWriter io.Writer */
	
	// based on type riteral
	type Person struct {
		Name string
	}
	
	var n int = 100

	m := MyInt(n)
	n = int(m)
}
