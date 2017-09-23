package main

import (
	"fmt"
	"sort"
)
//his exercise 1 was different than mine
type people []string

func (p people) Len()int{return len(p)}
func (p people) Swap(i, j int){p[i],p[j]=p[j],p[i]}
func (p people) Less(i,j int)bool  {return p[i]< p[j]}


func main(){
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println(studyGroup)

}
// his #2 and #3 were the same.