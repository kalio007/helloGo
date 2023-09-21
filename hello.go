package main

import (
	"fmt"
	"math"
	"time"

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

	var ToBe bool = false
	var s string
	fmt.Printf("Type: %T value: %v, %q", ToBe, ToBe, s)

	//you have to be very specific about your types
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	//in cont declaration we cant use :=
	const Truth = true
	fmt.Println("Go rules?", Truth)

	// for loop

	//declare
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//if statements
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//another
	if num := 7; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	}

	// time and switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Printf("its before noon")
	default:
		fmt.Println("Its after noon")
	}
}
