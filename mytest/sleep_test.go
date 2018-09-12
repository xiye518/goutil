package mytest

import (
	"time"
)

/*
go test -bench . -benchtime 5s		使用benchtime设定最小测试时间来增加循环次数，以便返回更准确的结果
*/

func sleep() {
	time.Sleep(time.Second)
}

//func BenchmarkSleep(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		sleep()
//	}
//}
