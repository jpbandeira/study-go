package classes

import "fmt"

func Examples2() {

	// IF & ELSE EXAMPLES

	// The structure of if can start with an initialization statement
	// (in which a value is given to a variable). This takes the form
	// (the ; after the initialization is mandatory):
	// if initialization; condition {
	// 	// do something
	// }
	max := 20
	if value := 10; value >= max {
		fmt.Println("\nMaior: ", value)
	} else {
		fmt.Println("\nMenor: ", value)
	}

	if value := sumAB(10, 10); value >= max {
		fmt.Println("Maior: ", value)
	} else {
		fmt.Println("Menor: ", value)
	}

	// ERROS ON FUNCTIONS
	v, ok := retErro()
	fmt.Println(v, ok)

	anInt, _ := retErro()
	fmt.Println(retErro())
	fmt.Println(anInt)

	// switch case with conditions and assigne a value to the variable inside switch
	switch num1 := 100; {
	case num1 < 0:
		fmt.Println("Number is negative")

	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")

	default:
		fmt.Println("Number is 10 or greater")
	}

	str := "Go is a beautiful language!"

	// for range in string
	// pos will get the position of char and char will get te caracter value
	for pos, char := range str {
		// using continue to jump to the next loop
		if pos == 5 {
			continue
		}
		fmt.Printf("Character on position %d is: %c \n", pos, char)

		// using break to out of looping
		if pos == 8 {
			break
		}
	}

	usingLabels()

	fmt.Println(teste(-10))	
}

func usingLabels() {
LABEL1: // adding a label for location
	for i := 0; i <= 5; i++ { // outer loop
		for j := 0; j <= 5; j++ { // inner loop
			if j == 4 {
				continue LABEL1 // jump to the label
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	// writing a loop without a for using goto
	i := 0

HERE: // adding a label for a location
	fmt.Printf("%d", i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}

func sumAB(a int, b int) (value int) {
	value = a + b
	return
}

func retErro() (value int, err error) {
	value = 21
	if value > 20 {
		return
	}
	err = fmt.Errorf("Some error")
	return
}

func challengeSeasonOfAMonth(monthNumber int) (season string) {
	switch monthNumber {
	case 1, 2, 12:
		season = "Winter"
		return
	case 3, 4, 5:
		season = "Spring"
		return
	case 6, 7, 8:
		season = "Summer"
		return
	case 9, 10, 11:
		season = "Autumn"
		return
	default:
		fmt.Println("Dont exist")
	}

	return
}

func teste(x int) int {

	if x < 0 {
		fmt.Println()
		fmt.Println(x)
		return -x
	}
	return x
}
