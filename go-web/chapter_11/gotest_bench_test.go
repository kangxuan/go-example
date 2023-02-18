package gotest

import "testing"

func BenchmarkDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

func BenchmarkDivision2(b *testing.B) {
	b.StopTimer() // 停止压力测试的时间计数

	// 做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	// 这样这些时间不影响我们测试函数本身的性能
	// ...

	b.StartTimer() // 重新开始压力测试的时间计数
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
