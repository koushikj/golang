package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	strs := strings.Split(scanner.Text(), " ")
	scanner.Scan()
	queryCount, _ := strconv.Atoi(scanner.Text())
	queries := getQueries(queryCount, scanner)

	//fmt.Println(strs)
	//fmt.Printf("%T \n",strs)
	nos := convertStringSliceToIntSlice(strs)
	//fmt.Printf("%T \n",nos)
	sort.Ints(nos)
	//fmt.Println(nos)
	//fmt.Println(queries)

	for _, v := range queries {
		//fmt.Println("OP",v)
		fmt.Println(binarySearch(0, len(nos)-1, v, nos))
		//result, _ := binarySearch2(nos, v)
		//fmt.Printf("the number was found in position %d after %d steps.\n\n", result, searchCount)
		//fmt.Println(result+1)
	}

}
func binarySearch2(a []int, search int) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid] > search:
		result, searchCount = binarySearch2(a[:mid], search)
	case a[mid] < search:
		result, searchCount = binarySearch2(a[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // found
	}
	searchCount++
	return
}

func binarySearch(low, high, query int, nos []int) int {
	for low <= high {
		//mid := (low+(high-low))/2
		mid := low + ((high - low) / 2)
		//fmt.Println("low high mid query nos",low,high,mid,query,nos)
		if nos[mid] < query {
			low = mid + 1
		} else if nos[mid] > query {
			high = mid - 1
		} else {
			return mid + 1
		}
	}
	return -1
}

func getQueries(n int, scanner *bufio.Scanner) []int {
	var queries []int
	for i := 0; i < n; i++ {
		scanner.Scan()
		j, _ := strconv.Atoi(scanner.Text())
		queries = append(queries, j)
	}
	return queries
}
func convertStringSliceToIntSlice(x []string) []int {
	var nos []int
	for _, v := range x {
		i, _ := strconv.Atoi(v)
		nos = append(nos, i)
	}
	return nos
}
