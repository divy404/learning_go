package main

import "fmt"

func main() {
	numbers:= []int{1,2,3,4,5}
	numbers = append(numbers, 2,10,23,25,45,34)
	fmt.Println("Number: ",numbers)
	fmt.Printf("number has data type of: %T\n ",numbers)
	fmt.Println("Length : ", len(numbers))
	// we have one more thing know as capacity which tells us the size of the
	// array which lies under the slice
	fmt.Println("capacity:",cap(numbers))

	name:=[]string{}
	fmt.Println("name",name)

	// now lets try using 'make' function 
	no := make([]int,3,5) 
	// basically here the len of slice is 3 and capacity is 5
	// thing to remember is that capacity increases automatically and gets doubled 
	no = append(no, 4)
	no = append(no, 5)
	no = append(no, 6)
	no[1]=6
	fmt.Println("Slice: , Length: , Capacity: ", no,len(no),cap(no))

}