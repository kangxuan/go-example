package split

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

//func TestMoreSplit(t *testing.T) {
//	got := Split("abcd", "bc")
//	want := []string{"a", "d"}
//	if !reflect.DeepEqual(want, got) {
//		t.Errorf("expected:%v, got:%v", want, got)
//	}
//}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:B:c", ":")
	}
}

func BenchmarkSplitParallel(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("沙河有沙又有河", "沙")
		}
	})
}

func ExampleSplit() {
	fmt.Println(Split("a:b:c", ":"))
	fmt.Println(Split("沙河有沙又有河", "沙"))
	// Output:
	// [a b c]
	// [ 河有 又有河]
}
