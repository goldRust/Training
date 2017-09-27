package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	c := sender()
	fmt.Println(c)

}

func sender() chan string {
	out := make(chan string)
	go func() {
		ln := " "
		message, _ := fmt.Scanf("%s", ln)
		out <- message
	}()
	return out
}
