package main

import (
	"fmt"

	"io/ioutil"
	"net/http"
)

func main() {
	//x:= 42
	//y:= "This is a string"
	//z:= true
	//fmt.Printf("%d - %b - %#X" ,x,x,x)
	doc, _ := http.Get("https://docs.google.com/document/d/1h5XVihDuqoXDk-gpjqnyx6P-nr5pctn9PPr_wixDZls/edit")
	docBod, _ := ioutil.ReadAll(doc.Body)
	fmt.Printf("%s", docBod)

}
