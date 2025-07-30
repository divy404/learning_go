package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}
type Contact struct {
	Email string
	Phone string
}
type Address struct {
	House int 
	Area string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	// we use struct to store diffrent types of data together
	// method 1
	var divi Person
	fmt.Println("Divi:",divi)
	divi.FirstName="divyansh"
	divi.LastName="jain"
	divi.Age=19
	fmt.Println("Divi:",divi)

	// method 2

	person1 := Person{
		FirstName: "pookiee",
		LastName: "hehe",
		Age: 10,
	}
	fmt.Println("first person:",person1)

	// method 3 using new keyword
	// new is a type of pointer 

	var person2 = new(Person)
	person2.FirstName="meow"
	person2.LastName="meowchann"
	person2.Age=1
	fmt.Println("person2:",person2)

	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "divy",
		LastName: "jain",
		Age: 19,
	}
	employee1.Person_Contact = Contact{
		Email: "divyanshj@gmail.com",
		Phone: "76948118",
	}
	employee1.Person_Address = Address{
		House: 5,
		Area: "kshir sagar",
		State: "M.P",
	}
	fmt.Println("Employee 1 details are:", employee1)
	// for specific values 
	fmt.Println("Employee 1 address: ", employee1.Person_Address.House)

}