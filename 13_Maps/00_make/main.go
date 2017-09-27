package _0_make

import "fmt"

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good Morning."
	myGreeting["Jenny"] = "Bonjour."
	fmt.Println(myGreeting)

}
