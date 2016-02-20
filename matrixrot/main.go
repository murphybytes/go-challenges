package main

import "fmt"

func rot(m [][]int, n *[][]int, step int) {
	if step > 1 {
		return
	}

	firstcol := 0 + step
	//	col := firstcol
	firstrow := 0 + step
	//	row := firstrow
	lastrow := len(m) - step - 1
	lastcol := len(m[0]) - step - 1

	fmt.Printf("fc %d fr %d lr %d lc %d\n", firstcol, firstrow, lastrow, lastcol)

	if firstcol >= lastcol || firstrow >= lastrow {
		return
	}

	//saved := m[firstrow][firstcol]

	//	var seed int
	// down the first col
	for i := firstrow + 1; i <= lastrow; i++ {
		fmt.Printf("(%d,%d)\n", i, firstcol)

		m[i][firstcol] = m[i-1][firstcol]

	}

	for i := firstcol + 1; i <= lastcol; i++ {
		fmt.Printf("(%d,%d)\n", lastrow, i)
		(*n)[lastrow][i] = m[lastrow][i-1]
		//    *n[lastrow, i] = m[lastrow,firstcol]
	}

	for i := lastrow - 1; i >= firstrow; i-- {
		fmt.Printf("(%d,%d)\n", i, lastcol)
		(*n)[i][lastcol] = m[i+1][lastcol]
	}

	for i := lastcol; i > firstcol; i-- {
		fmt.Printf("(%d,%d) val %d\n", firstrow, i, m[firstrow][i])
		(*n)[firstrow][i-1] = m[firstrow][i]
	}
	step++
	fmt.Println("-----------------------------")
	rot(m, n, step)

}

func main() {

	m := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	n := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	rot(m, &n, 0)

	for _, row := range n {
		for _, val := range row {
			fmt.Printf("%d, ", val)
		}
		fmt.Println("")
	}

	// n := [][]int{
	// 	{0, 0, 0, 0},
	// 	{0, 0, 0, 0},
	// 	{0, 0, 0, 0},
	// 	{0, 0, 0, 0},
	// 	{0, 0, 0, 0},
	// }
	//
	// fmt.Println("START N")
	// rot(n, nil, 0)
	//
	// o := [][]int{
	// 	{0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0},
	// }
	//
	// fmt.Println("START O")
	//
	// rot(o, nil, 0)

}
