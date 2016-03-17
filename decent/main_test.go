package main

import "testing"

func TestDecent(t *testing.T) {
	digits := 1
	_, _, p := getDecent(digits)
	if p == true {
		t.Error("Expected false")
	}

	_, _, p = getDecent(2)
	if p == true {
		t.Error("Expected false")
	}

	f, th, p := getDecent(3)
	if p == false {
		t.Error("Expected true")
	}
	if f != 3 && th != 0 {
		t.Error("Expected 3,0 got ", f, ", ", th)
	}

	f, th, p = getDecent(5)
	if p == false {
		t.Error("Expected true")
	}
	if f != 0 && th != 5 {
		t.Error("Expected 0,5 got ", f, ", ", th)
	}

	f, th, p = getDecent(11)
	if p == false {
		t.Error("Expected true")
	}
	if f != 6 && th != 5 {
		t.Error("Expected 6,5 got ", f, ", ", th)
	}
	f, th, p = getDecent(17)
	if p == false {
		t.Error("Expected true")
	}
	if f != 12 && th != 5 {
		t.Error("Expected 12,5 got ", f, ", ", th)
	}

	f, th, p = getDecent(23)
	if p == false {
		t.Error("Expected true")
	}
	if f != 18 && th != 5 {
		t.Error("Expected 18,5 got ", f, ", ", th)
	}

	f, th, p = getDecent(100)
	if p == false {
		t.Error("Expected true")
	}
	if f != 90 && th != 10 {
		t.Error("Expected 90,10 got ", f, ", ", th)
	}

}
