package main

import "fmt"

const VALUE int64 = 100

func main() {
	//constant can't be declared by using short hand operator :=
	//We can't able to re-initialize the value of constants
	//VALUE = 50
	fmt.Println("value = ", VALUE)
}
