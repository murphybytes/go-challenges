package main

// preallocate a, b, c
// skip addition operation for sequential get_c ops
// generate worse case file and profile
import "fmt"

const maxint = ^uint32(0)

func setBit(idx int, x byte, num *[]byte) {
	i := len(*num) - 1 - idx
	(*num)[i] = x
}

func getBit(idx int, num []byte) byte {
	i := len(num) - 1 - idx
	return num[i]
}

func add(a, b []byte) []byte {
	var carry byte
	//resultLength := len(a)*2 + 1
	res := make([]byte, len(a)+1, len(a)+1)
	for i := len(a) - 1; i >= 0; i-- {
		j := a[i] + b[i]
		if j == 2 {
			if carry == 1 {
				res[i+1] = 1
			} else {
				res[i+1] = 0
			}
			carry = 1
			continue
		}
		if j == 1 {
			if carry == 1 {
				res[i] = 0
			} else {
				res[i+1] = 1
			}
			continue
		}
		if j == 0 {
			res[i+1] = carry
			carry = 0
		}

	}
	res[0] = carry
	return res
}
func bytes2String(b []byte) string {
	for i := 0; i < len(b); i++ {
		b[i] = b[i] + 48
	}
	return string(b)
}

func string2Bytes(s string) []byte {
	buff := []byte(s)
	for i := 0; i < len(buff); i++ {
		buff[i] = buff[i] - 48
	}

	return buff
}

type problemReader interface {
	getInitial() (q int, a, b []byte)
	readOp() (op string, idx int, x byte)
}

type stdinReader struct {
	n, q int
	a, b []byte
}

func newStdinReader() problemReader {

	r := stdinReader{}
	intfs := []interface{}{&r.n, &r.q}
	fmt.Scanln(intfs...)
	var s string
	fmt.Scan(&s)
	r.a = string2Bytes(s)
	fmt.Scan(&s)
	r.b = string2Bytes(s)

	return r
}

func (r stdinReader) getInitial() (q int, a, b []byte) {
	return r.q, r.a, r.b
}

func (r stdinReader) readOp() (op string, idx int, x byte) {
	fmt.Scan(&op)
	fmt.Scan(&idx)
	if op != "get_c" {
		fmt.Scan(&x)
	}
	return
}

func run(reader problemReader) {
	q, a, b := reader.getInitial()
	var result string
	for i := 0; i < q; i++ {
		op, idx, x := reader.readOp()
		switch op {
		case "set_a":
			setBit(idx, x, &a)
		case "set_b":
			setBit(idx, x, &b)
		case "get_c":
			c := add(a, b)
			result += fmt.Sprint(getBit(idx, c))
		}
	}
	fmt.Println(result)
}

func main() {
	reader := newStdinReader()
	run(reader)
}
