package main

import "fmt"

func newMatrix(r int) []string {
	grid := make([]string, 0, r)
	for j := 0; j < r; j++ {
		var s string
		fmt.Scan(&s)
		grid = append(grid, s)
	}
	return grid
}

func findMatches(grid, pat string) []int {
	result := []int{}
	g := []byte(grid)
	p := []byte(pat)
	for i := 0; i < len(g); i++ {
		matches := 0
		k := i
		for j := 0; j < len(p) && k < len(g); j++ {
			if g[k] != p[j] {
				break
			}
			matches++
			if matches == len(pat) {
				result = append(result, i)
				i = k
				break
			}
			k++
		}

	}
	return result
}

func test(grd, pat []string) bool {
	for i := 0; i < len(grd); i++ {
		hits := findMatches(grd[i], pat[0])
		if len(hits) == 0 {
			continue
		}
		for h := 0; h < len(hits); h++ {

			p := 1
			g := i + 1
			for ; p < len(pat) && g < len(grd); p++ {
				if !matchAt(hits[h], grd[g], pat[p]) {
					break
				}

				g++
			}
			if p == len(pat) {
				return true
			}
		}
	}
	return false
}

func matchAt(idx int, grid, pat string) bool {
	g := []byte(grid)
	p := []byte(pat)
	i := idx
	for j := 0; j < len(p) && i < len(g); j++ {
		if g[i] != p[j] {
			return false
		}
		i++
	}
	if i < len(p) {
		return false
	}
	return true
}

func main() {
	var testCases int
	fmt.Scan(&testCases)
	for i := 0; i < testCases; i++ {
		var rowsGrid, colsGrid int
		fmt.Scanln(&rowsGrid, &colsGrid)
		grid := newMatrix(rowsGrid)
		var rowsPat, colsPat int
		fmt.Scanln(&rowsPat, &colsPat)
		pat := newMatrix(rowsPat)

		if test(grid, pat) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
