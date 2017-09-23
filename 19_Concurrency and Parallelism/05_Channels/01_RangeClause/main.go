package main

import "fmt"

func main (){
	channnel := make(chan int)

	go func() {
		for i:=0;i<10;i++{
			channnel<-i

		}
		close(channnel)
	}()
for n:= range channnel{
	fmt.Println(n)
}
}
