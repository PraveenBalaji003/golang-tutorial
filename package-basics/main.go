package main

/*
go mod init package name - go mod init package-basics used for initialize go projects
it must be need in our root directory when we work in package level
*/
import (
	"fmt"
	"package-basics/config"
)

func main() {
	fmt.Println("DB Username: ", config.USER_NAME)
	//variable/function name first character should be declared in caps' otherwise
	//it can't be available for usage outside the package
	//fmt.Println("DB Password: ", config.password)
}
