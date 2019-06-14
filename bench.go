package main

import (
	"fmt"
	"sync"
	"time"

	"flag"
)

var (
	maxFlag bool
	mode    string
)

func init() {
	flag.BoolVar(&maxFlag, "max", false, "this help")
	flag.StringVar(&mode, "mode", "mutex", "spin & spin_sched & mutex")
}

type handler interface {
	add()
}

func execute(locker handler) {
	wg := sync.WaitGroup{}

	start := time.Now()
	g := 200
	n := 200000

	for index := 0; index < g; index++ {
		wg.Add(1)
		go func() {
			var max int64
			for idx := 0; idx < n; idx++ {
				if !maxFlag {
					locker.add()
					continue
				}

				s := time.Now().UnixNano()
				locker.add()
				e := time.Now().UnixNano()
				c := e - s
				if c > max {
					max = c
				}
			}

			if maxFlag {
				fmt.Println("max in g: ", max/1000/1000, "ms")
			}
			wg.Done()
		}()
	}

	wg.Wait()
	end := time.Now()
	cost := end.Sub(start).Nanoseconds()
	fmt.Println("avg: ", cost/int64(g*n), "ns")
}

func main() {
	flag.Parse()
	var locker handler

	fmt.Printf("spinlock \n")
	locker = &spinLock{
		sched: false,
	}
	execute(locker)

	fmt.Printf("\nspinlock sched \n")
	locker = &spinLock{
		sched: true,
	}
	execute(locker)

	fmt.Printf("\nmutex \n")
	locker = &locked{}
	execute(locker)
}
