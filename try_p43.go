package main
import(
	"fmt"
	"math/rand"
	"time"
)

func main() {
	   t := time.Now().UnixNano()
       rand.Seed(t)
       s := rand.Intn(6)
       
       switch s + 1 {
			case 6:
			fmt.Printf("%d -  大吉", s)
			case 5, 4:
				fmt.Printf("%d - 中吉", s)
			case 3, 2:
				fmt.Printf("%d - 吉", s)
			default:
				fmt.Printf("%d - 凶", s)
		}
}
