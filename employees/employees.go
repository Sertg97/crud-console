package employee

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Employee struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Role    string `json:"role"`
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

// Read employees
func ReadEmployee(employees []Employee) {
	data, err := ioutil.ReadFile("employees.json")
	if err != nil {
		fmt.Println("Error al leer el archivo JSON:", err)
		return
	}

	err = json.Unmarshal(data, &employees)
	if err != nil {
		fmt.Println("Error al decodificar el archivo JSON:", err)
		return
	}
}

// The ID is self-incremented when new employees are created
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

// Delete employees
func DeleteEmployee() {
	/*for i, employee := range employees {
		if employee.Id == id {
			employees = append(employees[:i], employees[i+1:]...)
		}
	}
	return employees*/

	// Abrir el archivo JSON existente para leer
	file, err := os.Open("employees.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Leer el contenido del archivo JSON en memoria
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Decodificar los datos JSON en una estructura Go
	var employees []Employee
	err = json.Unmarshal(data, &employees)
	if err != nil {
		fmt.Println(err)
		return
	}
}
