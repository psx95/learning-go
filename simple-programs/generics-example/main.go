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

	testscoreMap := map[string]float64{
		"A": 11.4,
		"B": 21.2,
		"C": 31.1,
		"D": 41.0,
	}

	testscoreMapInt := map[string]int64{
		"A": 11,
		"B": 21,
		"C": 31,
		"D": 41,
	}

	clonedArray := clone(testscores)
	fmt.Println(&testscores[0], &clonedArray[0], clonedArray)

	genericCloneArray := genericClone[float64](heights)
	fmt.Println(&heights[0], &genericCloneArray[0], genericCloneArray)

	clonedMap := cloneMap(testscoreMap)
	fmt.Println(clonedMap)

	genericClonedMapInt := genericCloneMap[string, int64](testscoreMapInt)
	fmt.Println(genericClonedMapInt)

	genericClonedMapFloat := genericCloneMap[string, float64](testscoreMap)
	fmt.Println(genericClonedMapFloat)

}

// Clone creates a clone of an int64 slice
func clone(arr []int64) []int64 {
	result := make([]int64, len(arr))
	for i, v := range arr {
		result[i] = v
	}
	return result
}

// cloneMap creates a clone of a string to float64 map
func cloneMap(m map[string]float64) map[string]float64 {
	result := make(map[string]float64, len(m))
	for k, v := range m {
		result[k] = v
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

func genericCloneMap[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}
