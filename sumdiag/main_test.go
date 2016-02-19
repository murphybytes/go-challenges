package main

import "testing"

func Test_sumdiags(t *testing.T) {

	m := [][]int{
		{1, 2, 3},
		{3, 4, 5},
		{10, 7, 8},
	}

	lr, rl := sumdiags(m)
	if lr != 13 {
		t.Errorf("Expected %d, Got %d", 13, lr)
	}

	if rl != 17 {
		t.Errorf("Expected %d, Got %d", 17, rl)
	}

}

func Test_sumdiags2(t *testing.T) {

	m := [][]int{
		{1, 2, 3, -4},
		{3, 4, 5, 0},
		{10, 7, 8, 0},
		{-9, 0, 0, 5},
	}

	lr, rl := sumdiags(m)
	if lr != 18 {
		t.Errorf("Expected %d, Got %d", 13, lr)
	}

	if rl != -1 {
		t.Errorf("Expected %d, Got %d", 17, rl)
	}

}
