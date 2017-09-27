package main

import (
	"log"
	"net/http"

	"bufio"
	"fmt"
	"os"
)

func main() {

	res, err := http.Get("http://www.mieliestronk.com/corncob_lowercase.txt")
	if err != nil {
		log.Fatalln(err)
	}
	//bs,_:= ioutil.ReadAll(res.Body)
	//str:= string(bs)
	//fmt.Println(str)
	words := make(map[string]string)
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Println(os.Stderr, "reading input:", err)
	}

	var i int
	for k, _ := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++

	}

}
