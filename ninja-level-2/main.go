package main

import "fmt"

const(
	_ = iota
	yearCounter1 = 2020 + iota
	yearCounter2 = 2020 + iota
	yearCounter3 = 2020 + iota
	yearCounter4 = 2020 + iota

	untypedVar = 1234
	typedVar string = "this one is typed"

	)


func main() {

	exercise1()

	separator()

	exercise2()

	separator()

	exercise3()

	separator()

	exercise4()

	separator()

	exercise5()

	separator()

	exercise6()
}

func separator() {
	fmt.Println()
	fmt.Println()
}

func exercise1() {
	x := 24

	fmt.Printf("%d\t%b\t%#x", x, x, x)
}

func exercise2() {

	fmt.Println(42 == 42)
	fmt.Println(42 <= 42)
	fmt.Println(42 >= 42)
	fmt.Println(42 != 42)
	fmt.Println(42 < 43)
	fmt.Println(42 > 42)
}

func exercise3() {

	fmt.Println(untypedVar, typedVar)
}


func exercise4(){

	x := 123463

	fmt.Printf("%d\t%b\t%#x", x, x, x)

	y := x << 1

	fmt.Println()

	fmt.Printf("%d\t%b\t%#x", y, y, y)
}

func exercise5(){

	myLiteral := `"hello, 

this 				is a

lite	ral string"`
	fmt.Println(myLiteral)
}


func exercise6(){
	fmt.Printf("next year is %v and then is %v and bla is %v and then is %v",
				yearCounter1,yearCounter2, yearCounter3, yearCounter4)
}