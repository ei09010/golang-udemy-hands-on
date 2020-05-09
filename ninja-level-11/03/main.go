package main

import (
	"fmt"
)

type customError struct {
	reason string
	userToBlame string
}

func (c customError) Error() string {
	return fmt.Sprintf("the one to blame is %v because %v", c.userToBlame, c.reason)
}

func main() {

	foo(customError{"hulk is a disaster", "thor"})
}

func foo(e error){

	fmt.Println("foo ran - ", e.(customError).reason) //assertion : something implementing an interface belonging to type, we assert the type
}