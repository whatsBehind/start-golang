package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stringToFloat()
}

func stringToFloat() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input a number: ")

	input, err := reader.ReadString('\n')
	fmt.Printf("The var type is %T \n", input)

	if err != nil {
		fmt.Println(err)
	} else {
		// input has "\n" in the end, so trim space before parsing
		num, _ := strconv.ParseFloat(strings.TrimSpace(input), 32)
		fmt.Printf("The var type is %T ", num)
	}
}
