package main
import(
	"fmt"
	
)

type Employee struct {
	ID int
	FirstName string
	LastName string 
	Address string 
}

func main() {
	employee, err := getInformation(1001)
	if err != nil {
		fmt.Println("employee")
	}else {
		fmt.Println(employee)
	}
}

func getInformation(id int) (*Employee, error) {
	employee, err := apiCallEmployee(1000)
	return employee, err
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}