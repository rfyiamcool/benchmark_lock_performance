# benchmark_lock_performance

benchmark golang mutex & spinlock & spinlock sched perforemance 

## benchmark

**test condition**

* Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz (12 core)
* go1.16.7 darwin/amd64 

### custom test

cmd

```
go build; ./benchmark_lock_performance  -max=false
```

performance stdout:

```
spinlock
avg:  600 ns

spinlock with sched
avg:  32 ns

mutex
avg:  100 ns
```

**see max latency**

```
go build; ./benchmark_lock_performance  -max=true
```

### test std bench

cmd

```
go test -v -bench . -count=1 -benchmem
```

performance stdout:

```
goos: darwin
goarch: amd64
pkg: benchmark_lock_performance
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz

BenchmarkNotLockedAdd-12            	847063903	         1.305 ns/op	       0 B/op	       0 allocs/op
BenchmarkParalletNotLockedAdd-12    	454034742	         2.870 ns/op	       0 B/op	       0 allocs/op

BenchmarkLockedAdd-12               	81570625	        13.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkParalletLockedAdd-12       	13276042	       100.5 ns/op	       0 B/op	       0 allocs/op

BenchmarkSpinAdd-12                 	91128793	        12.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkParalletSpinAdd-12         	 1000000	      2158 ns/op	       0 B/op	       0 allocs/op

BenchmarkSpinSchedAdd-12            	83897924	        12.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkParalletSpinSchedAdd-12    	39286161	        32.51 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	benchmark_lock_performance	11.418s
```