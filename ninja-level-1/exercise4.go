package main

import "fmt"

type myFirstType int

var myVar myFirstType

func main() {

	fmt.Println(myVar)
	fmt.Printf("%T\n", myVar)
	myVar = 42
	fmt.Println(myVar)
}
