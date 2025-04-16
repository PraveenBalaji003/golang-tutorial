package main

import "fmt"

func main() {
	a, b, c := 773, 885, 9017
	fmt.Printf("Decimal\t\tBinay\t\tHexadecimal\n")
	fmt.Printf("%d\t\t%b\t\t%#X\n", a, a, a)
	fmt.Printf("%d\t\t%b\t\t%#X\n", b, b, b)
	fmt.Printf("%d\t\t%b\t\t%#X\n", c, c, c)
}
