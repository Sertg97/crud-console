package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	employee "github.com/setr4y/crud-console/employees"
)

func main() {
	file, err := os.OpenFile("employees.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var employees []employee.Employee

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		json.Unmarshal(bytes, &employees)
	} else {
		employees = []employee.Employee{}
	}

	if len(os.Args) < 2 {
		printUsage()
	}

	switch os.Args[1] {
	case "create":
		insert := bufio.NewReader(os.Stdin)

		fmt.Print("What's your name: ")
		name, _ := insert.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Print("What's your surname: ")
		surname, _ := insert.ReadString('\n')
		surname = strings.TrimSpace(surname)

		fmt.Print("How old are you: ")
		age, _ := insert.ReadString('\n')
		age = strings.TrimSpace(age)
		ageInt, _ := strconv.Atoi(age)

		fmt.Print("What's your role: ")
		role, _ := insert.ReadString('\n')
		role = strings.TrimSpace(role)

		employees = employee.CreateEmployee(employees, name, surname, ageInt, role)
		employee.SaveEmployee(file, employees)

	case "read":
		employee.ReadEmployee(employees)

	case "update":
		// Imprimir los datos existentes en la estructura
		fmt.Println("---Existing employee data---")
		for _, employee := range employees {
			fmt.Printf("%d\t%s\t%s\t%d\t%s\n", employee.Id, employee.Name, employee.Surname, employee.Age, employee.Role)
		}

		// Solicitar al usuario que ingrese los nuevos datos para actualizar
		var updateId int
		fmt.Println()
		fmt.Print("Enter the employee ID: ")
		fmt.Scan(&updateId)

		var newName, newSurname, updateRole string
		var newAge int

		fmt.Print("Enter the new name of the employee: ")
		fmt.Scan(&newName)

		fmt.Print("Enter the new surname of the employee: ")
		fmt.Scan(&newSurname)

		fmt.Print("Enter the new age of the employee: ")
		fmt.Scan(&newAge)

		fmt.Print("Enter the new role of the employee: ")
		fmt.Scan(&updateRole)

		// Actualizar los datos en la estructura
		for i, employee := range employees {
			if employee.Id == updateId {
				employees[i].Name = newName
				employees[i].Surname = newSurname
				employees[i].Age = newAge
				employees[i].Role = updateRole
				break
			}
		}

		file, err = os.Create("employees.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		err = json.NewEncoder(file).Encode(employees)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Employee data updated successfully.")

	case "delete":

	}

}

func SaveEmployee(file *os.File, employees []employee.Employee) {
	panic("unimplemented")
}

func printUsage() {
	fmt.Println("--- Welcome CRUD console ---")
	fmt.Println("Usage: [create|read|update|delete]")
}
