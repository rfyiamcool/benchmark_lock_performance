# benchmark_lock_performance

benchmark golang mutex & spinlock & spinlock sched perforemance 

## benchmark

**test condition**

* cpu 2.5hz, 8core
* go1.12.6 linux/amd64

### custom test

**cmd**

```
go build; ./benchmark_lock_performance  -max=false
```

**performance stdout:**

```
spinlock
avg:  429 ns

spinlock sched
avg:  42 ns

mutex
avg:  150 ns
```

**see max latency**

```
go build; ./benchmark_lock_performance  -max=true
```

### test std bench

**cmd**

```
go test -v -bench . -count=1 -benchmem
```

**performance stdout:**

```
goos: linux
goarch: amd64
BenchmarkLockedAdd-16               	100000000	        21.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkParalletAdd-16             	20000000	       111 ns/op	       0 B/op	       0 allocs/op
BenchmarkNotLockedAdd-16            	2000000000	         1.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkSpinAdd-16                 	100000000	        19.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkSpinSchedAdd-16            	100000000	        19.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkParalletSpinAdd-16         	 2000000	       648 ns/op	       0 B/op	       0 allocs/op
BenchmarkParalletSpinSchedAdd-16    	30000000	        41.8 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	_/root/benchmark_lock_performance	15.879s
```