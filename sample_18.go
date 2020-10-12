package main
import "fmt"


type Person struct {
	fmt.Stringer
	FirstName string 
	LastName string
	Age int
}

type StringerFunc func() string  //文字列を返す関数にfmt.Stringerインタフェースを実装させるための型

func (sf StringerFunc) String() string {
	return sf()
}

func BindStringer(p *Person, f func(p *Person) string) fmt.Stringer {
	return StringerFunc(func() string {
		return f(p)
	})
}

func NewPerson(firstName, lastName string, age int) (p *Person) {
	p = &Person{
		nil,
		firstName,
		lastName,
		age,
	}

	p.stringer = StringerFunc(func() string {
		return fmt.Sprintf("%s %s (%d)", p.FirstName, p.LastName, p.Age)
	})
	return
}

func (p *Person) SetStringer(sf func(p *Person) string) {
	p.stringer = StringerFunc(func() string {
		return sf(p)
	})
}

type stringer fmt.Stringer

func main(){
	p := Person{
		stringer: nil,
		FirstName: "Takahiro"
		LastName: "Nishijima"
		Age: 30,
	}

	fmt.Println(p.stringer)
	// p.Stringer = ??? //fmt.Stringerインターフェースを実装していれば代入できる
}