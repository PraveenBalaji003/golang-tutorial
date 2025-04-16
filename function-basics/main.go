package main

import "fmt"

// naked return
func swap(x, y string) (a, b string) {
	a, b = y, x
	return
}

// func swap(x string, y string) (string, string) {
// 	return y, x
// }

func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(5, 5)
	fmt.Println("result = ", result)
	str1, str2 := swap("Balaji", "Praveen")
	fmt.Println("", str1, str2)

}
