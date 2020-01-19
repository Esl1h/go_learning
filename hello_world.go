/*
The first line of the Golang code will always be the ‘package name’.
This is a compulsory statement and is the starting point to run your Golang code.
The second line of the Golang code is usually “import fmt (package name)”.
This is a preprocessor command that informs the Go compiler to include files in the package ‘fmt’.
*/
package main

import "fmt"

func main() {
	//in Golang a name is to be exported if the function begins with a capital letter
	fmt.Println("Hello World... again!")
}

// Save as hello_world.go and on terminal:
// go run hello_world.go
