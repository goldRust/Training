package main

import (
	"fmt"


)

func main(){

	c:= incrementer(2)

	var count int
	for n:= range c {
		count++
		fmt.Println(n)
	}
	fmt.Println("Final Count:", count)
}

func incrementer(n int) chan string{
	c:= make(chan string)
	done:= make(chan bool)

	for i:=1;i<=n;i++{
		go func(i int) {
			for k:=0;k<20;k++{
				c<- fmt.Sprint("Process: ", i," printing:",k)
			}
			done<-true
		}(i)

	}
	go func() {
		for i:=0; i<n; i++{
			<-done
		}
		close(c)
	}()

	return c
}