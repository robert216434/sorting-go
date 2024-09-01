# Sorting algorithms in Go

This repository aims to contain most used sorting algorithms written in Golang.
Benchmark included.

## Algorithms

List of algorithms currently in repository:

- Bubble sort
- Heap sort
- Insertion sort
- Merge sort
- Naive sort
- Quick sort
- Radix sort
- Selection sort
- Slice (go standard library package [sort](https://pkg.go.dev/sort)), uses pattern-defeating quicksort (pdqsort)
- SliceStable (go standard library package [sort](https://pkg.go.dev/sort)), uses insertion sort
- SortFunc (go standard library package [slices](https://pkg.go.dev/slices)), uses pattern-defeating quicksort (pdqsort)
- SortStableFunc (go standard library package [slices](https://pkg.go.dev/slices)), uses insertion sort

## Benchmark

At init the benchmark generates a slice of size 10_000 with random integers between [0, 1_000_000_000] and passes a copy of the slice to each sorting algorithm benchmark.

## Running benchmark

To run the benchmark and compare the algorithms performance on your machine, open terminal at the root of the repository and run:

```bash
go test -bench . -benchmem -benchtime=1s
```

Example of result:

```bash
generated slice with size 10000, max value 1000000000 in 3680456 nanoseconds / 3 milliseconds
goos: linux
goarch: amd64
pkg: github.com/robert216434/sorting-go
cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz
Benchmark_BubbleSort-16                                       24          47843150 ns/op           81920 B/op          1 allocs/op
Benchmark_HeapSort-16                                       1575            748924 ns/op           81920 B/op          1 allocs/op
Benchmark_InsertionSort-16                                   100          10018571 ns/op           81920 B/op          1 allocs/op
Benchmark_MergeSort-16                                      1278            890244 ns/op         1194630 B/op      10000 allocs/op
Benchmark_NaiveSort-16                                        15          70182863 ns/op           81920 B/op          1 allocs/op
Benchmark_QuickSort-16                                      2749            427977 ns/op           81920 B/op          1 allocs/op
Benchmark_RadixSort-16                                      1996            558941 ns/op          819202 B/op         10 allocs/op
Benchmark_SelectionSort-16                                    32          34003523 ns/op           81920 B/op          1 allocs/op
Benchmark_SliceSortPackage-16                               1440            797447 ns/op           81976 B/op          3 allocs/op
Benchmark_SliceStableSortPackage-16                          637           1859706 ns/op           81976 B/op          3 allocs/op
Benchmark_SortFuncSlicesSlicesPackage-16                    1777            662052 ns/op           81920 B/op          1 allocs/op
Benchmark_SortStableFuncSlicesSlicesPackage-16               844           1412044 ns/op           81920 B/op          1 allocs/op
PASS
ok      github.com/robert216434/sorting-go      14.559s
```

## How to interpret results

`iteration` is a for loop pass inside of `Benchmark` function. By default it tries to do `1_000_000_000` iterations in the given time.

- First column is the name of the algorithm
- Second column is how many iterations were made in the given time in `-benchtime`, default is 1 second. (**higher** is better)
- `ns/op` means nanoseconds per operation, how much time it took per iteration. (**less** is better)
- `B/op` means how many bytes were allocated per iteration. (**less** is better)
- `allocs/op` means how many memory allocations occurred per iteration. (**less** is better)
