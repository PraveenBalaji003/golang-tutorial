package main

import "fmt"

/*
%v - default value
%s - string
%d - decimal
%T - data-type
\n - newline
\t - spacing
*/

func main() {
	a, b := 1, "two"
	fmt.Printf("a = %d\tb = %s\na = %T\tb = %T\n", a, b, a, b)
}
