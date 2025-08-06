package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error getting:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("error getting response:", res.Status)
		return
	}
	//  data,err1:=ioutil.ReadAll(res.Body)
	//  if err1 != nil {
	// 	fmt.Println("error",err1)
	// 	return
	//  }
	//  fmt.Println("Data: ", string(data))
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("error in decoding:", err)
		return
	}
	fmt.Println("Todo: ", todo)

}
func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Divii",
		Completed: true,
	}
	// Convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return
	}
	// convert json data to string
	jsonString := string(jsonData)

	//  convert string data to reader
	jsonReader := strings.NewReader(jsonString)
	myUrl := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("error sending request: ", err)
		return
	}
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response:", string(data))

}
func performUpdateRequest() {
	todo := Todo{
		UserID:    20,
		Title:     "Diviiiii",
		Completed: false,
	}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marhsalling", err)
		return
	}
	// json data to string
	jsonString := string(jsonData)
	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	// create PUT Request
	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT Request:", err)
		return
	}
	req.Header.Set("Content-type", "application/json")
	// Send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending request", err)
		return
	}
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response:", string(data))
	fmt.Println("Response status:", res.Status)

}
func main() {
	fmt.Println("learning crud")
	// performGetRequest()
	// performPostRequest()
	performUpdateRequest()

}
