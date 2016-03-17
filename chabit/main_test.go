package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"
)

func checkExpectedInt(exp, act uint32, t *testing.T) {
	if exp != act {
		t.Error("Expected ", exp, " Actual ", act)
	}
}

type fileReader struct {
	n, q int
	a, b []uint32
	f    *os.File
	s    *bufio.Scanner
}

func (fr fileReader) getInitial() (q int, a, b []uint32) {
	return fr.q, fr.a, fr.b
}

func (fr fileReader) readOp() (op string, idx int, x byte) {
	fr.s.Scan()
	parts := strings.Split(fr.s.Text(), " ")
	op = parts[0]
	idx, _ = strconv.Atoi(parts[1])
	if op != "get_c" {
		i, _ := strconv.Atoi(parts[2])
		x = byte(i)
	}
	return
}

func (fr fileReader) close() {
	fr.f.Close()
}

func newFileReader(path string) (problemReader, error) {
	reader := fileReader{}
	var e error
	if reader.f, e = os.Open(path); e == nil {

		reader.s = bufio.NewScanner(reader.f)
		reader.s.Scan()
		parts := strings.Split(reader.s.Text(), " ")
		reader.n, _ = strconv.Atoi(parts[0])
		reader.q, _ = strconv.Atoi(parts[1])
		reader.s.Scan()
		reader.a = string2Ints(reader.s.Text())
		reader.s.Scan()
		reader.b = string2Ints(reader.s.Text())

	} else {
		return nil, e
	}
	return reader, e
}

var projectRoot string

func init() {
	projectRoot = os.Getenv("PROJECT_ROOT")
	if projectRoot == "" {
		fmt.Println("Missing PROJECT_ROOT environment variable. Set PROJECT_ROOT to the location of test**.txt files")
		os.Exit(1)
	}
}

func TestEndToEnd(t *testing.T) {
	path := projectRoot + "/test1.txt"
	if p, err := newFileReader(path); err == nil {
		defer p.(fileReader).close()
		result := run(p)
		if result != "1001111001011000000101110011000001" {
			t.Error("Expected 1001111001011000000101110011000001 got ", result)
		}
	} else {
		t.Error(err)
	}

}

func TestEndToEnd5(t *testing.T) {
	actual := "10111011111000100011100001010100001000011110111011111001010001100111111110011111010010110100000011110101111101110111110110101011000100001010101000110011010110111000110110000111111010011101011001111000011010111100010101111001000000110011100100110000100100000101011011100011011111101100110011011000001001011110111001111001110100001001101110011101100000001001010111010101101001010110000011111100000011111010100111111010001101010011010111011111011111011011000110010011000111011000010010111101111011010101110000001000011000001100010110010111011101101001000000110100010111110111011010100000000011101011110000011110000110100010111111111110101100001111110101001101110100101110111101000000010101100101100100011010010101001000100111100110100100010010000111111010110101010010111011000110111101011010100001010001111001000000111110110011110101000111111000010111100001010110000010100010010010100101111100101110110000011010000111001111001101010011000110010111001000110001011111110000001011101101001100101010001001100101010001101010111100110011000100000101100111111011011111110010110001111101011010011111011001100110001001100111011000101110110000110100110010100111111100010111101001111101110101010011100111110101111001001011100111100101110010001001111111100000001100000111100011011100101001110110101000011000100100001010101100001010010110010010100010100101111100101100110100011101111100000000010100010111111100011100110110001100100011000001111111010010111100110101110101110001001111110000100111000011101010011100000010010001101101100010010101010110100100110111001101101010011000110011111011000011101100100101100101011101011101001111101110100010111101001001101001011000000110110111111"
	path := projectRoot + "/test5.txt"
	if p, err := newFileReader(path); err == nil {
		defer p.(fileReader).close()
		result := run(p)
		if result != actual {
			t.Error("Fail")
		}
	}
}

func TestEndToEnd2(t *testing.T) {
	path := projectRoot + "/test2.txt"
	if p, err := newFileReader(path); err == nil {
		defer p.(fileReader).close()
		result := run(p)
		if result != "11" {
			t.Error("Expected 11 got ", result)
		}
	} else {
		t.Error(err)
	}

}

func BenchmarkRun(b *testing.B) {
	path := projectRoot + "/test3.txt"
	for i := 0; i < b.N; i++ {
		if p, err := newFileReader(path); err == nil {
			defer p.(fileReader).close()
			run(p)

		} else {
			fmt.Println(err)
		}
	}

}

