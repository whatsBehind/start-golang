package main

import (
	"fmt"
	"sort"
)

var fruits = []string{"apple", "peach", "banana"}

func main() {
	createSlice()
	appendValue()
	subSlice()
	sortSlice()
	deleteAtIndex()
}

func createSlice() {
	fmt.Printf("The type of fruits: %T\n", fruits)
}

func appendValue() {
	var appendedFruits = append(fruits, "pineapple")
	fmt.Println("Appended fruits slice: ", appendedFruits)
}

func subSlice() {
	// no copy is created, modify the subSlice will change value in original slice
	var subFruits = fruits[1:3]
	fmt.Println("Sub slice of fruits: ", subFruits)
	subFruits[0] = "peaches"
	fmt.Println("Original fruits", fruits)
}

func sortSlice() {
	sort.Strings(fruits)
	fmt.Println("Sorted fruits: ", fruits)
}

func deleteAtIndex() {
	var deletedAtIndexOneFruits = append(fruits[:1], fruits[2:]...)
	fmt.Println("Fruits deleted at index one: ", deletedAtIndexOneFruits)
}
