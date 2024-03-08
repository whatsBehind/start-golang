package main

import "fmt"

func main() {
	getAnItem()
	updateAnItem()
	iterateMap()
}

func makeAMap() map[string]string {
	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["py"] = "Python"
	languages["cpp"] = "C++"
	return languages
}

func getAnItem() {
	var languages = makeAMap()
	fmt.Println(languages["js"])
}

func updateAnItem() {
	var languages = makeAMap()
	languages["js"] = "javascript"
	fmt.Println(languages["js"])
}

func iterateMap() {
	for key, value := range makeAMap() {
		fmt.Printf("key is %v, and value is %v\n", key, value)
	}
}
