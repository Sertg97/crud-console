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
	}

}

func printUsage() {
	fmt.Println("--- Welcome CRUD console ---")
	fmt.Println("Usage: [create|read|update|delete]")
}
