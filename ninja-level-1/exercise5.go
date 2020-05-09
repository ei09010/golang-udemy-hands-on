package main

import "fmt"

type myFirstType5 int

var myVar5 myFirstType5

var myVarEx int

func main() {

	fmt.Println(myVar5)
	fmt.Printf("%T\n", myVar5)
	myVar5 = 42
	fmt.Println(myVar5)

	myVarEx = int(myVar5)

	fmt.Println("new variable after conversion: ",myVarEx)
	fmt.Printf("new variable type: %T", myVarEx)
}
