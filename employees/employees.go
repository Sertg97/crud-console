package employee

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Employee struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Role    string `json:"role"`
}

// Read employees
func ReadEmployee(employees []Employee) {
	if len(employees) == 0 {
		fmt.Println("No employees found :(")
		return
	}

	/*for _, employee := range employees {

		status := " "
		if employee.Name {
			status = "✔"
		}

		fmt.Printf("[%s] [%d] [%s]\n", status, employee.Id, employee.Name)
	}*/
}

// Create employees
func CreateEmployee(employees []Employee, name string, surname string, age int, role string) []Employee {
	newEmployee := Employee{
		Id:      GetNextId(employees),
		Name:    name,
		Surname: surname,
		Age:     age,
		Role:    role,
	}
	return append(employees, newEmployee)
}

// Update employees
func UpdateEmployee() {
	file, err := os.Open("employees.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var employees []Employee
	err = json.NewDecoder(file).Decode(&employees)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func GetNextId(employees []Employee) int {
	if len(employees) == 0 {
		return 1
	}
	return employees[len(employees)-1].Id + 1
}

// Save employees
func SaveEmployee(file *os.File, employees []Employee) {
	bytes, err := json.Marshal(employees)
	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
