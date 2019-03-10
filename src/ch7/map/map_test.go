package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	t.Log(m["a"], len(m))
	m2 := map[string]int{}
	m2["aaa"] = 5
	t.Log(len(m2))
	m3 := make(map[string]int, 10)
	t.Log(len(m3))
}

func TestMapEleExist(t *testing.T) {
	m := map[string]int{}
	t.Log(m["1"])
	m["1"] = 1
	if v, ok := m["1"]; ok {
		t.Log("m[1] exist,value", v, ok)
	} else {
		t.Log("m[1] not exist,value", v, ok)
	}
}

func TestTravelMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		t.Log(k, v)
	}
}
