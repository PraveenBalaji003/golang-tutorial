package main

import "gobuild/util"

// go build main.go - to compile a single file
// go build - use when you're in the root of a go module, it compile the entire project
// go build . - use when you are in a subdirectory of a Go module, it builds the module from that specific directory
// go install . - compile and generate go executables on go path (go env GOPATH)
func main() {
	util.WelcomeMessage()
}
