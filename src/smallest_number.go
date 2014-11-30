package main

import "fmt"

func main() {

	x := []int{
		96,86,68,57,82,63,70,37,
		34,83,27,19,97,9,17,7,94,
	}
	
	smallest := x[0]
	
	for _, value := range x {
	
		if value < smallest {
		
			smallest = value
			
		}
		
	}
	
	fmt.Println("smallest number in array: ", smallest)

}