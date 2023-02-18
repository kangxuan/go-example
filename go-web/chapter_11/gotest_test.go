package gotest

import "testing"

func TestDivision1(t *testing.T) {
	if i, e := Division(6, 3); i != 2 || e != nil {
		t.Error("第一个除法函数测试没通过")
	} else {
		t.Log("第一个测试通过")
	}
}

func TestDivision2(t *testing.T) {
	if i, e := Division(6, 0); i != 2 || e != nil {
		t.Error("第二个除法函数没通过")
	} else {
		t.Log("第二个测试通过")
	}
}
