package main

import "fmt"

type person struct {
	firstName string
	lastName string
	age int
}

func main() {

	exercise1()
	fmt.Println()
	exercise2()
	fmt.Println()

}

func exercise1(){
	x := 42
	fmt.Println(&x)
}

func exercise2(){

	p1 := person{"tony", "stark", 28}

	fmt.Println("before change: ", p1)

	changeMe(&p1,"steve", "rogers", 100 )

	fmt.Println("after change: ", p1)
}


func changeMe(p *person, newName string, newLastName string, newAge int){

	*p = person{
		firstName: newName,
		lastName: newLastName,
		age: newAge,
	}
}