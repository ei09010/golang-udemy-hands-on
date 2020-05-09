package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {

		// they both work:

		// return -1, sqrtError{"lat 23", "long 245", errors.New("this error was returned")}
		return -1, sqrtError{"lat 23", "long 245", fmt.Errorf("the value passed was: %v", f)}
	}
	return 42, nil
}