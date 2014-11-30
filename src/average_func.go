package main

import "fmt"

func average(nums []float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

func main() {
	myNums := []float64{1,2,3,4,5}
	fmt.Println(average(myNums))
}