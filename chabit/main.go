package main

// preallocate a, b, c
// skip addition operation for sequential get_c ops
// generate worse case file and profile
import "fmt"

const maxint = ^uint32(0)
const bigmask = uint64(maxint)

func setBit(idx int, x byte, num []uint32, cache []int64) {

	chunk := len(num) - idx/32 - 1
	shift := uint(idx % 32)
	mask := uint32(1 << shift)
	if num[chunk]&mask == mask {
		if x == 0 {
			cache[chunk] = -1
		}
	} else {
		if x == 1 {
			cache[chunk] = -1
		}
	}

	if x == 1 {
		num[chunk] |= mask
	} else {
		num[chunk] &= ^mask
	}

}

func getBit(idx int, num []uint32) byte {

	chunk := len(num) - idx/32 - 1
	shift := uint(idx % 32)
	mask := uint32(1 << shift)
	if (mask & num[chunk]) > uint32(0) {
		return 1
	}
	return 0
}

func add(a, b []uint32, cache []int64) []uint32 {
	var carry uint32
	//resultLength := len(a)*2 + 1
	res := make([]uint32, len(a)+1, len(a)+1)

	for i := len(a) - 1; i >= 0; i-- {
		if cache[i] != -1 {
			res[i+1] = uint32(cache[i])
			continue
		}

		var sum uint64

		sum = uint64(a[i] + b[i] + carry)

		if sum > uint64(maxint) {
			fmt.Println("carry")
			carry = uint32(sum >> 32)
			res[i+1] = uint32(sum & bigmask)
			cache[i] = int64(res[i+1])
			if i > 0 {
				cache[i-1] = -1
			}
		} else {
			carry = 0
			res[i+1] = uint32(sum)
			cache[i] = int64(res[i+1])
		}

	}

	res[0] = carry

	return res

}

func bytes2String(b []byte) string {
	cp := make([]byte, len(b), len(b))
	for i := 0; i < len(b); i++ {
		cp[i] = b[i] + 48
	}
	return string(cp)
}

func getSeqBounds(currByte int, s string) (right int, left int) {
	right = (currByte * 32) + len(s)%32
	if currByte > 0 {
		left = right - 32
	}
	return
}

func string2Ints(s string) []uint32 {
	padding := 0
	if len(s)%32 > 0 {
		padding = 1
	}

	sz := len(s)/32 + padding
	res := make([]uint32, sz, sz)
	for bt := sz - 1; bt >= 0; bt-- {
		right, left := getSeqBounds(bt, s)

		shift := uint(0)
		for i := right - 1; i >= left; i-- {
			if s[i] == '1' {
				res[bt] |= 1 << shift
			}
			shift++
		}

	}
	return res
}

type problemReader interface {
	getInitial() (q int, a, b []uint32)
	readOp() (op string, idx int, x byte)
}

type stdinReader struct {
	n, q int
	a, b []uint32
}

func newStdinReader() problemReader {

	r := stdinReader{}
	intfs := []interface{}{&r.n, &r.q}
	fmt.Scanln(intfs...)
	var s string
	fmt.Scan(&s)
	r.a = string2Ints(s)
	fmt.Scan(&s)
	r.b = string2Ints(s)

	return r
}

func (r stdinReader) getInitial() (q int, a, b []uint32) {
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

func run(reader problemReader) (result string) {
	q, a, b := reader.getInitial()
	cache := make([]int64, len(a), len(a))
	for k := range cache {
		cache[k] = -1
	}

	for i := 0; i < q; i++ {
		op, idx, x := reader.readOp()
		switch op {
		case "set_a":
			setBit(idx, x, a, cache)
		case "set_b":
			setBit(idx, x, b, cache)
		case "get_c":
			c := add(a, b, cache)
			result += fmt.Sprint(getBit(idx, c))
		}

	}

	return result
}

func main() {
	reader := newStdinReader()
	fmt.Println(run(reader))
}
