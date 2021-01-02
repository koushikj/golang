package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	find the position of element in a given array using linear search
	INPUT :
	7 3
	1 2 3 4 5 6 3

	OUTPUT:
	7

*/
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	lineOne := createIntSlice(parts)

	scanner.Scan()
	parts = strings.Split(scanner.Text(), " ")
	lineTwo := createIntSlice(parts)

	//fmt.Println(lineOne)
	//fmt.Println(lineTwo)

	fmt.Println("position of", lineOne[1], "from first =", findPositionOfElement(lineOne, lineTwo))
	fmt.Println("position of", lineOne[1], "from last  =", findPositionOfElementFromTheLast(lineOne, lineTwo))

}
func findPositionOfElementFromTheLast(a, b []int) int {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == a[1] {
			return i + 1
		}
	}
	return -1
}

func findPositionOfElement(a, b []int) int {
	for i, v := range b {
		if v == a[1] {
			return i + 1
		}
	}
	return -1
}

func createIntSlice(nums []string) []int {
	var r []int
	for _, v := range nums {
		i, _ := strconv.Atoi(v)
		r = append(r, i)
	}
	return r
}
