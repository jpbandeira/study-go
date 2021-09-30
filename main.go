package main

import (
	"fmt"
)

func Sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(Sum(1, 2))
	fmt.Println("TEST")

	var number float32 = 5.2 // Declared a floating point variable
	fmt.Println(number)      // Printing the value of variable
	fmt.Println(int(number)) // Convert the float value to int type

	// Constants

	const PI = 3.14159

	const B1 = "hello"
	const B2 string = "B constant"

	// const C2 = Sum(1, 2) not okay, becouse constante can't be evaluetade out of compile time

	// Overflow cases

	// const Ln2 = 0.693147180559945309417232121458\176568075500134360255254120680009
	// const Log2E = 1 / Ln2 // this is a precise reciprocal
	const BILLION = 1e9 // float constant
	fmt.Println(BILLION)
	const HARD_EIGHT = (1 << 100) >> 97
	fmt.Println(HARD_EIGHT)

	// Multiple assignments#

	// Untyped constants
	const BEEF, TOW, C = "meat", 2, "veg"
	// Typed contants
	const ONE, TWO2, THRE, FOR int = 1, 2, 3, 4

	// Enumerations
	const (
		ENUM1 int    = 1
		ENUM2 string = "two"
		ENUM3        = "untyped"
	)

	//The first use of iota gives 0. Whenever iota is used again on a new line, its value is incremented by 1
	const (
		UNKNOWN = iota
		FEMALE  = iota
		MALE    = iota
	)

	fmt.Println(UNKNOWN)
	fmt.Println(FEMALE)
	fmt.Println(MALE)

	// The naming of identifiers for variables follows the camelCasing rules (start with a small letter,
	//and every new part of the word starts with a capital letter). But if the variable has to be exported,
	//it must start with a capital letter, as discussed earlier in this chapter.

	// %d specifies format for integral values.
	// %s specifies format for string values.
	// %v specifies the general default format.

	fmt.Printf("MALE: %d\n", MALE)
	fmt.Printf("MALE: %v\n", MALE)

	var decision bool = true
	//We use %t, as a format specifier for booleans.
	fmt.Printf("Value of decision: %t\n", decision)

	// 	The architecture independent types have a fixed size (in bits) indicated by their names. For integers:

	// int8 (-128 to 127)
	// int16 (-32768 to 32767)
	// int32 (− 2,147,483,648 to 2,147,483,647)
	// int64 (− 9,223,372,036,854,775,808 to 9,223,372,036,854,775,807)
	// For unsigned integers:

	// uint8 (with the alias byte, 0 to 255)
	// uint16 (0 to 65,535)
	// uint32 (0 to 4,294,967,295)
	// uint64 (0 to 18,446,744,073,709,551,615)
	// For floats:

	//Unlike other languages, a float type on its own does not exist in Golang.
	//We have to specify the bits. For example, float32 or float64.
	// float32
	//Use float64 whenever possible, because all the functions of the math package expect that type.
	// float64

	var a int
	var b int32
	a = 15
	fmt.Println("a: ", a)
	// b = a + a -> compiler error = cannot use a + a (type int) as type int32 in assignment
	b = b + 5 // ok: 5 is a constant
	fmt.Println("b: " ,b)

	//So, to do assign attribute a to b, we need to convert this like this
	b = int32(a + a)
	fmt.Println("b: " ,b)

	//continue to Elementary Types
}
