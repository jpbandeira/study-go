package main

func main() {
	//classe.BasicConstructsAndElementaryDataTypes()
	//classe.ControlEstructures()
	//classe.Functions()
}

func sumInts(list ...int) (sum int) {	
	for _, value := range list {
		sum += value
	}
	return
}
