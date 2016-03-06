package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

type fileReader struct {
	n, q int
	a, b []byte
	f    *os.File
	s    *bufio.Scanner
}

func (fr fileReader) getInitial() (q int, a, b []byte) {
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
		reader.a = string2Bytes(reader.s.Text())
		reader.s.Scan()
		reader.b = string2Bytes(reader.s.Text())

	} else {
		return nil, e
	}
	return reader, e
}

func TestEndToEnd(t *testing.T) {
	path := "/Users/jam/code/go-challenges/chabit/test1.txt"
	if p, err := newFileReader(path); err == nil {
		defer p.(fileReader).close()
		run(p)
	} else {
		t.Error(err)
	}

}

func TestSetBitAtIdx(t *testing.T) {
	b := []byte("000000000000000")
	setBit(1, 1, &b)
	exp := getBit(1, b)
	if exp != 1 {
		t.Error("Incorrect result")
	}
	setBit(1, 0, &b)
	exp = getBit(1, b)
	if exp != 0 {
		t.Error("Error")
	}

}

func TestAdd(t *testing.T) {
	a := string2Bytes("0001")
	b := string2Bytes("0001")
	c := add(a, b)
	x := getBit(1, c)
	if x != 1 {
		t.Error("Error ", x)
	}

	a = string2Bytes("0011")
	b = string2Bytes("0011")
	s := bytes2String(add(a, b))
	if s != "00110" {
		t.Error("Error ", s)
	}
	a = string2Bytes("00011")
	b = string2Bytes("10010")
	s = bytes2String(add(a, b))
	if s != "010101" {
		t.Error("Error ", s)
	}
	a = string2Bytes("011111")
	b = string2Bytes("000001")
	s = bytes2String(add(a, b))
	if s != "0100000" {
		t.Error("Error ", s)
	}

	a = string2Bytes("11111")
	b = string2Bytes("11111")
	s = bytes2String(add(a, b))
	if s != "111110" {
		t.Error("Error ", s)
	}

}

func TestString2Bytes(t *testing.T) {
	a := string2Bytes("00010")
	fmt.Print(a)
	exp := getBit(1, a)
	if exp != 1 {
		t.Error("Error")
	}
	exp = getBit(2, a)
	if exp != 0 {
		t.Error("Error")
	}

}
