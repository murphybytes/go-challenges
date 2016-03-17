package main

import "fmt"

func getDecent(n int) (fives, threes int, poss bool) {
	if n < 3 {
		return
	}
	r := n % 3
	if r == 0 {
		return n, 0, true
	}
	fives = n - r
	threes = r

	for {
		fives -= 3
		threes += 3
		if threes%5 == 0 {
			poss = true
			break
		}
		if fives <= 0 {
			break
		}

	}

	return
}

func main() {
	var tests int
	fmt.Scan(&tests)
	for i := 0; i < tests; i++ {
		var digits int
		fmt.Scan(&digits)
		fives, threes, poss := getDecent(digits)
		if poss {
			for j := 0; j < fives; j++ {
				fmt.Print("5")
			}
			for k := 0; k < threes; k++ {
				fmt.Print("3")
			}
		} else {
			fmt.Print("-1")
		}
		fmt.Println("")
	}

}
