package map_test

import "testing"

func TestMapFun(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(pa int) int { return pa }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapSet(t *testing.T) {
	m := map[int]bool{}
	m[10] = true
	n := 10
	if m[n] {
		t.Log("exist")
	} else {
		t.Log("not exist")
	}
	delete(m, 10)
	if m[n] {
		t.Log("exist")
	} else {
		t.Log("not exist")
	}
}
