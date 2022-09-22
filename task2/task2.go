package task2

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}
type company struct {
	companyName string
	employees   []employee
}

func printCompany() {
	emp1 := employee{"Amir", 80000, "Full-Stack Developer"}
	emp2 := employee{"Mehmood", 90000, "Backend Developer"}
	emp3 := employee{"Faiez", 60000, "Frontend Developer"}
	emplys := []employee{emp1, emp2, emp3}
	comp := company{"Tetra", emplys}
	fmt.Println("=====================")
	fmt.Println("Company : ", comp.companyName)
	fmt.Println("=====================")
	fmt.Println("Employees : ")
	for _, emp := range comp.employees {
		fmt.Println("Name : ", emp.name)
		fmt.Println("Salary : ", emp.salary)
		fmt.Println("Position : ", emp.position)

		fmt.Println("=====================")
	}
}
