package mytest

/*
go test -bench .
go test -bench . -cpu 1,2,4  使用cpu参数设定多个并发限制来观察结果
*/

func add(x, y int) int {
	return x + y
}

//性能测试
//func BenchmarkAdd(b *testing.B) {
//	//println("B.N=", b.N)
//
//	for i := 0; i < b.N; i++ {
//		_ = add(1, 2)
//	}
//}

/*
标准库现有的benchmark演示：
go test -run NONE -bench . -memprofile mem.out -cpuprofile cpu.out net/http
*/
