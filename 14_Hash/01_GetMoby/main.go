package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main(){
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil{
		log.Fatal(err)
	}
	page,err:= ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil{
		log.Fatal(err)
	}
	//Call me Ishmale
	fmt.Printf("%s",page[35072:35088])
}
