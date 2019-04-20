package main

import "fmt"

func arrayOps(repoName string, projectName string) {
	var slice = []int{3, 5, 4} // slice of variable size
	fmt.Println(slice[0])

	fmt.Println(len(slice))

	for i := 0; i < len(slice); i++ {
		fmt.Println("a----", slice[i])
	}

	for in, val := range slice {
		fmt.Println("index ", in)
		fmt.Println("value", val)
	}

	newSlice := mutableArray(slice)
	fmt.Println("Mutated Array: ", newSlice)

	var array = [3]int{0, 2, 4} // array of size 3
	immutableArray(array)
	fmt.Println("Immutable Array:", array)

}

func mutableArray(a []int) []int { // slices passes a reference to the underlying array, hence values will be modified in the original
	b := append(a, 5, 6)
	return b
}

func immutableArray(a [3]int) { // array passes a copy of original array, hence original array is not modified
	a[0] = 89
	fmt.Println("Changed array value:", a)
}
