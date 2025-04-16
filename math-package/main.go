package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Pi = ", math.Pi)
	var a int
	for a < 5 {
		//It will prints the random number from 0 to 1
		fmt.Println(rand.Intn(2))
		a++
	}

	fmt.Println("Square root of 9 is", math.Sqrt(9))
}
