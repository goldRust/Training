package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)

}

func main() {
	txt, err := os.Open("not a file.txt")
	if err != nil {
		//You can just send an error report to the command line.
		fmt.Println("Error happened: ", err)
		//Or you can send the same info to a log file.
		log.Println("Error Happened: ", err)
		//Or you can send a log and close the program.
		log.Fatalln(err)
		//Or you can just close the program.
		panic(err)
	}

	fmt.Println(txt)
}

/*
//Println formats using the default formats for its operands and writes to standard output.
//Package log implements a simple logging package... writes to standard error and prints the datea nd time
of each logged message.
... The Fatal function calls os.Exit(1) after writing the log message ...
... The Panic functions call panic after writing the log message.



*/
