package main

import (
	"fmt"
	"time"
)

// The number is fixed, you can't change the value
var DefaultTimeFormat string = "01-02-2006 15:04:05 Monday"

func main() {
	printNow()
	printAnyDate()
}

func printNow() {
	currentTime := time.Now().Format(DefaultTimeFormat)
	fmt.Println(currentTime)
}

func printAnyDate() {
	anyDate := time.Date(1994, time.July, 31, 7, 0, 0, 0, time.UTC).Format(DefaultTimeFormat)
	fmt.Println(anyDate)
}
