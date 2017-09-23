package main

import "fmt"

func main(){
	var x= 12
	var y = 12.1230123
	//Conversion - int to float *Widening*
	fmt.Println(y*float64(x))

	//Conversion - foat to int *Narrowing*
	fmt.Println(int(y)*x)
}
