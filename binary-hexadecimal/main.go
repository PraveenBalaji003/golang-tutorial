package main

import "fmt"

func main() {
	a, b, c, d, e := 1, 2, 3, 4, 5
	fmt.Printf("Value\tBinary\tHexadecimal\n")
	fmt.Printf("%v\t%b\t%#X\n", a, a, a)
	fmt.Printf("%v\t%b\t%#X\n", b, b, b)
	fmt.Printf("%v\t%b\t%#X\n", c, c, c)
	fmt.Printf("%v\t%b\t%#X\n", d, d, d)
	fmt.Printf("%v\t%b\t%#X\n", e, e, e)
}
