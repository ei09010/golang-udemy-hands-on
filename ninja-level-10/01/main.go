package main

import "fmt"

func main() {
	ca := make(chan int, 1)

	go func (){
		ca <- 42
	}()

	fmt.Println(<-ca)
}
