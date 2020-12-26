package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Println(x)
	fmt.Printf("from main %v, %T\n", x, x)
	fmt.Printf("from main %v, %T\n", &x, &x)
	foo2(&x)

	fmt.Printf("from main %v, %T\n", x, x)
	fmt.Printf("from main %v, %T\n", &x, &x) // note the same address but different value

}

func foo2(s *int) {
	*s = 1002
	fmt.Printf("from foo  %v, %v\n", s, *s)

}
