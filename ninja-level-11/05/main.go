package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main() {

}

func myProfileFinder(p person) string{

	return fmt.Sprintf("a given person profile has the name %v and %v with age %v", p.first, p.last, p.age)
}
