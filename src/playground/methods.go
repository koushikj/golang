package main

import (
	"fmt"
)

type employee struct {
	name string
	id int
	dept string
}

/* attach method to the struct */
func (s employee) printEmpDetails() string  {
	return fmt.Sprint(s.name,"\t",s.dept,"\t",s.id)
}

func main()  {
	e1:= employee{
		name:"abc",
		id:12,
		dept: "sw",
	}
	e2 := employee{
		"def",
			21,
			"hw",
	}
	fmt.Println(e1.printEmpDetails())
	fmt.Println(e2.printEmpDetails())
}
