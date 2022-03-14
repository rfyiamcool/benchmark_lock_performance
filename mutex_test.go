package main

import (
	"testing"
)

func TestNotLockedAdd(t *testing.T) {
	c := notLocked{}
	c.add()
	if c.i != 1 {
		t.Errorf("Expected 1 but got %d", c.i)
	}
}

func BenchmarkNotLockedAdd(t *testing.B) {
	c := notLocked{}
	for i := 0; i < t.N; i++ {
		c.add()
	}
	if c.i != t.N {
		t.Errorf("Expected %d, but got %d", t.N, c.i)
	}
}

func BenchmarkParalletNotLockedAdd(b *testing.B) {
	c := notLocked{}
	b.SetParallelism(50)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.add()
		}
	})
}

func TestLockedAdd(t *testing.T) {
	c := locked{}
	c.add()
	if c.i != 1 {
		t.Errorf("Expected 1 but got %d", c.i)
	}
}

// defualt call 5 times, order call, not concurrent call
func BenchmarkLockedAdd(t *testing.B) {
	c := locked{}
	for i := 0; i < t.N; i++ {
		c.add()
	}
	if c.i != t.N {
		t.Errorf("Expected %d, but got %d", t.N, c.i)
	}
}

func BenchmarkParalletLockedAdd(b *testing.B) {
	c := locked{}
	b.SetParallelism(50)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.add()
		}
	})
}

// go test -v -bench . -count=1 -benchmem -cpu=8
