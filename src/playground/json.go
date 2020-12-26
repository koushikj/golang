package main

import (
	"encoding/json"
	"fmt"
)

type emp struct {
	Name string `json:"Employee Name"`
	Dept string `json:"Department Name"`
	Id   int    `json:"Employee Id"`
}

func main() {
	jsonString := marshall()
	employeesData := unmarshall(jsonString)
	for _, v := range employeesData {
		fmt.Println(v.Dept, v.Id, v.Name)
	}

}

var err error

func marshall() string {

	e1 := emp{"david", "support", 10019}
	e2 := emp{"goggins", "helpdesk", 10211}
	employees := []emp{e1, e2}
	fmt.Println(employees)
	by, err := json.Marshal(employees)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(by)
	fmt.Println(string(by))
	return string(by)
}

func unmarshall(s string) []emp {
	var employees []emp
	fmt.Println(employees)

	err = json.Unmarshal([]byte(s), &employees)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(employees)
	return employees
}
