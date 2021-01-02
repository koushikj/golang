package main

import (
	"fmt"
)

// go is always pass by value
func main() {
	x := 10
	fmt.Println(x)
	fmt.Printf("from main - orignal value %v, %T\n", x, x)
	fmt.Printf("from main - original address %v, %T\n", &x, &x)

	foo1(x) // pass by value
	fmt.Printf("from main - after foo1 %v, %T\n", x, x)
	fmt.Printf("from main - after foo1 %v, %T\n", &x, &x) // note the same address but different value

	foo2(&x)                                              // pass by value - and the value is an address
	fmt.Printf("from main - after foo2 %v, %T\n", x, x)   // x has modified
	fmt.Printf("from main - after foo2 %v, %T\n", &x, &x) // note the same address but different value

}
func foo1(x int) {
	x = 1002
	fmt.Printf("from foo1 - %v, %v\n", x, &x)
}
func foo2(s *int) {
	*s = 1002
	fmt.Printf("from foo2 - %v, %v, %v\n", s, *s, &s)

}
