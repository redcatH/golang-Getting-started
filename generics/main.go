package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumInits(m map[string]int64) int64 {
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

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := map[string]int64{"a": 1, "b": 2, "c": 3}
	floats := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}

	fmt.Printf("None-generic  sum:%v and sum:%v \n", SumInits(ints), SumFloats(floats))
	fmt.Printf("generic  sum:%v and sum:%v \n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
}
