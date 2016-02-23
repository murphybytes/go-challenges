package main

import (
	"fmt"
	"math/big"
)

func fact(val int64, curr *big.Int) {

	val--
	nextval := big.NewInt(val)
	var product big.Int
	product.Mul(curr, nextval)

	if val == 0 {
		fmt.Println(curr.String())
		return
	}

	fact(val, &product)
}

func main() {
	var val int64
	fmt.Scan(&val)
	curr := big.NewInt(val)
	fact(val, curr)
}
