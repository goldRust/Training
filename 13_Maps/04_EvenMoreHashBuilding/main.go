package main

import "fmt"

func main() {
	word := "Hello"
	letter := rune(word[0])
	fmt.Println(letter)
	fmt.Printf("%T\n", letter)
	for i := 65; i < 123; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12)
	}
	n := HashBucket("Go", 12)
	fmt.Println(n)

}

func HashBucket(word string, buckets int) int {
	//letter := rune(word[0])
	letter := int(word[1])
	bucket := letter % buckets
	return bucket
}
