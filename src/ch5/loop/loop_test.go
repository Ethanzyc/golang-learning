package loop_test

import (
	"fmt"
	"testing"
)

func TestWhileLoop(t *testing.T) {
	i := 0
	for i < 5 {
		t.Log(i)
		i++
	}
}

func TestForLoop(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(i)
	}
}

func TestIf(t *testing.T) {
	i := 5
	if i < 10 {
		t.Log("小于10")
	} else {
		t.Log("大于10")
	}
}

func TestIfCondition(t *testing.T) {
	// i := 5
	// if v, error := beError(); error == nil {
	// 	t.Log("小于10")
	// } else {
	// 	t.Log("大于10")
	// }
}

func beError() {
	i := 1 / 1
	fmt.Print(i)
}
