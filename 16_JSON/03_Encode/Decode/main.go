package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p1 Person
	rdr := strings.NewReader(`{"First":"James","Last":"Bond","Age": 20}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1.First, p1.Last, p1.Age)
}
