package main

import (
	"fmt"
	"udemy-hands-on/ninja-level-12/dog"
)


type canine struct{
	name string
	age int
}


func main() {
	fido := canine{"fido", dog.Years(28)}

	fmt.Println(fido.name ,"is",fido.age, "dog years")
}
