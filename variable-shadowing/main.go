package main

import "fmt"

var x int

func main() {
	//Shadowing - Redeclare the variable and initializing the value
	x := 50
	fmt.Println("main()- x: ", x)
	printX()
	fmt.Println("main()- x:", x)
}

func printX() {
	fmt.Println("printX()- x:", x)
	fmt.Println("Update...")
	x = 100
}
