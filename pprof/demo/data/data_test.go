package data

import (
	"testing"
)
// 单元测试使用pprof
// go test -bench=. -cpuprofile=cpu.prof

const url = "https://github.com/flowerwang2012"

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Error("Test.Add error!")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}