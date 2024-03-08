package main

import "fmt"

func main() {
	loopIndex()
	loopIndexInRange()
	loopIndexAndValue()
	breakLoop()
	gotoDemo()
}

func makeASlice() []string {
	return []string{"Monday", "Tuesday", "Friday", "Saturday", "Sunday"}
}

func loopIndex() {
	var days = makeASlice()

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
}

func loopIndexInRange() {
	var days = makeASlice()

	for i := range days {
		fmt.Println(days[i])
	}
}

func loopIndexAndValue() {
	var days = makeASlice()
	for index, value := range days {
		fmt.Printf("Index is %v, and value is %v\n", index, value)
	}
}

func breakLoop() {
	var index = 1
	for index < 10 {
		fmt.Printf("Index is %v\n", index)
		if index == 5 {
			break
		}
		index++
	}
}

func gotoDemo() {
	var index = 1
	for index < 10 {
		fmt.Printf("Index is %v\n", index)
		if index == 2 {
			goto gotoDemo
		}
		index++
	}

gotoDemo:
	fmt.Println("I am from goto")
}
