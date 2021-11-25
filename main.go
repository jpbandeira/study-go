package main

import (
	classe "study-go/classes"
)

func main() {
	//classe.BasicConstructsAndElementaryDataTypes()
	//classe.ControlEstructures()
	classe.Functions()
}

// Chalend to sum int with Variable Number of Arguments
func sumInts(list ...int) (sum int) {
	for _, value := range list {
		sum += value
	}
	return
}
