package main

import "fmt"

func main() {

	var mask int = 4294967295

	r := 1 ^ mask
	fmt.Println(r)

}
