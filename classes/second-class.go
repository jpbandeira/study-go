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
