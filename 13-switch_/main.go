package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	switchAmongOneToSix()
}

func switchAmongOneToSix() {
	rand.Seed(time.Now().UnixNano())
	randomValue := rand.Intn(6) + 1

	switch randomValue {
	case 1:
		fmt.Printf("random value %v\n", randomValue)
	case 2:
		fmt.Printf("random value %v\n", randomValue)
	case 3:
	case 4:
		fmt.Printf("random value %v\n", randomValue)
	case 5:
		fmt.Printf("random value %v\n", randomValue)
	default:
		fmt.Printf("random value %v\n", randomValue)
	}
}
