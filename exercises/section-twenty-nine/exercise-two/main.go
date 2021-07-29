package main

import (
	"fmt"
	"github.com/ThompsonJonM/go-learning/exercises/section-twenty-nine/exercise-two/quote"
	"github.com/ThompsonJonM/go-learning/exercises/section-twenty-nine/exercise-two/word"
	"strings"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
