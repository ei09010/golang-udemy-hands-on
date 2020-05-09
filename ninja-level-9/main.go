package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var myWaitGroup sync.WaitGroup
var myMutex sync.Mutex

type person struct{
	first string
	last string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.last, p.first)
}

func saySomething(h human){
	h.speak()
}


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
}

	func exercise1(){

		fmt.Println("CPu count: \t", runtime.NumCPU())
		fmt.Println("Go Routine count: ", runtime.NumGoroutine())

		myWaitGroup.Add(2)

		go func(){
			fmt.Println("this is routine n.1")
			myWaitGroup.Done()
		}()

		go func(){
			fmt.Println("this is routine n.2")
			myWaitGroup.Done()
		}()
		fmt.Println("CPu count mid: \t", runtime.NumCPU())
		fmt.Println("Go Routine count mid: ", runtime.NumGoroutine())


		myWaitGroup.Wait()

		fmt.Println("And we are done")

		fmt.Println("CPu count in end : \t", runtime.NumCPU())
		fmt.Println("Go Routine count  in end : ", runtime.NumGoroutine())

	}

	func exercise2(){

		tony := person{ "tony","stark"}

		saySomething(&tony)


		//saySomething(tony) //not possible since the pointer receiver implements speak and not the type person receiver itself

	}

	func exercise3(){

		var counter int32
		routineAmount := 50

		myWaitGroup.Add(routineAmount)

		for i := 0; i < routineAmount; i++{
			go func(){
				tempCounter := counter
				tempCounter++
				counter = tempCounter

				fmt.Println("end value: ", counter)
				myWaitGroup.Done()
			}()
		}

		myWaitGroup.Wait()
		fmt.Println("about to finish this exercise")
	}

	func exercise4(){
		var counter int32
		routineAmount := 50

		myWaitGroup.Add(routineAmount)

		for i := 0; i < routineAmount; i++{
			go func(){

				myMutex.Lock()

				tempCounter := counter
				tempCounter++
				counter = tempCounter

				fmt.Println("end value: ", counter)
				myMutex.Unlock()
				myWaitGroup.Done()
			}()
		}

		myWaitGroup.Wait()
		fmt.Println("about to finish this exercise")

	}

	func exercise5(){

		var counter int32
		routineAmount := 50

		myWaitGroup.Add(routineAmount)

		for i := 0; i < routineAmount; i++{
			go func(){
				atomic.AddInt32(&counter, 1)

				fmt.Println("end value: ", atomic.LoadInt32(&counter))

				myWaitGroup.Done()
			}()
		}

		myWaitGroup.Wait()
		fmt.Println("about to finish this exercise")
	}

	func exercise6(){

		fmt.Println("my GOOS: ", runtime.GOOS)
		fmt.Println("my GOARCH: ", runtime.GOARCH)

	}