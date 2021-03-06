package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"James","Last":"Bond","Age":20}`)

	json.Unmarshal(bs, &p1)
	fmt.Println(p1.First, p1.Last, p1.Age)

}
