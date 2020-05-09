package main

import "fmt"

var a int = 42
var b string = "James Bond"
var c bool = true

func main() {

		s := fmt.Sprintf("%v\t%v\t%v" ,a, b,c)

		fmt.Print(s)
}
