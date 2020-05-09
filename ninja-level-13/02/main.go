package main

import (
	"fmt"
	quote2 "udemy-hands-on/ninja-level-13/02/quote"
	word2 "udemy-hands-on/ninja-level-13/02/word"
)

func main() {
	fmt.Println(word2.Count(quote2.SunAlso))

	for k, v := range word2.UseCount(quote2.SunAlso) {
		fmt.Println(v, k)
	}
}
