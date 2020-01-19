package main

import "fmt"

// Variable declaration
var a int //declaration of a variable ‘a’ of type integer.
var b int
var c int = 5

func main() {
	a := 20
	b = 10
	fmt.Println(a, b, c)
}

// constants cannot be declared using the := syntax

/*
data = Numeric, String and Boolean types
Numeric = integers, floating-point and complex numbers

Integer types are signed or unsigned and attached int/uint (number of bits that data type can hold)
*/

var hexvalue = 0xff //prefix 0x for hexadecimal numbers
var octvalue = 034  // prefix 0 for octal numbers
