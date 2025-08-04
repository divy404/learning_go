package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}
func main () {
	fmt.Println("learning crud")
	 res , err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	 if err != nil {
		fmt.Println("error getting:",err)
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
		fmt.Println("error in decoding:",err)
		return
	}
	fmt.Println("Todo: ", todo)

}