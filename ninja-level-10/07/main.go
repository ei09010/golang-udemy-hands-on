package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	numberRoutines := 10

	for i := 0; i <= numberRoutines; i++{

		go func() {

			for i := 0; i <= 10; i++ {
				c <- i
			}

		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

}
