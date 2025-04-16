package main

import "fmt"

func main() {
	var a float32 = 244.12
	b := 100.21
	/*since default data type of float values will be f64. So, we can't able to store
	f64 value inside the f32 data type. So, we needs to do the type conversion */
	b = float64(a)
	fmt.Println("b = ", b)

	var x float32 = 25.01
	y := int(x)
	fmt.Printf("x = %d\t x = %T\n", y, y)

}
