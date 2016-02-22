package main

import "testing"

func TestCreateTransform(t *testing.T) {
	m := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	res := getTransform(m, 4, 4, 0, 0)
	if len(res) != 12 {
		t.Errorf("Expected 16 got %d", len(res))
	}

	expected := []int{1, 5, 9, 13, 14, 15, 16, 12, 8, 4, 3, 2}
	for i := range expected {
		if expected[i] != res[i] {

			t.Errorf("Incorrect result at index %d. Expected %d got %d", i, expected[i], res[i])
		}

	}

	res = getTransform(m, 2, 2, 1, 0)
	expected = []int{6, 10, 11, 7}
	for i := range expected {
		if expected[i] != res[i] {
			t.Errorf("Incorrect result at index %d. Expected %d got %d", i, expected[i], res[i])
		}
	}

}

func TestSmallTransform(t *testing.T) {

	m := [][]int{
		{1, 2},
		{3, 4},
	}
	res := getTransform(m, 2, 2, 0, 0)
	expected := []int{1, 3, 4, 2}
	if len(res) != len(expected) {
		t.Errorf("Expected %d got %d", len(expected), len(res))
	}

	for i := range expected {
		if expected[i] != res[i] {
			t.Errorf("Incorrect result at index %d. Expected %d got %d", i, expected[i], res[i])
		}
	}
}

func TestLayerRotation(t *testing.T) {
	m := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	res := getTransform(m, 4, 4, 0, 1)
	if len(res) != 12 {
		t.Errorf("Expected 16 got %d", len(res))
	}

	expected := []int{2, 1, 5, 9, 13, 14, 15, 16, 12, 8, 4, 3}
	for i := range expected {
		if expected[i] != res[i] {
			//		fmt.Print(*res[i])
			t.Errorf("Incorrect result at index %d. Expected %d got %d", i, expected[i], res[i])
		}

	}

	mm := [][]int{
		{2, 3, 4, 8},
		{1, 6, 7, 12},
		{5, 10, 11, 16},
		{9, 13, 14, 15},
	}

	applyTransform(&m, 4, 4, 0, res)

	for r := range mm {
		for c := range mm[r] {
			if mm[r][c] != m[r][c] {
				t.Errorf("Expected %d got %d", mm[r][c], m[r][c])
			}
		}
	}

}

func TestRotation(t *testing.T) {
	m := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	rotate(&m, 4, 4, 1)

	mm := [][]int{
		{2, 3, 4, 8},
		{1, 7, 11, 12},
		{5, 6, 10, 16},
		{9, 13, 14, 15},
	}

	for r := range mm {
		for c := range mm[r] {
			if mm[r][c] != m[r][c] {
				t.Errorf("Expected %d got %d", mm[r][c], m[r][c])
			}
		}
	}

}
