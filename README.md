# go-with-tests
Exercises from https://quii.gitbook.io/learn-go-with-tests

```bash
cd $MODULE_DIR

# run tests and testable examples
go test example.com/hello/maps

# get test coverage
go test example.com/hello/di -cover

# run benchmark tests
go test example.com/hello/concurrency -bench=.

# run memory benchmark tests
go test example.com/hello/concurrency -bench=. -benchmem

# # example memory benchmark output
# goos: linux
# goarch: amd64
# pkg: example.com/hello/iteration
# cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
# BenchmarkRepeat-16      57534962                20.53 ns/op            8 B/op          1 allocs/op
# PASS
# ok      example.com/hello/iteration     1.205s

# B/op: the number of bytes allocated per iteration
# allocs/op: the number of memory allocations per iteration
```