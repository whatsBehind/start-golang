package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	printRandomValue()
}

func printRandomValue() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number 0, 1, 2
	randomValue := rand.Intn(3)

	if randomValue == 0 {
		fmt.Println("Random value 0")
	} else if randomValue == 1 {
		fmt.Println("Random value 1")
	} else {
		fmt.Println("Random value 2")
	}
}
