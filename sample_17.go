package main
import "fmt"

type Person interface {
	Title() string
	Name() string
}

type person struct {
	firstName string 
	lastName string
}

func (p *person) Name() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

type Gender int 

const (
	Female = iota
	Male
)
//Goでは小文字で始まる型は(female, male)そのパッケージからしか使用できない
type female struct {
	*person 
}

func (f *female) Title() string {
	return "Ms."
}

type male struct {
	*person 
}

func (m *male) Title() string {
	return "Mr."
}

func NewPerson(gender Gender, firstName, lastName string) Person {
	p := &person{firstName, lastName}

	if gender == Female {
		return &female{p}
	}else{
		return &male{p} 
	}
	
}

func main() {
	var human Person = NewPerson(0, "sato", "daiki")
	gender := human.Title()
	full_name := human.Name()
	fmt.Printf("性別: %s%s", gender, full_name)
}