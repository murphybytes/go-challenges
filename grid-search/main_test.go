package main

import "testing"

func TestRow(t *testing.T) {
	m := findMatches("foobar", "foo")
	if len(m) != 1 {
		t.Errorf("Expected 1 got %d", len(m))
	}
	if m[0] != 0 {
		t.Errorf("Index should be 0")
	}

	m = findMatches("foobarfoo", "foo")
	if len(m) != 2 {
		t.Errorf("Expected 2 got %d", len(m))
	}
	if m[0] != 0 {
		t.Errorf("Index should be 0")
	}

	if m[1] != 6 {
		t.Errorf("Expected 6 got %d", m[1])
	}

	m = findMatches("foobarfoofo", "foo")
	if len(m) != 2 {
		t.Error("Incorrect length")
	}

	m = findMatches("561212", "12")
	if len(m) != 2 {
		t.Error("incorrect index %d\n", len(m))
	}

}

func TestIndexAt(t *testing.T) {
	match := matchAt(0, "foobar", "foo")
	if !match {
		t.Error("Expected true")
	}

	match = matchAt(0, "foobar", "bar")
	if match {
		t.Error("Expected false")
	}

	match = matchAt(3, "xxxfoozzz", "foo")
	if !match {
		t.Error("should have matched")
	}
}

func TestTest(t *testing.T) {
	g := []string{
		"123",
		"456",
		"789",
	}
	p := []string{
		"23",
		"56",
	}

	r := test(g, p)
	if !r {
		t.Error("expected true")
	}

	g = []string{
		"555",
		"456",
		"789",
	}

	r = test(g, p)
	if r {
		t.Error("expected false")
	}

	g = []string{
		"55555",
		"23423",
		"78456",
	}

	r = test(g, p)
	if !r {
		t.Error("expected true")
	}

	g = []string{
		"123412",
		"561212",
		"123634",
		"781288",
	}
	p = []string{
		"12",
		"34",
	}
	r = test(g, p)
	if !r {
		t.Error("expected true")
	}

}
