package main

import "fmt"

// For matrix M of dimension n X n, let X = the sum of the diagonal [0,0]..[n-1, n-1]
// and Y the sum of the diagonal [0, n-1] to [n-1,0]. Print |X-Y| to STDOUT
// Usage: Enter the dimension of the matrix on the first line, then enter each
// following row on a seperate line, for example a 3x3 matrix is as follows
// 3
// 1 0 0
// 0 1 0
// 0 0 1

func sumdiags(m [][]int) (lr, rl int) {
	jj := len(m) - 1
	for i, row := range m {
		for j, val := range row {

			if i == j {
				lr += val
			}

			if jj == j {
				rl += val
				jj--
			}
		}
	}
	return

}

func main() {
	var dim int
	fmt.Scanln(&dim)
	//fmt.Printf("dim %d", dim)
	m := make([][]int, dim, dim)
	for row := 0; row < dim; row++ {

		m[row] = make([]int, dim, dim)
		r := make([]interface{}, dim, dim)
		for i := range m[row] {
			r[i] = &m[row][i]
		}
		fmt.Scanln(r...)
	}

	lr, rl := sumdiags(m)

	result := lr - rl
	if result < 0 {
		result *= -1
	}

	fmt.Printf("%d\n", result)
}
