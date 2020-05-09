package main

import "fmt"

func main() {

	exercise1()
	fmt.Println()
	exercise2()
	fmt.Println()
	exercise3()
	fmt.Println()
	exercise4()
	fmt.Println()
	exercise5()
	fmt.Println()
	exercise6()
	fmt.Println()
	exercise7()
}

func exercise1(){

	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func exercise2(){
	for i := 65; i <= 90; i++{
		fmt.Println(i)

		for j := 0; j < 3; j++{
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func exercise3(){
	i := 1991
	for i <= 2019{
		fmt.Println(i)
		i++
	}
}

func exercise4(){
	i := 1991

	for{
		fmt.Println(i)
		if i == 2020{
			break
		}
		i++
	}
}

func exercise5(){
	for i := 10; i <= 100; i++{
		fmt.Println(i % 4)
	}
}

func exercise6(){
	x := "James Bond"

	if x == "James Bond" {
		fmt.Println(x)
	}
}


func exercise7() {

	x := "mario pereira"

	if x == "nice"{
		fmt.Println("nice")
	} else if x == "mario pereira"{
		fmt.Println("found it")
	}
}

func exercise8(){
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}

func exercise9(){
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}

}