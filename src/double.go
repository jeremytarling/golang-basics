package main

import "fmt"

func main() {
	
	fmt.Print("enter a nuber: ")
	
	var input float64
	
	fmt.Scanf("%f", &input)
	
	output := input * 2
	
	fmt.Println("double ", input, " = ", output)

}