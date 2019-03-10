package constant_test

import "testing"

const (
	a = 11111
	b
	c
	d
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota // 1
	Writable             // 11
	Eatable              // 111
)

func TestConstant(t *testing.T) {
	t.Log(a, b, c, d)
	t.Log(Monday, Tuesday, Wednesday)
	t.Log(Readable, Writable, Eatable)
}
