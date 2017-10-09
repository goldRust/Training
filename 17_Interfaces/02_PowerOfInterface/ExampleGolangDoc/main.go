package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byName []person

func (n byName) Len() int           { return len(n) }
func (n byName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n byName) Less(i, j int) bool { return n[i].Name < n[j].Name }

type byNameLength []person

func (nl byNameLength) Len() int           { return len(nl) }
func (nl byNameLength) Swap(i, j int)      { nl[i], nl[j] = nl[j], nl[i] }
func (nl byNameLength) Less(i, j int) bool { return len(nl[i].Name) < len(nl[j].Name) }

func main() {
	people := []person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)
	sort.Sort(byName(people))
	fmt.Println(people)
	sort.Sort(byNameLength(people))
	fmt.Println(people)

}
