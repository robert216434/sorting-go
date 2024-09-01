# Sorting algorithms in Go

This repository aims to contain most used sorting algorithms written in Golang.
Benchmark included.

## Running benchmark

To run the benchmark and compare the algorithms performance, open terminal and run:

```bash
go test -bench . -benchmem
```

Example of result:

```bash
generated slice with size 10000, max value 1000000000 in 3735708 nanoseconds
goos: linux
goarch: amd64
pkg: github.com/robert216434/sorting-go
cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz
Benchmark_MergeSort-16          1000000000               0.0008672 ns/op               0 B/op          0 allocs/op
Benchmark_NaiveSort-16          1000000000               0.05752 ns/op         0 B/op          0 allocs/op
PASS
ok      github.com/robert216434/sorting-go      0.491s
```
