// package to build must be main
package main

import(
	"fmt"
	_ "os"
	)
	// _ allows an unused import

// { must follow function
func main(){
	fmt.Println("This is a smaple Go program")
}

// go build filename.go (build executeable)
// go run filename.go (scripting)