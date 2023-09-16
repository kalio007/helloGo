package main

import (
	"fmt"
	"math"

	"rsc.io/quote"

	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("my favorite number is", rand.Intn(10))
	fmt.Printf("now lets print %g this. \n", math.Sqrt(7))
	fmt.Println(add(4, 5))
	fmt.Println(math.Pi)
	fmt.Println(quote.Go())
}
