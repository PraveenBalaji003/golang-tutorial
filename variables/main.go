package main

import "fmt"

/*Statically Typed Language
We declare a variable of a certain type. It can only holds the values of that type.

Aggregate type -
Aggregate many values together.
Example - array, slice

Compound type -
Include values of different types.
Example - struct
*/

/*Default Values
int = 0
float = 0.0
boolean = false
string = ""
pointers, functions, interfaces, slices, channels, maps = nil
*/

/*Go Basic Types

bool

string

int int8 int16 int32 int64
uint uint8 uint16 uint32 uintptr

byte // alias for uint8

rune // alias for int32 // represents a Unicode Code Point

float32 float64

complex64 complex128
*/

func main() {
	//Use var when specificity is required
	var a int
	fmt.Println("a = ", a)
	a = 100
	fmt.Println("a = ", a)

	//short declaration operator
	//short declaration operators can be used only inside the functions
	b := 214.11
	fmt.Println("b = ", b)

	//blank identifier - discards the value
	c, d, e, _ := 1, 0.0, true, "string"
	fmt.Printf("c = %v\nd = %v\ne = %v\n", c, d, e)

}
