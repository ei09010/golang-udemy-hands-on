package main

import "fmt"


func main() {
	c := make(chan int)
	gen(c)
	receive(c)

	fmt.Println("about to exit")
}

func gen(c chan int) <-chan int {

	go func(){
		for i := 0; i < 5; i++ {
			fmt.Println("pushing value", i, "into the channel")
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int){

	for v := range c {
		fmt.Println("value pulled through range", v)
	}

}
