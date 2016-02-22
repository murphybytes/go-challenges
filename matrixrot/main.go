package main

/*
Matix rotation
Reads input on stdin, first line is rows, columns and rotation moves
Subsequent lines are each row of the matrix
Example
4 4 2
1 2 3 4
5 6 7 8
9 10 11 12
13 14 15 16

Output
3 4 8 12
2 11 10 16
1 7 6 15
5 9 13 14
*/

import "fmt"

func rotate(m *[][]int, rows, cols, moves int) {
	for layer := 0; ; layer++ {
		currrows := rows - (2 * layer)
		currcols := cols - (2 * layer)

		if currrows <= 0 || currcols <= 0 {
			break
		}

		transform := getTransform(*m, currrows, currcols, layer, moves)
		applyTransform(m, currrows, currcols, layer, transform)

	}
}

func getTransform(m [][]int, currrows, currcols, layer, moves int) []int {

	elements := (2 * currrows) + (2 * (currcols - 2))
	partition := make([]int, elements, elements)
	partitionMoves := moves % elements
	partitionIndex := 0

	row := 0 + layer
	col := 0 + layer

	for ; row < currrows+layer; row++ {
		insertion := (partitionIndex + partitionMoves) % elements

		partition[insertion] = m[row][col]
		partitionIndex++

	}

	row = layer + currrows - 1
	col = layer + 1

	for ; col < currcols+layer; col++ {
		insertion := (partitionIndex + partitionMoves) % elements
		partition[insertion] = m[row][col]
		partitionIndex++
	}

	col = currcols + layer - 1
	row = currrows + layer - 2

	for ; row >= layer; row-- {
		insertion := (partitionIndex + partitionMoves) % elements
		partition[insertion] = m[row][col]
		partitionIndex++

	}

	row = layer
	col = currcols + layer - 2

	for ; col > layer; col-- {
		insertion := (partitionIndex + partitionMoves) % elements
		partition[insertion] = m[row][col]
		partitionIndex++
	}

	return partition
}

func applyTransform(m *[][]int, currrows, currcols, layer int, transform []int) {

	curr := 0

	row := 0 + layer
	col := 0 + layer

	fmt.Printf("row %d col %d\n", row, col)

	for ; row < currrows+layer; row++ {
		(*m)[row][col] = transform[curr]
		curr++
	}

	row = layer + currrows - 1
	col = layer + 1

	for ; col < currcols+layer; col++ {
		(*m)[row][col] = transform[curr]
		curr++
	}

	col = currcols + layer - 1
	row = currrows + layer - 2

	for ; row >= layer; row-- {
		(*m)[row][col] = transform[curr]
		curr++
	}

	row = layer
	col = currcols + layer - 2

	for ; col > layer; col-- {
		(*m)[row][col] = transform[curr]
		curr++
	}

}

func main() {
	dims := make([]int, 3, 3)
	intf := make([]interface{}, 3, 3)
	for i := range intf {
		intf[i] = &dims[i]
	}
	fmt.Scanln(intf...)
	rows := dims[0]
	cols := dims[1]
	rot := dims[2]

	m := make([][]int, 0, rows)

	for i := 0; i < rows; i++ {
		row := make([]int, cols, cols)
		intf := make([]interface{}, cols, cols)
		for j := range intf {
			intf[j] = &row[j]
		}
		fmt.Scanln(intf...)
		m = append(m, row)
	}

	rotate(&m, rows, cols, rot)

	for row := range m {
		for col := range m[row] {
			if col > 0 {
				fmt.Print(" ")
			}
			fmt.Print(m[row][col])
		}
		fmt.Println("")
	}

}
