package main

import (
	"fmt"
	"sort"
)

type Person struct{
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s : %d", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int { return len(a)}
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }


type ByName []Person

func (b ByName) Len() int { return len(b)}
func (b ByName) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByName) Less(i, j int) bool { return b[i].Name < b[j].Name }


func main() {

	sortByAge()
	sortByName()
}


func sortByAge(){

	people := []Person{
		{"Tony Stark", 28},
		{"Steve Rogers", 100},
		{"Natasha Romanoff", 27},
		{"Bruce Banner", 80},
		{"Clinton Barton", 35},
		{"Thor Odinson", 500},
	}

	fmt.Println(people)

	//conversion for ByAge type
	sort.Sort(ByAge(people))

	fmt.Println(people)

}


func sortByName(){

	people := []Person{
		{"Tony Stark", 28},
		{"Steve Rogers", 100},
		{"Natasha Romanoff", 27},
		{"Bruce Banner", 80},
		{"Clinton Barton", 35},
		{"Thor Odinson", 500},
	}

	fmt.Println(people)

	//conversion for ByAge type
	sort.Sort(ByName(people))

	fmt.Println(people)

}