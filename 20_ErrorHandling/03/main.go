package main

import (
	"fmt"

	"log"
)

//Customizing Error Messages with fmt


func main() {



_,err:=sqrt(-10)
if err != nil{
log.Fatalln(err)
}
}

func sqrt(f float64)(float64, error){
	if f<0 {
		return 0, fmt.Errorf("math error: square root of negative number: %v", f)
	}
	//implementation
	return 42,nil

}