package array_test

import "testing"

func TestArray(t *testing.T) {
	var arr [3]int
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 4, 6, 7}
	t.Log(arr)
	t.Log(arr1[1])
	t.Log(len(arr2))
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 4, 6, 7}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
	// for i, v := range arr {
	// 下划线用于占位
	for _, v := range arr {
		t.Log(v)
	}
}

// 数组截取
func TestArrayRange(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	// 数组截取的时候含头不含尾
	arr1 := arr[1:3]
	arr2 := arr[1:len(arr)]
	arr3 := arr[:2]
	arr4 := arr[1:]
	t.Log(arr)
	t.Log(arr1)
	t.Log(arr2)
	t.Log(arr3)
	t.Log(arr4)
}
