package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)

	studentGrades["Prince"] = 100
	studentGrades["alice"] = 150
	studentGrades["bob"] = 190

	// fmt.Println("bob marks: ",studentGrades["bob"])
	delete(studentGrades,"bob")
	fmt.Println("bob marks: ",studentGrades["bob"])

	// to check if key exist 
	grades, exists := studentGrades["Alice"]
	fmt.Println("grade: ", grades)
	fmt.Println("Alice exists: ", exists)

// when we are paasing the dclared variables Grades,Exists we can name it whatever we like 
// first variable will show the value of the selected key and next one will check if it exists
	Grades, Exists := studentGrades["Prince"]
	fmt.Println("Prince marks: ", Grades)
	fmt.Println("Prince exists: ", Exists)

	for index,value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n",index,value)
	}
	// map returns the data in uordered way 

	person := map[string]int{
		"1st" : 100,
		"2nd" : 200,
		"3rd" : 300,
	}
	for index,value := range person {
		fmt.Printf("key is:%s and value is:%d\n",index,value)
	}



}
