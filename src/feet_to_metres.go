package main

import "fmt"

func main() {
	
	fmt.Print("enter a measurement in feet: ")
	
	var length float64
	
	fmt.Scanf("%f", &length)

	output := length * 0.3048

	fmt.Println(length, " ft = ", output, " metres")
}

