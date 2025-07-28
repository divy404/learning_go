package main

import "fmt"

func main() {
	for i := range 4 {
		fmt.Println("numbers are:",i)
	}

	counter := 0
	for {
		fmt.Println("Infinite loop")
		counter++;
		// to make infinite loop 
		if counter == 2{
			break
		}
	}
	numbers := []int{1, 2, 34, 5}
	for index, value := range numbers {
		fmt.Printf("index of numbers is %d, and value is %d\n", index, value)
	}

	data:= "Hello my pookieee"
	for index, char := range data {
		fmt.Printf("Index of Data is %d, and value is %c\n",index,char)
	}
}