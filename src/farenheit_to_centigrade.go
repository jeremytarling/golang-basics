package main

import "fmt"

func main() {
	
	fmt.Print("enter a temperature in Fahrenheit: ")
	
	var tempFar float64
	
	fmt.Scanf("%f", &tempFar)
			
	tempCen := (tempFar - 32) * 5/9
	
	fmt.Println(tempFar, "Fahrenheit = ", tempCen, "Centigrade")
	
}