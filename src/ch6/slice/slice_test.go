package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s []int
	t.Log(len(s), cap(s))
	s = append(s, 1)
	t.Log(s, len(s), cap(s))

	// 注意slice切片与array的区别，array需要初始化数组的容量
	s1 := []int{1, 2, 3, 4, 5}
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 1)
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 1)
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 1)
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 1)
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 1)
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 1)
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 1, 1, 1)
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 1, 1, 1)
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 1, 1, 1)
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 1, 1, 1)
	t.Log(s1, len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	t.Log(s2[3])
}

func TestShareSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t.Log(len(s), cap(s))
	s1 := s[2:5]
	t.Log(s1, len(s1), cap(s1))
	s2 := s[3:9]
	t.Log(s2, len(s2), cap(s2))
	s2[0] = 1111
	t.Log(s1, len(s1), cap(s1))
	t.Log(s2, len(s2), cap(s2))
}

func TestShareSplite(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t.Log(getSum(s...))
}

func getSum(args ...int) int {
	var sum = 0
	for _, arg := range args {
		sum += arg
	}

	return sum
}
