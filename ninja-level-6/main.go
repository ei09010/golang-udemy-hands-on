package main

import (
	"fmt"
	"math"
	"reflect"
)

type person struct{
	first string
	last string
	age int
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	calculateArea() float64
}

func main() {

	defer exercise1()
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
	fmt.Println()
	exercise8()
	fmt.Println()
	exercise9()
	fmt.Println()
	exercise10()
}


func foo() int {
	return 123 * 123
}

func bar() (int, string){

	myString := fmt.Sprint("this is the best")
	var myInt int = 7

	return myInt, myString
}

func exercise1(){
	x, y := bar()
	fmt.Println(foo(), "\n", x, "\t", y)
}

func exercise2(){

	sliceArg := []int{1,2,4, 5,6, 7}

	x := foo2(sliceArg...)
	y := bar2(sliceArg)

	fmt.Println(x, y)
}


func foo2(x ...int) int{
	totalSum := 0

	for _,v := range x{
		totalSum += v
	}

	return totalSum
}

func bar2(mySlice []int) int {
	totalSum := 0

	for _,v := range mySlice{
		totalSum += v
	}

	return totalSum
}

func exercise3(){
	defer foo3()
	fmt.Println("hey you, this will run first")
}

func foo3(){

	defer func(){
		fmt.Println("foo DEFER ran")
	}()

	fmt.Println("not defered statement")
}


func(p person) speak(){
	fmt.Printf("hi, my name is %v %v and I'm %v years old\n", p.first, p.last, p.age)
}


func exercise4(){
	tonyStark := person{"Tony", "Stark", 28	}
	tonyStark.speak()
}


func exercise5(){
	sq1 := square{21.35}

	circ1 := circle{ 4}

	info(sq1, circ1)
}


func (s square) calculateArea() float64{
	return s.length * s.length
}

func (c circle) calculateArea() float64{
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s ...shape){
	for _,v := range s{
		fmt.Println("the area for shape", reflect.TypeOf(v),"is",v.calculateArea())
		}
}


func exercise6(){

	retFromAnonymous := func(x float64) float64{
		return math.Sqrt(x)
	}(4)

	fmt.Println("Sqr root is: ", retFromAnonymous)
}

func exercise7(){

	calculateSqr := func(x float64) float64{
		return math.Sqrt(x)
	}

	fmt.Println("Sqr root in a variable assigned function is: ", calculateSqr(4))
}


func exercise8(){
	returnedFunc := willRetFunc()
	fmt.Println(returnedFunc())
}

func willRetFunc() func() int {
	return func() int{
		return 2
	}
}


func exercise9(){

	fmt.Println(receivesNewFuncAsArg(enhanceName, "mario"))
}

func enhanceName(x string) string{

	return fmt.Sprint("Hello, and welcome mr.", x)
}

func receivesNewFuncAsArg(f func(x string) string, x string) string{

	return f(x)
}

func exercise10(){

	f := foo10()
	fmt.Println(f())
	fmt.Println(f())

	g := foo10()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}

func foo10() func() int {
		x := 0
		return func() int {
			x++
			return x
		}
}