package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	
	var f *os.File
	f = os.Stdin
	// prevent closing when run
	defer f.Close()
	// creating an input variable
	scanner := bufio.NewScanner(f)
	// loop over input
	for scanner.Scan(){
		fmt.Println(">", scanner.Text())
	}
}