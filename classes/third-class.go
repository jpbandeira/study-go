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

	// built-in functions

	// close#
	// It is used in channel communication.

	// len and cap#
	/*
		"len" gives the length of a number of types (strings, arrays, slices, maps, channels).
		"cap" is the capacity, the maximum storage (only applicable to slices and maps).
	*/

	// new and make#
	/*
		Both new and make are used for allocating memory.
		new(type)
		make(type)
		The function new is used for value types and user-defined types like structs.
	*/
	v1 := new(int)
	fmt.Println(v1)
	fmt.Println(*v1)
	*v1 = 10
	fmt.Println(*v1)
	/*
		Whereas, make is used for built-in reference types (slices, maps, channels).
		They are used like functions with the type as its argument:
	*/

	// copy and append#
	/*
		These are used for copying and concatenating slices
	*/

	// panic and recover#
	/*
		These both are used in a mechanism for handling errors
	*/

	// Recursive calls
	fmt.Println()
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i) // function call
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	fmt.Println()

	// A call to functions that using mutually recursive to know if a number is even or odd
	fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))   // 17 is odd: is true
	fmt.Printf("%d is odd: is %t\n", 18, odd(18))   // 18 is odd: is false

	fmt.Println(Factorial(5))

	//High order functions
	//Function can be used as a parameters in another functions
	//An example is a function that use another function with parameter to call it:
	//func callBack(y int, f func(int,int)) using func Add(a, b int) to call it and sum 2 numbers
	callBack(1, Add)

	//Functions used as a filter
	//In this case, I created a function that do a filter in a array with interger values and
	//the function odd user in parameter is the criteria to filter
	//filterFunction(arrayValues, odd) filter odd numbers in arrayValues
	arrayValues := []int{1, 2, 3, 4, 5, 7}
	fmt.Println(filterFunction(arrayValues, odd))
	fmt.Println(filterFunction(arrayValues, even))

	fmt.Println(filter(arrayValues, even))

	//We can create functions without name, these functions need to be assigned to a variable and this variable
	//have a reference to that function
	//or use auto call
	//This is when a function call byourself to execute
	//Three examples is:
	referenceToFunction := func(a, b int) int { return a + b }
	fmt.Println(referenceToFunction(1, 2))

	func(a, b int) {
		fmt.Println(a + b)
	}(1, 2)

	func() {
		fmt.Println(1 + 2)
	}()


	numberIncremented := incrementNumber(1)
	fmt.Println(numberIncremented(1))
}

//We can create a function that use closure function in return statement
//Example:
func incrementNumber(a int) func(b int) int {
	return func(b int) int {
		return b + a
	}
}

type flt func(int) bool

func filterFunction(array []int, toFilter flt) []int {
	var result []int
	for _, value := range array {
		if toFilter(value) {
			result = append(result, value)
		}
	}

	return result
}

//CHALLENGE to create a filter function to number from odd and even
func filter(sl []int, toFilter flt) (yes, no []int) {
	for _, value := range sl {
		if toFilter(value) {
			yes = append(yes, value)
		} else {
			no = append(no, value)
		}
	}

	return
}

func callBack(y int, f func(int, int)) {
	f(y, 2)
}

func Add(a, b int) {
	fmt.Println(a + b)
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

// Function to calculate the nth fibonacci number
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2) // recursive call
	}
	return
}

func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1) // even calls odd
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1) // odd calls even
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}

func Factorial(n uint64) (fac uint64) {
	if n <= 1 {
		return 1
	}
	fac = n * Factorial(n-1)
	return
}
