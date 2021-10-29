package classes

import (
	"fmt"
)

func Functions() {

	fmt.Println()
	// Example of a function calling another functing in params
	fmt.Println(f1(f2(1, 2)))
	fmt.Println()

	// We can't create functions with the same name even though with diferents params and returns

	// Blank identifier
	// Here i dont need the second return, so the caractere _ is used to ignore this.
	result1, _ := namedReturnVariables(2)
	fmt.Println(result1)

	// Here we assigne the variable reply with the position memory to n and passing this in param to a function
	// to change this initial value
	n := 0
	reply := &n
	fmt.Println("Before Multiplication:", n)
	fmt.Println("Before Multiplication:", *reply)

	// Passing reply with the position memory from n in params
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
	fmt.Println("Multiply:", n)      // Multiply: 50

	// Passing a variable number of parameters
	// The parameter (...type) is used to pass any quantity values with the type used
	// Example: func Greeting(prefix int, who ...string) like Greeting(4, "Joe", "Anna", "Eileen")
	// This can be used if you dont know how many params are be used in function
	value1, value2 := Greeting(1, "value1", "value2")
	fmt.Println(value1, value2)
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

// Named return variables
func namedReturnVariables(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	//return x2, x3
	// Here we dont need to use variables in retur. Just a empty return are suficient
	// And variables are assigned with your zero value
	return
}

// Using a pointer to create a side effect param
// So here we are changin reply value just with a pointer and dont need to return this
func Multiply(a, b int, reply *int) {
	*reply = a * b // side-effect(changing n)
}

func Greeting(prefix int, who ...string) (int, []string) {
	return prefix, who
}


