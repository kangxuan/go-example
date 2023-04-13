package gotest

import (
	"strconv"
	"testing"
)

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

func TestDivisionGroup(t *testing.T) {
	type test struct {
		a    float64
		b    float64
		want float64
	}
	tests := []test{
		{1, 2, 0.5},
		{2, 1, 2},
		{3, 1, 3},
		{3, 0, 0},
	}
	for k, tc := range tests {
		if i, err := Division(tc.a, tc.b); i != tc.want || err != nil {
			t.Errorf("num:%d, expected:%v, got:%v, err:%s", k+1, tc.want, i, err.Error())
		} else {
			t.Logf("num:%d, expected:%v, got:%v", k+1, tc.want, i)
		}
	}
}

func TestDivisionGroupRun(t *testing.T) {
	type test struct {
		a    float64
		b    float64
		want float64
	}
	tests := []test{
		{1, 2, 0.5},
		{2, 1, 2},
		{3, 1, 3},
		{6, 2, 3},
		{3, 0, 0},
	}
	var num = 0
	for k, tc := range tests {
		num = k + 1
		t.Run("第"+strconv.Itoa(num)+"个", func(t *testing.T) {
			if i, err := Division(tc.a, tc.b); i != tc.want || err != nil {
				t.Errorf("expected:%v, got:%v, err:%s", tc.want, i, err.Error())
			}
		})
	}
}
