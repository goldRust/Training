package main

import (
	"fmt"
	"sort"
)

func main(){
	student:= make([]string,0)
	student = append(student,"Brett")
	student = append(student,"Justin")
	student = append(student,"Bert")
	student = append(student,"Dustin")
	students:= make([][]string,0)
	students = append(students,student)
	where := sort.SearchStrings(students[0],"Dustin")
	fmt.Println(where,"---Big cap: ", cap(students),"---Lil cap:",cap(student))

}