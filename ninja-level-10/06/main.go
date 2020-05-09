package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 101; i++ {
			c <- i
		}
		close(c)
	}()

	for x := range c {
		fmt.Println(x)
	}

	fmt.Println(" About to exit")
}
