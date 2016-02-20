package main

import "fmt"

func fib(first, second, curr, end int) int64 {
	if curr < end {
		curr++
		newsecond := first + second

		return fib(second, newsecond, curr, end)
	}

	return first

}

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(fib(0, 1, 0, number))

}
