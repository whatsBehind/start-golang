package main

import "fmt"

// "go env" returns env parameters
// "go build" build the module with main() function
// You can build with env parameter declaration like 'GOOS="windows" go build"
// Above command set env parameter GOOS to windows and then build the package
func main() {
	fmt.Println("Hello World")
}
