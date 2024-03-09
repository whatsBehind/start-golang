package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	writeFile()
	readFile("my-file.txt")
}

func writeFile() {
	file, err := os.Create("./my-file.txt")
	checkError(err)
	file.WriteString("Hello World")

	defer file.Close()
}

func readFile(fileName string) {
	var bytes, err = ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("The content in the file:", string(bytes))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
