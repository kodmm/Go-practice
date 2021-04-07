package main 

import(
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Person struct {
		Name string `json:"name"`
		Age int `json:"age"`
	}

	p := &Person{Name: "tenntenn", Age: 31}

	var buf bytes.Buffer 
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	var p2 Person
	dec := json.NewDecoder(&buf)
	if err := dec.Decode(&p2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p2)
}