// func TestSetBitAtIdx(t *testing.T) {
// 	b := []byte("000000000000000")
// 	setBit(1, 1, &b)
// 	exp := getBit(1, b)
// 	if exp != 1 {
// 		t.Error("Incorrect result")
// 	}
// 	setBit(1, 0, &b)
// 	exp = getBit(1, b)
// 	if exp != 0 {
// 		t.Error("Error")
// 	}
//
// }
//
// func TestAdd(t *testing.T) {
// 	a := string2Bytes("0001")
// 	b := string2Bytes("0001")
// 	c := add(a, b)
// 	x := getBit(1, c)
// 	if x != 1 {
// 		t.Error("Error ", x)
// 	}
//
// 	a = string2Bytes("0011")
// 	b = string2Bytes("0011")
// 	c = add(a, b)
// 	s := bytes2String(c)
// 	if s != "00110" {
// 		t.Error("Error ", s)
// 	}
//
// 	addBit(&c, 0, 1)
// 	s = bytes2String(c)
// 	if s != "00111" {
// 		t.Error("Error expected 00111 got ", s)
// 	}
//
// 	addBit(&c, 0, 0)
// 	s = bytes2String(c)
// 	if s != "00110" {
// 		t.Error("Expected 00110 got ", s)
// 	}
//
// 	a = string2Bytes("1111")
// 	b = string2Bytes("1111")
// 	c = add(a, b)
// 	fmt.Printf("%q\n", c)
// 	addBit(&c, 1, 1)
// 	fmt.Printf("%q\n", c)
// 	s = bytes2String(c)
// 	if s != "10010" {
// 		t.Error("expected 10010 got ", s)
// 	}
//
// 	a = string2Bytes("00011")
// 	b = string2Bytes("10010")
// 	c = add(a, b)
// 	s = bytes2String(c)
// 	if s != "010101" {
// 		t.Error("Error ", s)
// 	}
// 	a = string2Bytes("011111")
// 	b = string2Bytes("000001")
// 	c = add(a, b)
// 	s = bytes2String(c)
// 	if s != "0100000" {
// 		t.Error("Error ", s)
// 	}
//
// 	a = string2Bytes("11111")
// 	b = string2Bytes("11111")
// 	c = add(a, b)
// 	s = bytes2String(c)
// 	if s != "111110" {
// 		t.Error("Error ", s)
// 	}
//
// }
//

func TestSetBit(t *testing.T) {
	a := string2Ints("0")
	cache := make([]int64, 1000, 1000)
	setBit(1, 1, a, cache)
	checkExpectedInt(uint32(1), uint32(len(a)), t)
	checkExpectedInt(uint32(2), a[0], t)
	if getBit(1, a) != 1 {
		t.Error("Expected 1")
	}
	if getBit(2, a) != 0 {
		t.Error("Expected 0")
	}
	if getBit(0, a) != 0 {
		t.Error("expected 0")
	}
	bi := new(big.Int)
	bi.SetUint64(uint64(maxint) + 4)
	a = string2Ints(bi.Text(2))
	fmt.Println(bi.Text(2))
	fmt.Printf("%b\n", a[0])
	fmt.Printf("%b\n", a[1])
	if getBit(32, a) != 1 {
		t.Error("expected 1")
	}
	if getBit(33, a) != 0 {
		t.Error("Expected 0")
	}
	if getBit(63, a) != 0 {
		t.Error("Expected 0")
	}
	if getBit(0, a) != 1 {
		t.Error("expected 1")
	}
	setBit(0, 0, a, cache)
	if getBit(0, a) != 0 {
		t.Error("Expected 0")
	}
	setBit(33, 1, a, cache)
	if getBit(33, a) != 1 {
		t.Error("Expected 1")
	}

}
func TestString2Bytes(t *testing.T) {
	bi := new(big.Int)
	a := string2Ints("00010")
	checkExpectedInt(uint32(1), uint32(len(a)), t)
	checkExpectedInt(uint32(2), a[0], t)

	a = string2Ints("0")
	checkExpectedInt(uint32(1), uint32(len(a)), t)
	checkExpectedInt(uint32(0), a[0], t)

	a = string2Ints("01100001001")
	bi.SetString("01100001001", 2)
	checkExpectedInt(uint32(1), uint32(len(a)), t)
	checkExpectedInt(uint32(bi.Uint64()), a[0], t)

	a = string2Ints("0000100000000001")
	bi.SetString("0000100000000001", 2)
	checkExpectedInt(uint32(1), uint32(len(a)), t)

	checkExpectedInt(uint32(bi.Uint64()), a[0], t)

	bi.SetUint64(uint64(maxint) + 4)
	//fmt.Println(bi.Text(2))
	a = string2Ints(bi.Text(2))
	checkExpectedInt(uint32(2), uint32(len(a)), t)
	checkExpectedInt(uint32(1), a[0], t)
	checkExpectedInt(uint32(3), a[1], t)

}
