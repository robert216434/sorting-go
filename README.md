# Sorting algorithms in Go

This repository aims to contain most used sorting algorithms written in Golang.
Benchmark included.

## Algorithms

List of algorithms currently in repository:

- Merge sort
- Quick sort
- Naive sort
- SliceStable (go standard library), uses insertion sort

## Running benchmark

To run the benchmark and compare the algorithms performance, open terminal and run:

```bash
go test -bench . -benchmem -benchtime=1s
```

Example of result:

```bash
generated slice with size 10000, max value 1000000000 in 3667273 nanoseconds / 3 milliseconds
goos: linux
goarch: amd64
pkg: github.com/robert216434/sorting-go
cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz
Benchmark_MergeSort-16                      1345            875341 ns/op         1194629 B/op      10000 allocs/op
Benchmark_NaiveSort-16                        16          69950942 ns/op           81920 B/op          1 allocs/op
Benchmark_QuickSort-16                      2821            417973 ns/op           81920 B/op          1 allocs/op
Benchmark_SliceStablePackage-16              654           1841209 ns/op           81976 B/op          3 allocs/op
PASS
ok      github.com/robert216434/sorting-go      5.083s
```

## How to interpret results

- First column is the name of the algorithm
- Second column is how many iterations were made in the given time in `-benchtime`, default is 1 second. (higher is better)
- `ns/op` means nanoseconds per operation, how much time it took per iteration. (less is better)
- `B/op` means how many bytes were allocated per iteration. (less is better)
- `allocs/op` means how many memory allocations occurred per iteration. (less is better)
