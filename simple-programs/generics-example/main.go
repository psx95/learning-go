package main

import "fmt"

func main() {
	testscores := []int64{
		10,
		20,
		30,
		40,
	}

	heights := []float64{
		1.3,
		1.7,
		1.8,
	}

	clonedArray := clone(testscores)
	fmt.Println(&testscores[0], &clonedArray[0], clonedArray)

	genericCloneArray := genericClone[float64](heights)
	fmt.Println(&heights[0], &genericCloneArray[0], genericCloneArray)
}

// Clone creates a clone of an int64 slice
func clone(arr []int64) []int64 {
	result := make([]int64, len(arr))
	for i, v := range arr {
		result[i] = v
	}
	return result
}

// This is a generic function - defined by a pair of sqaure brackets []
// in between the function name and the parameters.
// The square brackets contain the generic type - in this case V along with
// the type constraint - which in this case is the builtin type any - indicating
// that we can take any type in place of V
func genericClone[V any](arr []V) []V {
	result := make([]V, len(arr))
	for i, v := range arr {
		result[i] = v
	}
	return result
}
