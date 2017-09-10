package main

import (
	"sort"
	"fmt"
)

//This is from the "Sort Package" Video
type people []string


//This first main answers #1 exercise
//func main(){
//	studyGroup := people{"Zeno","John","Al","Jenny",}
//	fmt.Println(studyGroup)
//	sort.Strings(studyGroup)
//	fmt.Println(studyGroup)
//}


// Reverse #1
//func main(){
//	studyGroup := people{"Zeno","John","Al","Jenny",}
//	fmt.Println(studyGroup)
//	sort.Strings(studyGroup)
//	fmt.Println(studyGroup)
//	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
//	fmt.Println(studyGroup)
//}


//This answers #2 exercise
//func main()  {
//	s:=[]string{"Zeno","John","Al","Jenny"}
//	fmt.Println(s)
//	sort.Strings(s)
//	fmt.Println(s)
//}

//Reverse #2

//func main()  {
//	s:=[]string{"Zeno","John","Al","Jenny"}
//	fmt.Println(s)
//	sort.Strings(s)
//	fmt.Println(s)
//	sort.Sort(sort.Reverse(sort.StringSlice(s)))
//	fmt.Println(s)
//}



//#3 exercise
//func main() {
//	n:=[]int{7,4,8,2,9,19,12,32,3}
//	fmt.Println(n)
//	sort.Ints(n)
//	fmt.Println(n)
//}

// Reverse #3
func main() {
	n:=[]int{7,4,8,2,9,19,12,32,3}
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}