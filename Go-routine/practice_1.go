package main
import "fmt"

func main() {
	txtCh := make(chan string)
	
	defer close(txtCh)

	go task(txtCh)

	s := <- txtCh
	fmt.Println(s)
}

func task(txtCh chan <- string) {
	txtCh <- "I'm goroutine"
}

