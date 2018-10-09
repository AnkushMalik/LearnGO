package main

import "fmt"

func main() {

	//*******Slices*********
	// slice1 := []int{3, 5, 2, 1, 4}
	// fmt.Println(slice1)
	// // slice2 := slice1[2:]
	// slice2 := make([]int, 5, 10)

	// copy(slice2, slice1)
	// // slice2 = append(slice2, 0, 1)

	// fmt.Println(slice2)

	//*******functions*********
	listNum := []int{5, 3, 5, 1, 4, 8}
	fmt.Println("Sum of all elements : ", Sumofeles(listNum))

}

// Sumofeles : unexported method
func Sumofeles(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
