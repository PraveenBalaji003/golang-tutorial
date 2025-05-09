package main

//get from main - go get github.com/Praveenmj005/calculator@main
//get from latest git tag version - go get github.com/Praveenmj005/calculator@latest
//get from specific version - go get github.com/Praveenmj005/calculator@v1.0.1
//clean mod cache - go clean -modcache
import (
	"fmt"

	"github.com/Praveenmj005/calculator"
)

func main() {
	fmt.Println("1 + 1 = ", calculator.Add(1, 1))
	fmt.Println("1 - 1 = ", calculator.Sub(1, 1))
	fmt.Println("2 * 2 = ", calculator.Mutiply(2, 2))
	fmt.Println("1 + 1 = ", calculator.Divide(3, 3))
	fmt.Println("1 + 1 = ", calculator.Modulo(1, 1))
}
