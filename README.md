# Sorting algorithms in Go

This repository aims to contain most used sorting algorithms written in Golang.
Benchmark included.

## Algorithms

List of algorithms currently in repository:

- Merge sort
- Quick sort
- Naive sort
- Slice (go standard library package [sort](https://pkg.go.dev/sort)), uses pattern-defeating quicksort (pdqsort)
- SliceStable (go standard library package [sort](https://pkg.go.dev/sort)), uses insertion sort
- SortFunc (go standard library package [slices](https://pkg.go.dev/slices)), uses pattern-defeating quicksort(pdqsort)

## Running benchmark

To run the benchmark and compare the algorithms performance on your machine, open terminal at the root of the repository and run:

```bash
go test -bench . -benchmem -benchtime=1s
```

Example of result:

```bash
generated slice with size 10000, max value 1000000000 in 3761525 nanoseconds / 3 milliseconds
goos: linux
goarch: amd64
pkg: github.com/robert216434/sorting-go
cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz
Benchmark_MergeSort-16                      1296            889164 ns/op         1194630 B/op      10000 allocs/op
Benchmark_NaiveSort-16                        15          73348896 ns/op           81920 B/op          1 allocs/op
Benchmark_QuickSort-16                      2780            413782 ns/op           81920 B/op          1 allocs/op
Benchmark_SlicePackage-16                   1480            787783 ns/op           81976 B/op          3 allocs/op
Benchmark_SliceStablePackage-16              632           1858785 ns/op           81976 B/op          3 allocs/op
Benchmark_SortFuncSlices-16                 1798            649498 ns/op           81920 B/op          1 allocs/op
PASS
ok      github.com/robert216434/sorting-go      7.482s
```

## How to interpret results

`iteration` is a for loop pass inside of `Benchmark` function. By default it tries to do `1_000_000_000` iterations in the given time.

- First column is the name of the algorithm
- Second column is how many iterations were made in the given time in `-benchtime`, default is 1 second. (**higher** is better)
- `ns/op` means nanoseconds per operation, how much time it took per iteration. (**less** is better)
- `B/op` means how many bytes were allocated per iteration. (**less** is better)
- `allocs/op` means how many memory allocations occurred per iteration. (**less** is better)
