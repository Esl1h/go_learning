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

	namevalue := "Hello Everyone"
	fmt.Println(namevalue)

	// constants cannot be declared using the := syntax

	/*
				   data = Numeric, String and Boolean types
				   Numeric = integers, floating-point and complex numbers

				   Integer types are signed or unsigned and attached int/uint (number of bits that data type can hold)


		var hexvalue = 0xff //prefix 0x for hexadecimal numbers
		var octvalue = 034  // prefix 0 for octal numbers

				var abc bool = false // data type bool (boolean)

				   Logical operators like:
				   && (logical conjunction)
				   || (logical disjunction)
				   ! (logical negation)
				   == equality
				   != inequality - can be used on the boolean data types.

		var check [3]float32
		var check = [3]float32{50.0, 7.0, 8.0}

	*/
	var sample int = 10
	fmt.Println("%", &sample)
}
