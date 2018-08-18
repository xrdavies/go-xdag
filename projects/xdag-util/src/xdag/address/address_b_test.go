package address

import (
	"testing"
)

func Benchmark_XdagAddress2hash(b *testing.B) {
	b.SetParallelism(200)
	b.RunParallel(func(pb *testing.PB) {
		var testAddress = []uint8("0ZMzmll399pSV45ELTp6TdY8SpMGt6u3")
		for pb.Next() {
			XdagAddress2hash(testAddress)
		}
	})
}

func Benchmark_XdagHash2address(b *testing.B) {
	b.SetParallelism(200)
	b.RunParallel(func(pb *testing.PB) {
		var testHash = []uint64{15778211046238688209, 5582838654177269586, 13234873168827137238, 324235667}
		for pb.Next() {
			XdagHash2Address(testHash)
		}
	})
}
