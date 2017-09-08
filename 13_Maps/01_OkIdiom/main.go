package main

import "fmt"

func main(){
	myGreeting := map[string][]float64{
		"Julie": {1,452e1/2,3,4%3,5+2,6/2,},

		"Ben": {456,32,111,45,2,},
		"Terry": {33,4,1,222,3},
		"Some Guy": {5,5,5,55,5,5,5,5,},
	}
	//fmt.Println(myGreeting)
	if val, exists:=myGreeting["Julie"];exists{

		fmt.Println("Val: ", val)
		fmt.Println("Exists: ", exists)

	}else{
		fmt.Println("That value doesn't exist.")
		fmt.Println("Val: ", val)
		fmt.Println("Exists: ", exists)
	}
//	fmt.Println(myGreeting)
}
