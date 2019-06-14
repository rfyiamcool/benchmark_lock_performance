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
