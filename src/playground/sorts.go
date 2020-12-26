package main

import (
	"fmt"
	"sort"
)

func main() {
	defaultSort()
	fmt.Println("--------------------")
	customSort()

}

type Person struct {
	name string
	age  int
}

/* create a interface for custom sort by age*/
type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].age < a[j].age

}

/*create a interface for custom sort by name*/
type ByName []Person

func (a ByName) Len() int {
	return len(a)
}
func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByName) Less(i, j int) bool {
	return a[i].name < a[j].name
}
func customSort() {
	var persons []Person
	p1 := Person{"david", 13}
	p2 := Person{"alex", 29}
	p3 := Person{"bob", 89}
	p4 := Person{"peter", 43}
	persons = append(persons, p1)
	persons = append(persons, p2)
	persons = append(persons, p3)
	persons = append(persons, p4)
	fmt.Println("original data :", persons)
	sort.Sort(ByAge(persons))
	fmt.Println("after sorting by age:", persons)
	sort.Sort(ByName(persons))
	fmt.Println("after sorting by name:", persons)

}

func defaultSort() {
	i := 043 // number begins with 0 so treated as base 8
	xi := []int{190, 2910, 210, 1, 1002, 10, 43, 32}
	xs := []string{"abc", "ABC", "DEF", "122", "spw", "pqer"}
	fmt.Println(xi)
	fmt.Println(xs)
	fmt.Println(i)

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)
}
