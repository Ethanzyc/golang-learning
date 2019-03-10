package case_test

import "testing"

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 1, 3:
			t.Log("奇数")
		case 0, 2:
			t.Log("偶数")
		default:
			t.Log("not in 0-3")
		}
	}
}

func TestSwitchCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 1:
			t.Log("奇数", i)
		case i%2 == 0:
			t.Log("偶数", i)
		default:
			t.Log(i)
		}
	}
}
