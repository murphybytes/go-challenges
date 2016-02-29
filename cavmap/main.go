package main

import "fmt"

func main() {
	var dim int
	fmt.Scan(&dim)
	m := make([][]int, 0, dim)
	for r := 0; r < dim; r++ {
		var inp string
		fmt.Scan(&inp)
		row := make([]int, 0, dim)
		for c := 0; c < len(inp); c++ {
			val := int(inp[c] - 48)
			row = append(row, val)
		}
		m = append(m, row)

	}

	for r := range m {
		for c := range m[r] {
			if r == 0 || r == (dim-1) {
				fmt.Print(m[r][c])
				continue
			}
			if c == 0 || c == (dim-1) {
				fmt.Print(m[r][c])
				continue
			}

			if m[r-1][c] < m[r][c] && m[r][c+1] < m[r][c] && m[r+1][c] < m[r][c] && m[r][c-1] < m[r][c] {
				fmt.Print("X")
			} else {
				fmt.Print(m[r][c])
			}

		}
		fmt.Println("")
	}

}
