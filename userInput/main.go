package main

import (
	"bufio"
	"fmt"
	"os"
)
func main () {
	fmt.Println("What's your name? ")
	// var name string
	// fmt.Scan(&name) 
	// // noteee ->> if we use scan it only asks for input before the white space 
	// fmt.Println("Hello Mr.",name) 

	reader := bufio.NewReader(os.Stdin)
	name,_ := reader.ReadString('\n')
	fmt.Println("Hello Mr.", name)

}