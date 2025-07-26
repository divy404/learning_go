package main

import "fmt"

// func divide(a,b float64) (float64,error) {
// 	if b==0 {
// 		return 0, fmt.Errorf("denominator must not be zero")
// 	}
// 	return  a/b, nil
// }
func divide(a,b float64) (float64,string) {
	if b==0 {
		return 0, "denominator must not be zero"
	}
	return  a/b, "nil"
}

func main() {
	fmt.Println("handling errors in go")
	// ans,_:=divide(5,2); this is one way to avoid error here using _ or we can do like this ->
	ans,_:=divide(5,2)
	// if err != nil {
	// 	fmt.Println("error handling")
	// }
	fmt.Println("the divison will give us: ", ans)

}