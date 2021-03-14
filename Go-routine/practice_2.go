package main
import (
	"time"
	"fmt"
)

func main() {
	txtCh := make(chan string, 10)

	defer close(txtCh)

	go task(txtCh)

	for i := 0; i < 10; i++ {
		txtCh <- fmt.Sprintf("push txt%d.", i)
	}
}

func task(txtCh <-chan string) {
	time.Sleep(1 * time.Second)

	<- txtCh
}