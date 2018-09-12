package mytest

/*
go test -bench . -benchmem -gcflags "-N -l"
*/

func heap() []byte {
	return make([]byte, 1024*10)
}

//func BenchmarkHeap(b *testing.B) {
//	//设置为总是输出内存分配信息，无论使用benchmem参数与否
//	//go test -bench .  -gcflags "-N -l"
//	b.ReportAllocs()
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		_ = heap()
//	}
//}
