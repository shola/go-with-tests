# go-with-tests
Exercises from https://quii.gitbook.io/learn-go-with-tests

```bash
cd $MODULE_DIR

# run tests and testable examples
go test shola.com/gwt/maps

# get test coverage
go test shola.com/gwt/di -cover

# run benchmark tests
go test shola.com/gwt/concurrency -bench=.

# test for race conditions
go test shola.com/gwt/concurrency -race

# run memory benchmark tests
go test shola.com/gwt/concurrency -bench=. -benchmem

# find subtle bugs in code
go vet shola.com/gwt/sync

# # example memory benchmark output
# goos: linux
# goarch: amd64
# pkg: shola.com/gwt/iteration
# cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
# BenchmarkRepeat-16      57534962                20.53 ns/op            8 B/op          1 allocs/op
# PASS
# ok      shola.com/gwt/iteration     1.205s

# B/op: the number of bytes allocated per iteration
# allocs/op: the number of memory allocations per iteration
```