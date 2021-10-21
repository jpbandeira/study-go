package main

import (
	"fmt"
	"strings"
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
	fmt.Println("b: ", b)

	//So, to do assign attribute a to b, we need to convert this like this
	b = int32(a + a)
	fmt.Println("b: ", b)

	// 	Format specifiers#
	// In format-strings, %d is used as a format specifier for integers (%x or %X can be used for a hexadecimal representation).
	// The %g is used for float types (%f gives a floating-point, and %e gives a scientific notation).
	// The %0nd shows an integer with n digits, and a leading 0 is necessary.
	// The %n.mg represents the number with a precision of m digits and width of n digits.
	// Instead of g, e and f can also be used. For example, the %5.2e formatting of the value 3.4 gives 3.40e+00.

	var c1 complex64 = 5 + 10i // Declaring complex num (real +imaginary(¡))
	fmt.Printf("The value is: %v", c1)

	// If re and im are of type float32, a variable c of type complex64 can be made with the function complex:
	// c = complex(re, im)

	var ch1 byte = 'A'
	fmt.Printf("\nThe value is: %v\n", ch1)
	fmt.Printf("The charactere is: %c\n", ch1)
	var ch2 byte = 65
	fmt.Printf("The value is: %v\n", ch2)
	fmt.Printf("The charactere is: %c\n", ch2)
	var ch3 byte = '\x41'
	fmt.Printf("The value is: %v\n", ch3)
	fmt.Printf("The charactere is: %c\n", ch3)

	var ch11 int = '\u0041'
	var ch22 int = '\u03B2'
	var ch33 int = '\U00101234'

	fmt.Printf("integer %d - %d - %d\n", ch11, ch22, ch33)        // integer
	fmt.Printf("character %c - %c - %c\n", ch11, ch22, ch33)      // character
	fmt.Printf("UTF-8 bytes %X - %X - %X\n", ch11, ch22, ch33)    // UTF-8 bytes
	fmt.Printf("UTF-8 code point %U - %U - %U", ch11, ch22, ch33) // UTF-8 code point

	fmt.Printf("\n")
	fmt.Printf(`This is a raw string \n`)
	fmt.Printf(` This is a raw string \n`)

	//Prefixes and suffixes#
	var str string = "This is an example of a string"
	fmt.Printf("\n T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("\n Finding prefix %t", strings.HasPrefix(str, "Th")) // Finding prefix

	fmt.Printf("\n Does the string \"%s\" have suffix %s? ", str, "ting")
	fmt.Printf("\n Finding suffix %t", strings.HasSuffix(str, "ting"))    // Finding suffix
	fmt.Printf("\n Finding suffix %t\n", strings.HasSuffix(str, "tring")) // Finding suffix

	str = "Hi, I'm Marc, Hi."
	fmt.Printf("The position of the first instance of\"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc")) // Finding first occurence
	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi")) // Finding first occurence
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi")) // Finding last occurence
	// In the next exemplo, the word Burger isn't exist in string variable.
	// So the value returned is -1
	fmt.Printf("The position of the first instance of\"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))

	// We can replace the first n occurrences of old in str by new. A cpy of str is returned, and if n is -1, all occurrences are replaced.
	// strings.Replace(str, old, new string, n int)
	fmt.Printf("%s\n", strings.Replace(str, "Hi", "teste", 1))
	fmt.Printf("%s\n", strings.Replace(str, "Hi", "teste", 2))
	str = "Hi, I'm Marc, Hi, Hi, Hi, Hi."
	fmt.Printf("%s\n", strings.Replace(str, "Hi", "teste", -1))

	// 	Count counts the number of non-overlapping instances of substring str in s with:
	// strings.Count(s, str string) int
	fmt.Printf("Count function: %d\n", strings.Count(str, "Hi"))

	// The Repeat function returns a new string consisting of count copies of the string s:
	var origS string = "Hi there! "
	var newS string
	newS = strings.Repeat(origS, 3) // Repeating origS 3 times
	fmt.Printf("The new repeated string is: %s\n", newS)

	fmt.Printf("To lower case string is: %s\n", strings.ToLower(str))
	fmt.Printf("To upper case string is: %s\n", strings.ToUpper(str))

	str = "\nHi, I'm Marc, Hi, Hi, Hi, Hi.\n"
	fmt.Printf(`NO Trim str from \n in the left and right sid: %s`, str)
	fmt.Printf("Trim str from second param: %s\n", strings.Trim(str, "\n"))

	str = "Hi Hello OLA"
	//	strings.Fields(str) split the string with all instances of one or more consecutive with space
	var splitedString []string = strings.Fields(str)
	fmt.Println(splitedString)

	// split string with all instances of the second param
	str = "Hi;Hello;OLA"
	splitedString = strings.Split(str, ";")
	fmt.Println(splitedString)

	// Create a new string with all elements from array separeted with the second param
	fmt.Println(strings.Join(splitedString, ","))

	// CONTINUE FROM TIMES AND DATES

}

// aliasing type
type Celsius float32
type Fahrenheit float32

// Function to convert celsius to fahrenheit
func toFahrenheit(t Celsius) Fahrenheit {

	celsiusConverted := Fahrenheit(t)
	var temp Fahrenheit
	temp = (celsiusConverted * 9 / 5) + 32
	return temp

}
