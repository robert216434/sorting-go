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

At init the benchmark generates a slice of size `10_000` with random integers in range `[0, 1_000_000_000]` and passes a copy of the slice to each sorting algorithm benchmark.

## Running benchmark

To run the benchmark and compare the algorithms performance on your machine, open terminal at the root of the repository and run:

```bash
go test -bench . -benchmem -benchtime=1s
```

Example of result from GitHub Actions runner:

```bash
generated slice with size 10000, max value 1000000000 in 8384579 nanoseconds / 8 milliseconds
goos: linux
goarch: amd64
pkg: github.com/robert216434/sorting-go
cpu: AMD EPYC 7763 64-Core Processor
Benchmark_BubbleSort-4                          	      13	  84588590 ns/op	   81920 B/op	       1 allocs/op
Benchmark_HeapSort-4                            	    1302	    914729 ns/op	   81920 B/op	       1 allocs/op
Benchmark_InsertionSort-4                       	      74	  15858771 ns/op	   81920 B/op	       1 allocs/op
Benchmark_MergeSort-4                           	    1116	   1056431 ns/op	 1194633 B/op	   10000 allocs/op
Benchmark_NaiveSort-4                           	      21	  54114159 ns/op	   81920 B/op	       1 allocs/op
Benchmark_QuickSort-4                           	    2443	    481215 ns/op	   81920 B/op	       1 allocs/op
Benchmark_RadixSort-4                           	    1845	    662291 ns/op	  819202 B/op	      10 allocs/op
Benchmark_SelectionSort-4                       	      25	  47181361 ns/op	   81920 B/op	       1 allocs/op
Benchmark_SliceSortPackage-4                    	    1254	    943232 ns/op	   81976 B/op	       3 allocs/op
Benchmark_SliceStableSortPackage-4              	     518	   2278474 ns/op	   81976 B/op	       3 allocs/op
Benchmark_SortFuncSlicesSlicesPackage-4         	    1419	    835150 ns/op	   81920 B/op	       1 allocs/op
Benchmark_SortStableFuncSlicesSlicesPackage-4   	     680	   1750437 ns/op	   81920 B/op	       1 allocs/op
PASS
ok  	github.com/robert216434/sorting-go	15.232s
```

## How to interpret results

`iteration` is a for loop pass inside of `Benchmark` function. By default it tries to do `1_000_000_000` iterations in the given time.

- First column is the name of the algorithm
- Second column is how many iterations were made in the given time in `-benchtime`, default is 1 second. In other words, represents how many times the sorting algorithm managed to sort the slice in 1 second. (**higher** is better)

- `ns/op` means nanoseconds per operation, how much time it took per iteration. (**less** is better)
- `B/op` means how many bytes were allocated per iteration. (**less** is better)
- `allocs/op` means how many memory allocations occurred per iteration. (**less** is better)

## Remarks

Algorithms implementation might not be optimal. From the results in this specific usecase QuickSort is fastest.
