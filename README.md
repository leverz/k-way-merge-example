# k-way-merge-example

TODO:
- [x] simple merge sort
- [x] test case
- [x] 2-way merge sort
- [ ] k-way merge sort
- [ ] using heap sorting optimization k-way merge
- [ ] using multiple goroutine

## 2-way merge sort + multiple goroutine
```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/leverz/k-way-merge-example
BenchmarkSort-4              	2000000000	         0.15 ns/op
BenchmarkTwoSort-4           	2000000000	         0.10 ns/op
BenchmarkTwoSortLimitCPU-4   	2000000000	         0.16 ns/op
testing: BenchmarkTwoSortLimitCPU-4 left GOMAXPROCS set to 1
PASS
ok  	github.com/leverz/k-way-merge-example	11.911s
```
