# benchmark_lock_performance

benchmark golang mutex & spinlock & spinlock sched perforemance 

### benchmark

`cpu 2.5, 8core`

**cmd**

```
./benchmark_lock_performance  -max=false
```

**performance stdout:**

```
spinlock
avg:  40 ns

spinlock sched
avg:  40 ns

mutex
avg:  184 ns
```