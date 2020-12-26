package main

import "fmt"

func main() {
	defer closeConnection()
	foo()
	boo("bond")
	s := monkey("banana")
	fmt.Println(s)
	st, no, b := ping("pong")
	fmt.Println(st)
	fmt.Println(no)
	fmt.Println(b)
	fmt.Println("total :", sum1(1, 120, 1292, 1921, 192201, 12921, 19291, 0))
	xi := []int{112, 12, 121, 11, 11, 101, 122}
	fmt.Println("total :", sum("hello", xi...))
	fmt.Println("total :", sum("hi"))

	/* anonymous func*/
	func() {
		fmt.Println("im a anonymous function")
	}()

	func(x int, s string, ix ...int) {
		fmt.Println("im a anonymous function with arguments :", x, s, ix)
	}(1920102, "goLang", 1020, 29102, 12910)

	/* func expression*/
	f1 := func(op1 ...float32) {
		fmt.Println("im a functione and assigned to a variable")
		for _, v := range op1 {
			fmt.Println(v)
		}
	}
	f1(102, 101.20102, 1020301, 1202021.120201, 10201, 00001.1001, 1)

	/* functions that returns a func*/
	i1 := returnMeAFunctionThatReturnsInt()
	fmt.Printf("%T\n", i1)
	fmt.Println(i1())
	fmt.Println(returnMeAFunctionThatReturnsInt()()) // call a func which returns a func and call that returned func

	fmt.Println(total(1, 2, 3, 4, 5, 6, 7, 8, 9))
	/* call back functions */
	fmt.Println(even(total, 2, 3, 4, 5, 6, 7, 8, 9))

	// fact(5) = 5*4*3*2*1

	/* factorial using recursion */
	fmt.Println("fact(5)=", fact(5))
	/* factorial using loops */
	fmt.Println(factUsingLoop(5))

}

func fact(s int) int {
	if s == 1 {
		return 1
	}
	return s * fact(s-1)
}

func factUsingLoop(s int) int {
	result := 1
	for ; s > 0; s-- {
		result = result * s
	}
	return result
}
func foo() {
	fmt.Println("hi from foo")
}

func boo(s string) {
	fmt.Println("hello from boo", s)
}

func monkey(s string) string {
	return fmt.Sprint("hi from monkey ", s)
}

func ping(s string) (string, int, bool) {
	s1 := fmt.Sprint("hello ", s, " from ping")
	return s1, 121, false
}

func sum1(s ...int) int {
	return sum("", s...)
}

// variadic parameter - accepts 0 to many of certain type
func sum(s1 string, s ...int) int {
	fmt.Printf("%T \n", s)
	fmt.Println(s1)
	sum := 0
	for _, v := range s {
		sum = sum + v
		fmt.Println("\tadd ", v, "to sum, which is now", sum)
	}
	return sum
}

func returnMeAFunctionThatReturnsInt() func() int {
	return func() int {
		return 100
	}
}
func closeConnection() {
	fmt.Println("im a deferred function call - executes at the end of this program execution, irrespecive of where i called")
}

func total(s ...int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func even(f func(xi ...int) int, s ...int) int {
	var y []int
	for _, v := range s {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	//fmt.Println(f(y...))
	return f(y...)
}
