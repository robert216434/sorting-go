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
generated slice with size 10000, max value 1000000000 in 3736125 nanoseconds / 3 milliseconds
goos: linux
goarch: amd64
pkg: github.com/robert216434/sorting-go
cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz
Benchmark_MergeSort-16              1294            871662 ns/op         1194629 B/op      10000 allocs/op
Benchmark_NaiveSort-16                15          69740839 ns/op           81920 B/op          1 allocs/op
Benchmark_QuickSort-16              2775            409129 ns/op           81920 B/op          1 allocs/op
PASS
ok      github.com/robert216434/sorting-go      3.531s
```
