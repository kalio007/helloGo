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
func swap(x, y string) (string, string) {
	return x, y
}

// named retun value
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// declaring a variable
var c, python, java bool

func main() {

	fmt.Println("my favorite number is", rand.Intn(10))

	fmt.Printf("now lets print %g this. \n", math.Sqrt(7))

	fmt.Println(add(4, 5))

	fmt.Println(math.Pi)

	fmt.Println(quote.Go())

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)

	// var i, j int = 1, 2
	// k := 3
	// c, python, java := true, false, "no!"

}
