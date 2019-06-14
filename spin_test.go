package main

import (
	"testing"
)

func BenchmarkSpinAdd(t *testing.B) {
	c := spinLock{
		sched: false,
	}
	for i := 0; i < t.N; i++ {
		c.add()
	}
	if c.i != t.N {
		t.Errorf("Expected %d, but got %d", t.N, c.i)
	}
}

func BenchmarkSpinSchedAdd(t *testing.B) {
	c := spinLock{
		sched: true,
	}
	for i := 0; i < t.N; i++ {
		c.add()
	}
	if c.i != t.N {
		t.Errorf("Expected %d, but got %d", t.N, c.i)
	}
}

func BenchmarkParalletSpinAdd(b *testing.B) {
	c := spinLock{
		sched: false,
	}
	b.SetParallelism(200)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.add()
		}
	})
}

func BenchmarkParalletSpinSchedAdd(b *testing.B) {
	c := spinLock{
		sched: true,
	}
	b.SetParallelism(200)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.add()
		}
	})
}
