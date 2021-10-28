package classes

import "fmt"

func Functions() {

	fmt.Println()
	// Example of a function calling another functing in params
	fmt.Println(f1(f2(1, 2)))
	fmt.Println()

	// We can't create functions with the same name even though with diferents params and returns

}

func f1(a, b, c int) int { // taking three parameters and returning their sum
	fmt.Println(a, b, c)	
	return a + b + c
}

func f2(a, b int) (int, int, int) { // taking two parameters and returning their sum, difference and product
	n1 := a + b
	n2 := a - b
	n3 := a * b
	return n1, n2, n3
}
