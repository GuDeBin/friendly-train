package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {

	ints := map[string]int64{
		"fist":   1,
		"second": 2,
	}

	floats := map[string]float64{
		"first":  1.1,
		"second": 2.2,
	}

	fmt.Printf("Non-generic sums: %v and %v\n", SumInts(ints), SumFloats(floats))
	fmt.Printf("generic sums: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
	fmt.Printf("generic sums: %v and %v\n", SumNumbers(ints), SumNumbers(floats))
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
